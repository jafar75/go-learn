package main

import (
	"fmt"
	"strings"
)


func main() {
	str1 := string("");
	str2 := string("");
	fmt.Println("Enter the first string:")
	fmt.Scan(&str1);
	fmt.Println("Now please enter the second:");
	fmt.Scan(&str2);

	str1 = str1 + str1;

	if strings.Contains(str1, str2) {
		fmt.Printf("YES");
	} else {
		fmt.Printf("NO");
	}
}