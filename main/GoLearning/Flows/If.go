package main

import "fmt"

func main() {
	var a = 10
	if a < 20 {
		fmt.Println("a小于20")
	} else {
		fmt.Println("a的值不小于20")
	}
	fmt.Println("a=", a)
}
