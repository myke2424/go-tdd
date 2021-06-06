package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"

// Testing this will be slow, it will take num seconds to run the test
// We have a depdency on Sleeping which we need to extract so we can control in our tests
// We can mock time.sleep with dependency injection
func Countdown(out io.Writer, num int, sleeper Sleeper) {
	for i := num; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
    sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, 3, sleeper)
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep() {
    time.Sleep(1 * time.Second)
}
