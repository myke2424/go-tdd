package stringz

import "testing"

func TestReverse(t *testing.T) {
    got := Reverse("mike")
    want := "ekim"

    if got != want  {
        t.Errorf("got '%q' but want '%q'", got, want)
    }
}
