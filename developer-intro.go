package main

import (
	"fmt"
)

func InitDeveloperIntro() {
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

}
