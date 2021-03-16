package main

import "fmt"

// map,slice,chan,指针，interface默认以引用方式传递
func main() {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Print(a, b)
}

func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
