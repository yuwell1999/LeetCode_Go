package main

import "fmt"

func main() {
	type MyInt string
	fmt.Println(len(MyInt(111111))) // 输出为4

	type person struct { // 同一个包内名称不能重复
		//name string
		//age  int8
		//city string
		name, city string
		age        int
	}

	// 两种实例化方法
	var person1 person
	person1.name = "Jack"
	person1.city = "USA"
	person1.age = 29
	fmt.Println("p1", person1)

	p4 := person{
		name: "James",
		city: "LA",
		age:  36,
	}
	fmt.Println("p4", p4)

	var p2 = new(person) // p2是一个结构体指针
	p2.age = 30
	p2.name = "Henry"
	p2.city = "Australia"
	fmt.Println("p2:", p2) // p2: &{Henry Australia 30}

	p3 := &person{}
	p3.name = "Ford"
	p3.city = "Orlando"
	p3.age = 18
	fmt.Println("p3", p3)

	// 匿名结构体
	var user struct {
		name string
		age  int
	}
	user.name = "Rose"
	fmt.Printf("%#v\n", user)
}
