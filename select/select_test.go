package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Go stdlib has a http rest package where you can easily create a mock http server

// Make a WebsiteRacer function which takes two URLS and "races" them
// by hitting them w/ HTTP GET and returning the URL which returned first,
// if none of them return within 10seconds return an error.

func TestRacer(t *testing.T) {

	// Lets use mock http servers so we have reliable servers to test against that we control

	// httptest.NewServer takes a http.HandlerFunc which we are sending in as an anonymous function
	// http.HandlerFunc takes a response writer and a request pointer
	// This is standard for a HTTP server, this is how we would write a real HTTP server in go
	// httptest.NewServer finds an open port to listen on and we close it when we're done.
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	// By prefixing a function with 'defer' it willl now call that function at the end of this function
	// You implement defer usually for clean up actions (closing file, server etc), it improves code readability
	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got, _ := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

}
