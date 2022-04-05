package main

import (
	"fmt"
	"os"
)

func addofMatrices() {
	fmt.Println("\n>>> Addition of two Matriks 3x3 <<<\n")

	var matrixA [3][3]int
	fmt.Println("Making the first matrix")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Input matrix element A-%d%d: ", i+1, j+1)
			fmt.Scanln(&matrixA[i][j])
			if matrixA[i][j] == -0 {
				matrixA[i][j] = 0
			}
		}
	}
	fmt.Println("\nOutput matrix A:")
	for i := 0; i < 3; i++ {
		fmt.Printf("[ ")
		for j := 0; j < 3; j++ {
			fmt.Print(matrixA[i][j], " ")
			if j == 2 {
				fmt.Println("]")
			}
		}
	}
	var matrixB [3][3]int
	fmt.Println("\nMaking the second matrix")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Input matrix element B-%d%d: ", i+1, j+1)
			fmt.Scanln(&matrixB[i][j])
			if matrixB[i][j] == -0 {
				matrixB[i][j] = 0
			}
		}
	}
	fmt.Println("\nOutput matrix B:")
	for i := 0; i < 3; i++ {
		fmt.Printf("[ ")
		for j := 0; j < 3; j++ {
			fmt.Print(matrixB[i][j], " ")
			if j == 2 {
				fmt.Println("]")
			}
		}
	}
	var result [3][3]int
	fmt.Println("\nResult of the addition of two matrices is...")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = matrixA[i][j] + matrixB[i][j]
		}
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("[ ")
		for j := 0; j < 3; j++ {
			fmt.Print(result[i][j], " ")
			if j == 2 {
				fmt.Println("]")
			}
		}
	}
	fmt.Println("\n=====================================")
}

func subofMatrices() {
	fmt.Println("\n>>> Subtraction of two Matriks 3x3 <<<\n")

	var matrixA [3][3]int
	fmt.Println("Making the first matrix")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Input matrix element A-%d%d: ", i+1, j+1)
			fmt.Scanln(&matrixA[i][j])
			if matrixA[i][j] == -0 {
				matrixA[i][j] = 0
			}
		}
	}
	fmt.Println("\nOutput matrix A:")
	for i := 0; i < 3; i++ {
		fmt.Printf("[ ")
		for j := 0; j < 3; j++ {
			fmt.Print(matrixA[i][j], " ")
			if j == 2 {
				fmt.Println("]")
			}
		}
	}
	var matrixB [3][3]int
	fmt.Println("\nMaking the second matrix")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Input matrix element B-%d%d: ", i+1, j+1)
			fmt.Scanln(&matrixB[i][j])
			if matrixB[i][j] == -0 {
				matrixB[i][j] = 0
			}
		}
	}
	fmt.Println("\nOutput matrix B:")
	for i := 0; i < 3; i++ {
		fmt.Printf("[ ")
		for j := 0; j < 3; j++ {
			fmt.Print(matrixB[i][j], " ")
			if j == 2 {
				fmt.Println("]")
			}
		}
	}
	var result [3][3]int
	fmt.Println("\nResult of the subtraction of two matrices is...")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = matrixA[i][j] - matrixB[i][j]
		}
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("[ ")
		for j := 0; j < 3; j++ {
			fmt.Print(result[i][j], " ")
			if j == 2 {
				fmt.Println("]")
			}
		}
	}
	fmt.Println("\n=====================================")
}

