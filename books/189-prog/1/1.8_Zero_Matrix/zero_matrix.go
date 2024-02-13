package main


import (
	"fmt"
)


func setZeroes(matrix [][]int)  {
    c0 := 1;
    m := len(matrix);
    n := len(matrix[0]);

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0;
                if j == 0 {
                    c0 = 0; 
                } else {
                    matrix[0][j] = 0;
                }
            }
        }
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0;
            }
        }
    }
    if matrix[0][0] == 0 {
        for i := 0; i < n; i++ {
            matrix[0][i] = 0;
        }
    }
    if c0 == 0 {
        for j := 0; j < m; j++ {
            matrix[j][0] = 0;
        }
    }
}

func main() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	};

	fmt.Printf("original matrix:\n");
	fmt.Println(matrix);

	fmt.Printf("zero matrix:\n");
	setZeroes(matrix);
	fmt.Println(matrix);
}