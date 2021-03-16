package main

import (
	"fmt"
	"math"
)

func main() {
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(4))

	fn := func() { println("Hello World!") }
	fn()

	fns := []func(x int) int{
		func(x int) int {
			return x + 1
		},
		func(x int) int {
			return x + 2
		},
	}

	println(fns[1](100))

	// function as field
	d := struct {
		fn func() string
	}{
		fn: func() string {
			return "Hello World!"
		},
	}

	fmt.Println(d.fn())
}
