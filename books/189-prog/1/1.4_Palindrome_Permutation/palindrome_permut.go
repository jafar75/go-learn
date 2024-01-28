package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(str string) bool {
	counts := make(map[rune]uint)
	runes := []rune(str)
	for i := 0; i < (len(runes) - 1); i++ { // -1 to avoid newline character
		if string(runes[i]) != " " {
			counts[runes[i]]++
		}
	}
	odd_repeats := 0
	for _, count := range counts {
		if count%2 != 0 {
			odd_repeats++
		}
	}
	return odd_repeats <= 1
}

func main() {
	str := string("")
	fmt.Println("Enter the string:")

	in := bufio.NewReader(os.Stdin)

	str, _ = in.ReadString('\n')

	if check(str) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
