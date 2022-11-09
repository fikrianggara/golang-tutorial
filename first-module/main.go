package main

import (
	"fmt"

	secondModule "github.com/fikrianggara/golang-tutorial/second-module"
)

func main(){
	now :=secondModule.GetTime()
	fmt.Println(now)
}