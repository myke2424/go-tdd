package di

// Dependency Injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	// A buffer is a slice of bytes with read and write methods (append to slice /read values)
	// We can think of it as 'in-memory file'

	// This implements the Writer interface
	buffer := bytes.Buffer{}
	Greet(&buffer, "Mike")

	got := buffer.String()

	want := "Hello, Mike"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
