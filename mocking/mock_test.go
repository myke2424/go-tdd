package main

import (
	"bytes"
	"testing"
)

// Here we our mocking time.Sleep with our own type that implements the Sleeper interface
// "Spies" are a kind of mock which can record how a dependency is used.
// They can record arguments sent in, how many times it has been called etc
// In our case, we're keeping track of how many times Sleep() is called


type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	// Create a buffer we can use for our Writer interface
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, 3, spySleeper)

	got := buffer.String()

	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("Not enough calls to sleeper, expected 4 but got %d", spySleeper.Calls)
	}
}
