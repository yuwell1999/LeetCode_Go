package main

import "fmt"

func main() {
	//  Golang switch 分支表达式可以是任意类型，不限于常量。可省略 break，默认自动终止。
	grade := "B"
	_ = grade
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

	fmt.Printf("你的等级是 %s", grade)
}
