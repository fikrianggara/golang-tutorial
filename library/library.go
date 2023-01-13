package library

import "fmt"

func SayHello() {
	fmt.Println("hello from package library")
}

type Human struct {
	name string
	age  int
}

func CreateInstance(name string, age int) Human {
	return Human{name: name, age: age}
}

func (h Human) SayHello() {
	fmt.Println("Hello, my name is", h.name, "I'm", h.age, "years old")
}

func (h *Human) SetAge(age int) {
	h.age = age
}
