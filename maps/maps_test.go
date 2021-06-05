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

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		definition := "this is a test"

		err := dict.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, definition)

	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defintion := "this is a test"

		// Maps will overwrite values with the new provided value.
		dict := Dictionary{word: defintion}

		err := dict.Add(word, defintion)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, defintion)

	})
}

func TestUpdate(t *testing.T) {
     word := "test"
     def := "this is a test"

     dict := Dictionary{word:def}

     newDef := "updated test"
     err := dict.Update(word, newDef)

     assertDefinition(t, dict, word, newDef)
     assertError(t, err, nil)
}

func TestDelete(t *testing.T) {
    word := "test"
    def := "this is a test"

    dict := Dictionary{word: def}

    dict.Delete(word)

    _, err := dict.Search(word)

    if err != ErrNotFound {
        t.Errorf("expected %q to be deleted", word)
    }
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

func assertDefinition(t testing.TB, dict Dictionary, word, defintion string) {
	t.Helper()

	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("Failed to find added word", err)
	}

	if defintion != got {
		t.Errorf("got %q but want %q", got, defintion)
	}

}
