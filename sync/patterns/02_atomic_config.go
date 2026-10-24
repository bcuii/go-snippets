//go:build ignore

// Hot-swap config: readers Load the current *Config without a lock;
// a reload builds a whole new Config and Stores the pointer in one atomic step.
// Readers see either the old config or the new one, never a half-updated mix.
// Rule: never modify a Config after storing it; build a new one instead.
//
// Run:   go run ./sync/patterns/02_atomic_config.go
// Check: go run -race ./sync/patterns/02_atomic_config.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	Timeout time.Duration
	Debug   bool
}

var current atomic.Pointer[Config]

func main() {
	current.Store(&Config{Timeout: time.Second})
	handle() // before the reload: old config

	var wg sync.WaitGroup
	for range 3 {
		wg.Go(handle) // running while the config is swapped: old or new, never a mix
	}
	reload()
	wg.Wait()

	handle() // after the reload: new config
}

func reload() {
	current.Store(&Config{Timeout: 2 * time.Second, Debug: true}) // e.g. after the config file changed
}

func handle() {
	cfg := current.Load() // take one snapshot and use it for the whole request
	fmt.Println("timeout:", cfg.Timeout, "debug:", cfg.Debug)
}
