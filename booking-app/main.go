package main

import (
	"booking-app/morestrings"
	"fmt"

	"booking-app/hello"

	"github.com/fikrianggara/golang-tutorial/time"
)

func main() {
	fmt.Print(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(time.GetTime()) 
	hello.HelloInIndonesia()

}