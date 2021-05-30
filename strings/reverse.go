package stringz

import "strings"

func Reverse(str string) string {
    reversed := make([]string, len(str))

    for i := len(str) -1; i >= 0; i-- {
        reversed = append(reversed, string(str[i]))
    }

    // Convert the slice of bytes into a string
    return strings.Join(reversed, "")
}
