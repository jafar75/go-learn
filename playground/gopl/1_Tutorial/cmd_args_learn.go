package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	// Exercise 1.3  about compare running time of above with join in below. will be done later
	fmt.Println(strings.Join(os.Args[1:], " "))
}
