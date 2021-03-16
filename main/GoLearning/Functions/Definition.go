package main

import "fmt"

func main() {
	/**
	不支持重载和默认参数
	*/
}

func test(x, y int, s string) (int, string) {
	n := x + y
	return n, fmt.Sprintf(s, n)
}
