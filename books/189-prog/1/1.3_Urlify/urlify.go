package main

import (
	"fmt"
	"bufio"
	"os"
)


func urlify(str string, len uint) (string) {
	var i uint = 0;
	out := "";
	for i = 0; i < len; i++ {
		if str[i] != ' ' {
			out = out + string(str[i]);
		} else {
			out = out + "%20";
		}
	}
	return out;
}

func main() {
	len := uint(0);
	str := string("");
	fmt.Println("Enter the string:");

	in := bufio.NewReader(os.Stdin);

	str, _ = in.ReadString('\n');

	fmt.Println("Length:");
	fmt.Scan(&len);

	fmt.Printf("URLify: %s\n", urlify(str, len));
}