package main

import (
	hellos "golang-tutorial/package1"
	humans "golang-tutorial/package2"
)

//supaya variabel, properti atau method bisa diakses dari package lain
// huruf pertama pada variabel harus kapital
// package 'something' akan mengekspose file atau package dengan nama 'something'
// satu folder merupakan satu package

func initMultiPackageExample() {
	//from hello package, uwaw2.go file
	hellos.SayHelloFromUwaw2()
	//from hello package, uwaw.go file
	hellos.SayHello()

	//from human package human.go file
	var humanObj = humans.CreateInstance("fikri", 22)
	humanObj.SayHello()
	humanObj.SetAge(20)
	humanObj.SayHello()
}