func multiofMatrices() {
	fmt.Println("\n>>> Multiply of two matrices <<<\n")

	var matrixA [3][3]int
	fmt.Println("Making the first matrix")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Input matrix element A-%d%d: ", i+1, j+1)
			fmt.Scanln(&matrixA[i][j])
			if matrixA[i][j] == -0 {
				matrixA[i][j] = 0
			}
		}
	}
	fmt.Println("\nOutput matrix A:")
	for i := 0; i < 3; i++ {
		fmt.Printf("[ ")
		for j := 0; j < 3; j++ {
			fmt.Print(matrixA[i][j], " ")
			if j == 2 {
				fmt.Println("]")
			}
		}
	}
	var matrixB [3][3]int
	fmt.Println("\nMaking the second matrix")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Input matrix element B-%d%d: ", i+1, j+1)
			fmt.Scanln(&matrixB[i][j])
			if matrixB[i][j] == -0 {
				matrixB[i][j] = 0
			}
		}
	}
	fmt.Println("\nOutput matrix B:")
	for i := 0; i < 3; i++ {
		fmt.Printf("[ ")
		for j := 0; j < 3; j++ {
			fmt.Print(matrixB[i][j], " ")
			if j == 2 {
				fmt.Println("]")
			}
		}
	}
	// Multiplying two matrices
	var result [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = 0
			for m := 0; m < 3; m++ {
				result[i][j] += matrixA[i][m] * matrixB[m][j]
			}
		}
	}
	fmt.Println("\nResult of multiplying two matrices is...")
	for i := 0; i < 3; i++ {
		fmt.Printf("[ ")
		for j := 0; j < 3; j++ {
			fmt.Print(result[i][j], " ")
			if j == 2 {
				fmt.Println("]")
			}
		}
	}
	fmt.Println("\n=====================================")
}

func determinantMatrix() {
	fmt.Println("\n>>> Find the Determinant of a matrix <<<\n")

	var matrix [9]int
	m := 0
	fmt.Println("Making a matrix")
	fmt.Println("Input matrix's element..")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Row-%d Column-%d: ", i+1, j+1)
			fmt.Scanln(&matrix[m])
			if matrix[m] == -0 {
				matrix[m] = 0
			}
			m++
		}
	}
	m = 0
	fmt.Println("\nOutput matrix:")
	for i := 0; i < 3; i++ {
		fmt.Printf("[ ")
		for j := 0; j < 3; j++ {
			fmt.Print(matrix[m], " ")
			if j == 2 {
				fmt.Println("]")
			}
			m++
		}
	}
	// Calculating determinant of a matrix
	var determinant, top1, top2, top3, bot1, bot2, bot3 int
	top1 = matrix[0] * matrix[4] * matrix[8]
	top2 = matrix[1] * matrix[5] * matrix[6]
	top3 = matrix[2] * matrix[3] * matrix[7]
	bot1 = matrix[6] * matrix[4] * matrix[2]
	bot2 = matrix[7] * matrix[5] * matrix[0]
	bot3 = matrix[8] * matrix[3] * matrix[1]
	determinant = top1 + top2 + top3 - bot1 - bot2 - bot3

	// Show the result
	fmt.Printf("\nDeterminant = %d + %d + %d - %d - %d - %d\n", top1, top2, top3, bot1, bot2, bot3)
	fmt.Printf("The result of determinant matrix is = %d", determinant)
	fmt.Println("\n=====================================")
}

