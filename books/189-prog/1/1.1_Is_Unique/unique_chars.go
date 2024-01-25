package main

import (
	"fmt"
	"github.com/bits-and-blooms/bitset"
)


func checkUniqueCharacters(str string) (bool) {
	var hash bitset.BitSet;
	for _, char := range str {
		if hash.Test(uint(char)) {
			return false;
		}
		hash.Set(uint(char));
	}
	return true;
}

func main() {
	str := string("");
	fmt.Println("Enter the string:")
	fmt.Scan(&str);
	if checkUniqueCharacters(str) {
		fmt.Printf("Characters of the string: %s are all unique!\n", str);
	} else {
		fmt.Printf("There are some non-unique characters in string: %s\n", str);
	}
}