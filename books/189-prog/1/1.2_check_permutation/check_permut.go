package main

import (
	"fmt"
)


func check(str1 string, str2 string) (bool) {
	if len(str1) != len(str2) {
		return false;
	}

	hash := make(map[uint]int);
	for _, char := range str1 {
		hash[uint(char)] += 1;
	}

	for _, char := range(str2) {
		hash[uint(char)] -= 1;
		if hash[uint(char)] < 0 {
			return false;
		}
	}

	return true;
}

func main() {
	str1 := string("");
	str2 := string("");
	fmt.Println("Enter the first string:")
	fmt.Scan(&str1);
	fmt.Println("Now please enter the second:");
	fmt.Scan(&str2);

	if check(str1, str2) {
		fmt.Printf("YES\n");
	} else {
		fmt.Printf("NO\n");
	}
}