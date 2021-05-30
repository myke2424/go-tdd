package hello

import "testing"

// Test files need to be in the format: xxxxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T
// In order to use *testing.T type, you need to import "testing"
// t of type *testing.T, is our "hook" into the testing framework, so you can do things like t.Fail(), if you want to fail.

func TestHello(t *testing.T) {
	// Here we are introducing subtests. It is useful to group tests around a "thing" and then have
	// subsets for different scenarios

	// In go, you can declare funcs inside other functions and assign them to var (only in funcs)
	// t.Helper() is  needed to tell the test suite that this method is a helper..
	// If you don't include this, when the test fails, the error line will be line 22 (t.Errorf)
	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			// Errorf prints out the err msg and fail the test,
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people in spanish", func(t *testing.T) {
		got := Hello("charlie", "spanish")
		want := "Hello m8!charlie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello m8' in english when an empty string is supplied ", func(t *testing.T) {
		got := Hello("world", "english")
		want := "sup world"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say bonjour in french", func(t *testing.T) {
		got := Hello("myke", "french")
		want := "bonjourmyke"
		assertCorrectMessage(t, got, want)
	})
}

// TDD Cycle!
// Write a test...
// Make a compiler pass.
// Run the test, see that it fails and check the error msg is meaningful
// Write enough code to make the test pass
// REFACTOR!!!!

