package hello

import "fmt"

// Using a const over hard-coding a string will improve performance
const defaultPrefix = "Hello m8!"
const englishPrefix = "sup "
const frenchPrefix = "bonjour"
const french = "french"
const english = "english"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

// We have a named return value (prefix string)
// This will create a variable called prefix inside the function
// It will be assigned the "zero" value depending on the type
// Just by calling "return", will return the prefix var
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case english:
		prefix = englishPrefix
	default:
		prefix = defaultPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("mike", "englsish"))
}

// in Go public functions start with a capital letter and private one start with a lowercase.

