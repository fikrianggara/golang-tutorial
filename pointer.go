package main

import "fmt"

func setA(x *int, y int) {

	//mengubah nilai yang ditunjuk x menjadi y
	*x = y
	fmt.Println("change value with address :", x, "to:", y)
}

func initPointerExample() {
	var a int = 10
	var b *int
	b = &a

	fmt.Println(" a :", a)
	fmt.Println("&a :", &a)
	fmt.Println("b:", b)
	fmt.Println("*b", *b)
	fmt.Println("&b", &b)

	*b = 5
	fmt.Println(" a :", a)
	fmt.Println("*b", *b)

	setA(&a, 15)
	fmt.Println("a:", a)
	// ketika b merupakan pointer dari a,
	// maka berlaku two-way data binding.
	// artinya, jika nilai variabel yang ditunjung variabel pointer diubah (mengubah *b)
	// maka variabel yang ditunjuk (a), juga berubah

}
