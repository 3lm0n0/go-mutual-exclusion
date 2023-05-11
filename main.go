package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v map[string]int
}

// Inc increments the counter for the given key.
func (sc *SafeCounter) Inc(key string) {
	sc.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	sc.v[key] ++
	sc.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (sc *SafeCounter) Value(key string) int {
	sc.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer sc.mu.Unlock()
	return sc.v[key]
}

func main() {
	key := string("test")

	sc := SafeCounter{v: make(map[string]int)}
	for i:=0; i<1001; i++ {
		go sc.Inc(key) 
	}

	time.Sleep(time.Second)
	fmt.Println(sc.Value(key))
}