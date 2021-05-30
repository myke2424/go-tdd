package structs

import "math"

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Note: In Go you CANT have two methods of the same name, even if the parameters differ.
// However you could have functions with the same name that reside in different packages.
// We can define methods on our types, so rectangle and circle type can implement there own area method.
func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}

// Lets create a rectangle type that encapsulates both area and perimeter logic
// A method is a function with a reciever
// Methods are attached to a type

type Rectangle struct {
	Width  float64
	Height float64
}

// Were binding this method to the rectangle type
// r Rectangle is the mehod reciever
// It's go convetion to have the reciever variable be the first letter of the type
// e.g 'r' for Rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

// Interfaces allow  you to make functions that can be used with different types
// "Parametric polymorphism"
// This creates decoupled code while maintaing type saftey
// Lets create a shape interface type

type Shape interface {
	Area() float64
}

// This is quite different to interfaces in most other programming languages.
// Normally you have to do something like: class Rectangle implements interface Shape
// But in our case:
// Rectangle/Circle have a method called Area that returns float64, so it satisfies the shape interface
// In Go, interface resolution is implicit. If the type matches what the interface is asking for, it will compile.
