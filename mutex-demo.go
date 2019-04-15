package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 10000; i++ {
		go c.Inc("k")
	}
	time.Sleep(time.Second)
	defer fmt.Println(c.Value("k"))
}
