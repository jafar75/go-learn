package main

import (
	"fmt"
	"strconv"
)


func main() {
	str := string("");
	fmt.Println("Enter the string:")
	fmt.Scan(&str);

	out := string(str[0]);
	repeats := 1;
	for i := 1; i < len(str); i++ {
		if str[i] == str[i - 1] {
			repeats ++;
		} else {
			out = out + strconv.Itoa(repeats);
			repeats = 1;
			out = out + string(str[i]);
		}
	}
	out = out + strconv.Itoa(repeats);
	if len(out) == len(str) {
		out = str;
	}
	fmt.Println("compressed string: ", out);

}
