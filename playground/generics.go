// a simple code to learn basics of generics in go 1.18 and above
package main

import (
	"fmt"
)

func MaxInt64(nums []int64) (int64, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("invalid input!")
	}

	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	
	return max, nil
}

func MaxFloat(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("invalid input!")
	}

	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	
	return max, nil
}

func MaxGeneric[V int64 | float64](nums []V) (V, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("invalid input!")
	}

	// max := nums[0]
	var max V = nums[0]		// although above line is working, this way of declaration better shows the generics
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	
	return max, nil
}

func main() {
	ints := []int64{3, 4, 12, 1}
	floats := []float64{3.1, 2.2, 1.7, 11.5, 7.5}

	maxInt, _ := MaxInt64(ints)
	fmt.Println("MaxInt64: ", maxInt)

	maxFloat, _ := MaxFloat(floats)
	fmt.Println("MaxFloat: ", maxFloat)

	maxInt, _ = MaxGeneric[int64](ints)
	fmt.Println("MaxGeneric[int64]: ", maxInt)

	// maxFloat, _ = MaxGeneric[float64](ints)  // compile error
	maxFloat, _ = MaxGeneric[float64](floats)
	fmt.Println("MaxGeneric[float64]: ", maxFloat)

}