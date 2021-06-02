package maps

import "testing"

// In golang, maps will return an empty string if the key doesn't exist. It will not throw an error like most languages.

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {

		got, _ := dict.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)

	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")

		assertError(t, err, ErrNotFound)
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}

}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
