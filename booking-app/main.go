package main

import (
	"booking-app/morestrings"
	"fmt"

	"booking-app/hello"

	"github.com/fikrianggara/golang-tutorial/time/timestamp"
)

func main() {
	fmt.Print(morestrings.ReverseRunes("!oG ,olleH"))
	timestamp.GetTime()
	hello.HelloInIndonesia()

}