func inverseofMatrix() {
	fmt.Println("\n>>> Find the Inverse of a matrix <<<\n")

	var matrix [3][3]float64
	fmt.Println("Making a matrix")
	fmt.Println("Input matrix's element..")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Row-%d Column-%d: ", i+1, j+1)
			fmt.Scanln(&matrix[i][j])
			if matrix[i][j] == -0 {
				matrix[i][j] = 0
			}
		}
	}
	fmt.Println("\nOutput matrix:")
	for i := 0; i < 3; i++ {
		fmt.Printf("[ ")
		for j := 0; j < 3; j++ {
			fmt.Print(matrix[i][j], " ")
			if j == 2 {
				fmt.Println("]")
			}
		}
	}
	// Find the determinant
	var determinant, top1, top2, top3, bot1, bot2, bot3 float64
	top1 = matrix[0][0] * matrix[1][1] * matrix[2][2]
	top2 = matrix[0][1] * matrix[1][2] * matrix[2][0]
	top3 = matrix[0][2] * matrix[1][0] * matrix[2][1]
	bot1 = matrix[2][0] * matrix[1][1] * matrix[0][2]
	bot2 = matrix[2][1] * matrix[1][2] * matrix[0][0]
	bot3 = matrix[2][2] * matrix[1][0] * matrix[0][1]
	determinant = top1 + top2 + top3 - bot1 - bot2 - bot3
	fmt.Printf("\nThe determinant matrix is %.0f\n", determinant)

	// Find the cofactor
	var cofactor [3][3]float64
	cofactor[0][0] = (matrix[1][1] * matrix[2][2]) - (matrix[1][2] * matrix[2][1])
	if cofactor[0][0] == -0 {
		cofactor[0][0] = 0
	}
	cofactor[0][1] = -((matrix[1][0] * matrix[2][2]) - (matrix[1][2] * matrix[2][0]))
	if cofactor[0][1] == -0 {
		cofactor[0][1] = 0
	}
	cofactor[0][2] = (matrix[1][0] * matrix[2][1]) - (matrix[1][1] * matrix[2][0])
	if cofactor[0][2] == -0 {
		cofactor[0][2] = 0
	}
	cofactor[1][0] = -((matrix[0][1] * matrix[2][2]) - (matrix[0][2] * matrix[2][1]))
	if cofactor[1][0] == -0 {
		cofactor[1][0] = 0
	}
	cofactor[1][1] = (matrix[0][0] * matrix[2][2]) - (matrix[0][2] * matrix[2][0])
	if cofactor[1][1] == -0 {
		cofactor[1][1] = 0
	}
	cofactor[1][2] = -((matrix[0][0] * matrix[2][1]) - (matrix[0][1] * matrix[2][0]))
	if cofactor[1][2] == -0 {
		cofactor[1][2] = 0
	}
	cofactor[2][0] = (matrix[0][1] * matrix[1][2]) - (matrix[0][2] * matrix[1][1])
	if cofactor[2][0] == -0 {
		cofactor[2][0] = 0
	}
	cofactor[2][1] = -((matrix[0][0] * matrix[1][2]) - (matrix[0][2] * matrix[1][0]))
	if cofactor[2][1] == -0 {
		cofactor[2][1] = 0
	}
	cofactor[2][2] = (matrix[0][0] * matrix[1][1]) - (matrix[0][1] * matrix[1][0])
	if cofactor[2][2] == -0 {
		cofactor[2][2] = 0
	}
	fmt.Println("\nCofactor of a matrix:")
	for i := 0; i < 3; i++ {
		fmt.Printf("[ ")
		for j := 0; j < 3; j++ {
			fmt.Print(cofactor[i][j], " ")
			if j == 2 {
				fmt.Println("]")
			}
		}
	}

	// Find the adjoin
	fmt.Println("\nThe adjoint is...")
	var adjoint [3][3]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			adjoint[i][j] = cofactor[j][i]
		}
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("[ ")
		for j := 0; j < 3; j++ {
			fmt.Print(adjoint[i][j], " ")
			if j == 2 {
				fmt.Println("]")
			}
		}
	}
	fmt.Println("\nThe result of inverse a matrix is...")
	var result [3][3]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = (1 / determinant) * adjoint[i][j]
		}
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("[ ")
		for j := 0; j < 3; j++ {
			fmt.Printf("%.3f ", result[i][j])
			if j == 2 {
				fmt.Println("]")
			}
		}
	}
	fmt.Println("\n=====================================")
}

func main() {

	for true {
		var menu int
		fmt.Println(">>> Welcome to Matrix 3x3 Operation <<<\n\nChoose your option..")
		fmt.Println("1. Addition of two matrices\n2. Subtraction of two matrices\n3. Multiply of two matrices\n4. Find the determinant of a matrix\n5. Find the inverse of a matrix\n0. Exit from program")
		fmt.Printf("Your option here: ")
		fmt.Scanln(&menu)

		switch menu {
		case 1:
			addofMatrices()
		case 2:
			subofMatrices()
		case 3:
			multiofMatrices()
		case 4:
			determinantMatrix()
		case 5:
			inverseofMatrix()
		case 0:
			var quit string
			fmt.Print("Thank you for using our program! :)\nType something to quit from the program.\n")
			fmt.Scan(&quit)
			os.Exit(0)
		default:
			fmt.Println("\nError input.")
		}
	}
}
