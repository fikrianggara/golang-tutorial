package main

import (
	"booking-app/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Print(morestrings.ReverseRunes("!oG ,olleH"))
	cmp.Diff("Hello World", "Hello Go")

}