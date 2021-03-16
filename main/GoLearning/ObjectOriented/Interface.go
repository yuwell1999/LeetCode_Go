package main

import "fmt"

type Sayer interface {
	say()
}

type dog struct {
}

type cat struct {
}

func (d dog) say() {
	fmt.Println("Dog barking")
}

func (c cat) say() {
	fmt.Println("cat meowing")
}

func main() {
	dog := dog{}
	dog.say()
}
