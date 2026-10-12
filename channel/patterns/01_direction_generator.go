//go:build ignore

// The most common use of directional channels: a function creates a channel,
// sends and closes it in its own goroutine, and returns only the receive end.
// Callers can receive, but can't send or close. time.After and ctx.Done() work this way.
//
// It pays off when values are not available right away (I/O, events, time).
//
// Run: go run ./channel/patterns/01_direction_generator.go
package main

import "fmt"

func generate(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range n {
			out <- i
		}
		close(out)
	}()
	return out
}

func main() {
	nums := generate(3)
	for v := range nums {
		fmt.Println("got", v)
	}
	// nums <- 1   // compile error: cannot send to receive-only channel
	// close(nums) // compile error: cannot close receive-only channel
}

// ---- Use cases ----
// Common use cases of this pattern, one per pair of functions.
// They are for reading only, main doesn't call them.

// 1. Wait for something to finish.
func waitForJob() {
	done := startJob()
	<-done // blocks until the job closes done
}

func startJob() <-chan struct{} {
	done := make(chan struct{}) // struct{} carries no data, it's only a signal
	go func() {
		// do some work

		// No data to deliver, only "the job is done", so close instead of sending a value.
		// close changes the channel's state for good: every receiver, now or later, sees it,
		// and the closer never blocks. A sent value reaches only one receiver, once.
		// Same as ctx.Done().
		close(done)
	}()
	return done
}

// 2. Run slow calls concurrently and collect the results later.
func callServices() {
	user := fetch("user") // the three calls run at the same time
	order := fetch("order")
	recommend := fetch("recommend")
	fmt.Println(<-user, <-order, <-recommend) // total time is about the slowest call
}

func fetch(service string) <-chan string {
	ch := make(chan string, 1) // buffer 1: the goroutine can finish even if nobody receives
	go func() {
		// call a remote service

		// The caller needs the result, so send it. close can't carry data:
		// a receive after close only gets the zero value "".
		// One result for one caller who receives once, so no close is needed.
		ch <- service + " ok"
	}()
	return ch
}

// 3. Stream data that arrives over time, e.g. paging through an API.
// Chaining such stages, each taking a channel and returning a new one, gives a pipeline,
// e.g. save(parse(pages(3))).
//
// pages returns the channel right away; the pages then flow through it one by one:
//
//	goroutine in pages      out           streamPages
//	------------------      ---           -----------
//	                                      range: waiting for a value
//	fetch page 1
//	out <- 1         ---- 1 ---->         p = 1, processing page 1
//	fetch page 2                          (still processing page 1)
//	out <- 2         ---- 2 ---->         p = 2, processing page 2
//	fetch page 3
//	out <- 3         ---- 3 ---->         p = 3, processing page 3
//	close(out)       -- closed ->         range sees closed and drained, loop ends
//
// The order is always 1, 2, 3. How long each side waits depends on which side is slower.
func streamPages() {
	for p := range pages(3) {
		fmt.Println("processing page", p) // the next page is fetched meanwhile
	}
}

func pages(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			// fetch a page
			out <- i // send: the caller needs the data
		}
		close(out) // close: no more pages, so the caller's range can end
	}()
	return out
}
