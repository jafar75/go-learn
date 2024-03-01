// By using flag.Int(), flag will be pointer.
// But flag.IntVar() fills the variable as non-pointer

package main

import (
	"flag"
	"fmt"
)

func main() {
	var flagOne int

	flag.IntVar(&flagOne, "one", 1234, "1st flag in int")

	flagTwoPtr := flag.Int("two", 1234, "2nd flag in int (ptr)")

	flag.Parse()

	if flagOne > 100 {
		fmt.Println(">100")
	} else {
		fmt.Println("<=100")
	}

	fmt.Println("address of 2nd flag: ", flagTwoPtr)
	fmt.Println("amount of 2nd flag: ", *flagTwoPtr)
}
