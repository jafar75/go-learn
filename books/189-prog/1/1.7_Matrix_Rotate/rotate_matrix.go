
package main

import (
	"fmt"
)

func rotate(matrix [][]int)  {
    var mid int = len(matrix) / 2;
    length := len(matrix);
    for i := 0; i < mid; i++ {
        for j := 0; j < length; j++ {
            tmp := matrix[i][j];
            matrix[i][j] = matrix[length - 1 - i][j];
            matrix[length - 1 - i][j] = tmp;
        }
    }

    for i := 0; i < length; i++ {
        for j := 0; j < (i + 1); j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j];
        } 
    }
}


func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	};

	fmt.Printf("original matrix:\n");
	fmt.Println(matrix);

	fmt.Printf("rotated(90 clockwise) matrix:\n");
	rotate(matrix);
	fmt.Println(matrix);
}



