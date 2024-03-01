// recover() method captures what the panic() is called with
// calls to defer() is pushed to a list and will be executed in a LIFO manner

package main

import "fmt"

const number = 3

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f, panic arg: ", r)
		}
	}()
	fmt.Println("Calling divide.")
	divide(number, 0)
	fmt.Println("Returned normally from divide.")
}

func divide(number int, i int) {
	if i > number {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("panic called with %v", i))
	}
	defer fmt.Println("Defer in divide", i)
	fmt.Println("Printing in divide", i)
	divide(number, i+1)
}
