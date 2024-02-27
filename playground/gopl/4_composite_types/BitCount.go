package main

import (
	"crypto/sha256"
	"fmt"

	"go.learn/utils"
)

func BitsDifference(h1, h2 *[sha256.Size]byte) int {
	n := 0
	for i := range h1 {
		n += utils.BitsDifference(h1[i], h2[i])
	}
	return n
}

func main() {
	s1 := "unodostresquatro"
	s2 := "UNODOSTRESQUATRO"
	h1 := sha256.Sum256([]byte(s1))
	h2 := sha256.Sum256([]byte(s2))
	fmt.Println(BitsDifference(&h1, &h2))
}
