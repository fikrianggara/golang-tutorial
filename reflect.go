package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Age   uint
	Grade float32
}

func (s student) getPropertyInfo() {
	reflectValue := reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		//jika merupakan pointer, ambil elemen relevannya (properti object student)
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()
	fmt.Println(reflectType)
	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama		:", reflectType.Field(i).Name)
		fmt.Println("tipe Data	:", reflectType.Field(i).Type)
		//interface() berguna untuk menampilkan data yang bisa saja kompleks
		// ke bentuk string.
		fmt.Println("nilai		:", reflectValue.Field(i).Interface())
		fmt.Println(" ")
	}

}

func initReflectExample() {
	var student1 student = student{Name: "Fikri", Age: 14, Grade: 3.5}
	student1.getPropertyInfo()
}
