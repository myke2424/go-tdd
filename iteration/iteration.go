package iteration

func Repeat(char string, repeatCount int) string {
    // Use "var" when you want to zero value a variable, if you're initializing it, use the := syntax.
    var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += char
	}
	return repeated
}
