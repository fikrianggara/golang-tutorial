package humanClass

import "fmt"

type Human struct {
	Name string
	Age  int
}

func CreateInstance(name string, age int) Human {
	return Human{Name: name, Age: age}
}

func (h Human) SayHello() {
	fmt.Println("Hello, my name is", h.Name)
}
