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

func checkExactlyOneBitSet(bitVector int) bool {
	return (bitVector & (bitVector - 1)) == 0;
}

func createBitVector(str string) int {
	bitVector := 0;
	for _, c := range str {
		index := c - rune('a');
		if index < 0 {
			fmt.Println("only a-z");
			os.Exit(-1);
		}
		mask := 1 << index;
		bitVector = bitVector ^ mask;
	}
	return bitVector;
}

// assuming str only contains a-z
func checkOptimized(str string) bool {
	var bitVector int = createBitVector(str);
	return bitVector == 0 || checkExactlyOneBitSet(bitVector);
}

func main() {
	str := string("")
	fmt.Println("Enter the string:")

	in := bufio.NewReader(os.Stdin)

	str, _ = in.ReadString('\n')
	if len(str) <= 1 {
		fmt.Println("invalid len for string");
		return;
	}
	str = str[:len(str) - 1];	// remove \n

	if checkOptimized(str) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
