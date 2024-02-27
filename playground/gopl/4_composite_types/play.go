package main

import (
	"fmt"
)

func main() {
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "9", GBP: "!", RMB: "R", 6: "Six"}
	fmt.Println(RMB, symbol[RMB])
	fmt.Println(symbol[4])
	fmt.Println(len(symbol))
}
