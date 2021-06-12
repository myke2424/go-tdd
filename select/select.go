package main

import (
	"fmt"
	"net/http"
	"time"
)

// We're going to use the select statement which helps us synchronise processes easily
// Lets test the speed of the websites concurrently
// Return an error if it takes longer than 10 seconds.

var tenSecondTimeOut = 10 * time.Second

// Our client code can use this function where our tests can use the configurable racer so we can configure the timeout
func Racer(a, b string) (winner string, error error) {
     return ConfigurableRacer(a, b, tenSecondTimeOut)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// Whichever ping writes to its channel first will result in the url being returned (winner)
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

	// time.After is very handy when using select.
	// With channels, you could potentially write code tha blocks forever if the channels
	// we're listening to never return a value.

	// time.After returns a chan (like ping) and will send a signal down it after the amount of time we defined.
    
}

// Creates a chan struct{} and returns it.
// In our case, we don't care what type is sent to the channel, we just want to signal we are done so we close the channel
// chan struct{} is the smallest data type available from a memory perspective so that's why we use that over bool
// Since we're not sending anything on the channel, this is ok.
func ping(url string) chan struct{} {
	// Always use 'make' for channels, channels zero value is nil if declare with 'var'
	// If you send a value to a nil channel with '<-' it will block forever
	ch := make(chan struct{})

	// Start a goroutine that closes the channel once the http request is done.
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

// time.Since returns a time.Duration type
func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)

	return time.Since(start)
}
