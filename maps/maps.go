package maps

import "errors"

type Dictionary map[string]string

// We can return multiple values in GO, e.g. a string/error, remember errors can be nil
// Errors can be converted to a string with the .Error method

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	// Map lookup can return 2 values. The second value is a boolean which indicates if the key was found succesfully.
	// If the key isn't found, the value is an empty string.
	value, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	// Remember nil is considered as an error type, so this return works.
	return value, nil
}
