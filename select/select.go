package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// * For the main function and test the the fastest one
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// * For testing if a server doesn't respond within specified time
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// * select allows you to wait on multiple channels.
	// * the first one to send a value "wins" and the code underneath the case is executed
	// * Whichever one writes to its channel first will have is code executed in the select
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	// * time.After returns a channel (like ping) and will send a signal down it after the amount you define
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
func ping(url string) chan struct{} {
	// * we don't care what type is sent to the channel
	// * we just want to signal we are done and closing the channel
	// * struct{} is the smallest data type available from a memory perspective
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
