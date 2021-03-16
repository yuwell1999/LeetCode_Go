package main

import (
	"fmt"
)

func add(a, b int) (c int) {
	c = a + b
	return // 没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。
}

func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a + b) / 2

	return
}

func main() {
	var a, b = 1, 2
	c := add(a, b)
	sum, avg := calc(a, b)
	fmt.Println(a, b, c, sum, avg)
}
