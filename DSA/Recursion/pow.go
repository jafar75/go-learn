package main

import (
	"fmt"
)

func main() {
	var x float64 = 2.10000
	var n int = 3

	var neg bool = false
	if n < 0 {
		neg = true
	}
	out := PowPositive(x, uint(Abs(n)))
	if neg {
		out = 1 / out
	}
	fmt.Println("pow(x, n) = ", out)
}

func Abs(n int) uint {
	if n < 0 {
		return uint(-1 * n)
	}
	return uint(n)
}

func PowPositive(x float64, n uint) float64 {
	if n == 0 {
		return 1.0
	} else if n == 1 {
		return x
	}

	var out float64 = x
	p := uint(1)
	for ; p < n; p = p * 2 {
		out = out * out
	}
	return out / PowPositive(x, p-n)
}
