package humanClass

import "fmt"

func SayHello() {
	fmt.Println("hello from human's package, human.go")
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
