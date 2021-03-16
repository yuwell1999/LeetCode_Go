package main

import (
	"fmt"
)

func main() {
	//  Golang switch 分支表达式可以是任意类型，不限于常量。可省略 break，默认自动终止。
	grade := "B"
	marks := 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("优秀")
	case grade == "B", grade == "C":
		fmt.Println("良好")
	case grade == "D":
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	fmt.Printf("你的等级是 %s\n", grade)

	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型：%T\r\n", i)
	case int:
		fmt.Printf("x是int型")
	default:
		fmt.Print("未知型")
	}

	var k = 0
	switch k {
	case 0:
		fmt.Println("0")
		fallthrough // 无条件执行下面的case 1
		/*
		   Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；
		   而如果switch没有表达式，它会匹配true。
		   Go里面switch默认相当于每个case最后带有break，
		   匹配成功后不会自动向下执行其他case，而是跳出整个switch,
		   但是可以使用fallthrough强制执行后面的case代码。
		*/
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}
	//
}
