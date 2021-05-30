package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertCorrectSum := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got '%d', but want '%d'", got, want)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {

		arr := [5]int{10, 10, 10, 10, 10}

		got := Sum(arr)
		want := 50

		assertCorrectSum(t, got, want)

	})

	t.Run("collection of any size", func(t *testing.T) {
		slice := []int{5, 5, 5, 5}

		got := SumSlice(slice)
		want := 20

		assertCorrectSum(t, got, want)

	})

}

// Go does not let you test equality with slices, instead we can use reflect.DeepEqual to check if any
// two variables are the same.
// However, reflect isn't type safe, meaning it will compile even if use two different types
func TestSumAll(t *testing.T) {
	got := SumAll([]int{5, 5}, []int{10, 10, 10, 10})
	want := []int{10, 40}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' but want '%v'", got, want)
	}
}

func TestRecursiveSum(t *testing.T) {
	got := SumRecursive([]int{5, 5, 5})
	want := 15
	if got != want {
		t.Errorf("got '%d', but want '%d'", got, want)
	}

}

// A tail is all the items apart from the first one ("head")
func TestSumAllTails(t *testing.T) {
	assertCorrectTailSum := func(t testing.TB, got, want []int) {

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' but want '%v'", got, want)
		}
	}

	t.Run("Sum tails for two slices", func(t *testing.T) {
		got := SumAllTails([]int{2, 1}, []int{10, 20, 20})
		want := []int{1, 40}

		assertCorrectTailSum(t, got, want)
	})

	t.Run("Sum tails for an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{5, 10})
		want := []int{0, 10}

		assertCorrectTailSum(t, got, want)
	})

}
