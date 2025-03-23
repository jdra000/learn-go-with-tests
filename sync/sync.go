package main

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()         // * Any goroutine calling Inc will acquire the lock on Counter if they are first
	defer c.mu.Unlock() // * All the goroutines will have to wait for it to be unlocked before getting access
	c.value++
}
func (c *Counter) Value() int {
	return c.value
}
