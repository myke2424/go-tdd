package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		// f stands for float64 and .2 means print 2 decimal places
		t.Errorf("got '%.2f' but want '%.2f'", got, want)
	}
}

func TestArea(t *testing.T) {
	// By using the shape interface, we have decoupled the concrete types
	// This keeps things abstract
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			//%g gives us the full decimal
			t.Errorf("got '%g' but want '%g'", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)

	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}

// TDT here is referring to "Table driven tests"
// Each table entry is a complete test case with an input and expected result
// We define the input and expected result with an anonymous struct
// Using this approach, if we need to add a new shape w/ area, we can very easily
// add the table entry for that shape and expected value
// Usecases: Testing various implemenations of an interface or a lot of combinations of function arguments
func TestAreaTDT(t *testing.T) {

	// anonymous struct
	// we are declaring a slice of structs
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{ // We can use named values for the fields to be more explicit
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.15926535897934},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {

			got := test.shape.Area()

			if got != test.want {
				t.Errorf("%#v got '%g' want '%g'", test.shape, got, test.want)
			}

		})
	}
}
