package main

import (
	"fmt"
	"go-tutorial/humanClass"
)

func main() {
	var firstname string = "Fikri"
	lastname := "Septrian"
	var age int
	var isMan bool = true
	var sex string

	if isMan {
		sex = "Man"
	} else {
		sex = "Woman"
	}

	var academics []string = []string{"SD attaufiq", "SMPN 2 Kota Jambi", "SMAN 3 Kota Jambi", "Politeknik Statistika STIS"}

	age = 22
	var ageClasses string

	if age < 10 {
		ageClasses = "young"
	} else if age < 17 {
		ageClasses = "young teenage"
	} else if age < 25 {
		ageClasses = "mature teenage"
	} else {
		ageClasses = "mature"
	}

	fmt.Println("Hello, my name is", firstname, lastname)
	fmt.Println("I'm a ", ageClasses, sex)
	fmt.Println("my academics :", academics)

	// var fruits = [4]string{"apple",
	// 	"grape",
	// 	"banana",
	// 	"melon"}

	// // i dipake, kalo ngga dipake bisa dibuat sebagai '_' aja
	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
	// }

	// for _, fruit := range fruits {
	// 	fmt.Println("buah", fruit)

	// }

	// temp := fruits[0:3] //slicing
	// fmt.Println(temp)
	// fruits[0] = "strawberry"
	// fmt.Println(temp)
	//slicing menyimpan alamat dari elemen,
	//jadi ketika ada perubahan nilai dari fruits, maka temp juga ikut berubah

	var fikri humanClass.Human = humanClass.CreateInstance("Fikri", 22)
	fikri.SayHello()
}
