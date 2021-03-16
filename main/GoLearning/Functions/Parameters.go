package main

import "fmt"

// map,slice,chan,指针，interface默认以引用方式传递
func main() {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Print(a, b)

	println(test1("sum: %d", 1, 2, 3))
}

func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func test1(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}
