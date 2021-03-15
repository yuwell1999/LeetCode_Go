package main

import "fmt"

type person struct { // 同一个包内名称不能重复
	//name string
	//age  int8
	//city string
	name, city string
	age        int
}

// 构造函数
func newPersonConstructor(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func (p person) Method() {
	fmt.Println("person结构体的Method方法被调用！" + " " + p.name + "所在的城市为" + p.city)
}

// 指针类型接收者
func (p *person) setAge(age int) {
	p.age = age
}

func main() {
	type MyInt string
	fmt.Println(len(MyInt(111111))) // 输出为4

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

	// 不写键，直接写值
	// 必须初始化所有字段，且顺序与结构体中声明一致
	// 不能和键值初始化方式混用

	p5 := person{
		"prof",
		"北京",
		35,
	}
	fmt.Println("p5", p5)

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

	p6 := newPersonConstructor("White", "San Fransisco", 31)
	p6.setAge(30)
	fmt.Println("p6", p6) // p6 &{White San Fransisco 31}
	p6.Method()

	personInfo := PersonInfo{
		Name: "小红",
		City: "Changsha",
		Address: Address{
			road: "Wuyi Road",
		},
	}
	fmt.Println("personInfo", personInfo)
}

// 嵌套结构体
type Address struct {
	road string
}

type PersonInfo struct {
	Address Address
	Name    string
	City    string
}
