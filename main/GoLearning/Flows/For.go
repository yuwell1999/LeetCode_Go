package main

import "fmt"

func main() {
	/**
	for循环有三种形式
	    for init; condition; post { }
	    for condition { }
	    for { }

	*/

	var b = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}
	for a := 0; a < 10; a++ {
		fmt.Printf("a的值为：%d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a的值为：%d\n", a)
	}

	for i := range numbers {
		fmt.Printf("第%d位x的值=%d\n", i, numbers[i])
	}

	// 循环嵌套输出2到100的素数
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}

		if j > (i / j) {
			fmt.Printf("%d 是素数\n", i)
		}
	}
}
