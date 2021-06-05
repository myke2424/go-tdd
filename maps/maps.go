package maps

type Dictionary map[string]string

// We can return multiple values in GO, e.g. a string/error, remember errors can be nil
// Errors can be converted to a string with the .Error method


const (
 ErrNotFound = DictionaryErr("could not find the word you were looking for")
 ErrWordExists = DictionaryErr("Word already exists in dict")
 ErrWordDoesNotExist = DictionaryErr("Cannot update a word that doesnt exist")
)

type DictionaryErr string

// Here we create our own err type which implements the error interface (Error())
// This makes the errors more reusable and immutable
func (e DictionaryErr) Error() string {
    return string(e)
}

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

func (d Dictionary) Update(word, newValue string) error {
    _, err := d.Search(word)

    switch err {
        case ErrNotFound:
             // It's good to use a precise error here like this
             return ErrWordDoesNotExist
        case nil:
             d[word] = newValue
        default:
             return err
    }

    return nil
}

func (d Dictionary) Delete (word string) {
    // Go has a builtin delete keyword that works on maps
    // It takes two args, the map and key to be deleted
    // It returns nothing
    delete(d, word)
}

// An interesting property of maps is that you can modify them without passing an address to it
// e.g. &myMap

// Maps can also be nil values.
// A nil map behaves like a normal map when reading, but attemps to write to a nil map will cause a runtime err


func (d Dictionary) Add(key string, value string) error {
    _, err := d.Search(key)

    switch err {
    case ErrNotFound:
         d[key] = value
    case nil:
         return ErrWordExists
    default:
         return err
    }

    return nil

}
