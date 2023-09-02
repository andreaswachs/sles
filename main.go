package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// initialize an empty coefficients matrix
	coefficients := [][]float64{}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of equations: ")
	equations, _ := reader.ReadString('\n')
	equations = strings.TrimSuffix(equations, "\n")

	fmt.Print("Enter the number of variables: ")
	variables, _ := reader.ReadString('\n')
	variables = strings.TrimSuffix(variables, "\n")

	// convert the strings to integers
	equationsInt, _ := strconv.Atoi(equations)
	variablesInt, _ := strconv.Atoi(variables)
	variablesInt += 1 // add one for the augmented part

	// initialize the coefficients matrix
	for i := 0; i < equationsInt; i++ {
		coefficients = append(coefficients, make([]float64, variablesInt))
	}

	// read in one row of coefficients at a time
	for i := 0; i < equationsInt; i++ {
		fmt.Printf("Enter the coefficients for equation %d: ", i+1)
		coefficientsString, _ := reader.ReadString('\n')
		coefficientsString = strings.TrimSuffix(coefficientsString, "\n")

		// split the string into a slice of strings
		coefficientsSlice := strings.Split(coefficientsString, " ")

		// convert the strings to floats
		for j := 0; j < variablesInt; j++ {
			coefficients[i][j], _ = strconv.ParseFloat(coefficientsSlice[j], 64)
		}
	}

	// Now listen interactively for the user to do row operations.
	// Must support:
	// mul row scalar
	// add row row scalar
	// swap row row

	for {
		print(coefficients)

		fmt.Print("Enter a row operation: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		var row1, row2 int
		var scalar float64

		if strings.HasPrefix(input, "mul") {
			// parse the row and scalar
			fmt.Sscanf(input, "mul %d %f", &row1, &scalar)
			coefficients = mul(coefficients, row1-1, scalar)
		} else if strings.HasPrefix(input, "add") {
			// parse the rows and scalar
			fmt.Sscanf(input, "add %d %d %f", &row1, &row2, &scalar)
			coefficients = add(coefficients, row1-1, row2-1, scalar)
		} else if strings.HasPrefix(input, "sub") {
			// parse the rows and scalar
			fmt.Sscanf(input, "sub %d %d %f", &row1, &row2, &scalar)
			coefficients = add(coefficients, row1-1, row2-1, -scalar)
		} else if strings.HasPrefix(input, "swap") {
			// parse the rows
			fmt.Sscanf(input, "swap %d %d", &row1, &row2)
			coefficients = swap(coefficients, row1-1, row2-1)
		} else if strings.HasPrefix(input, "quit") {
			os.Exit(0)
		} else {
			fmt.Println("Invalid input")
		}

		println()
	}
}

func print(coefficients [][]float64) {
	// Print the matrix with a nice format
	for i := 0; i < len(coefficients); i++ {
		fmt.Print("[ ")
		for j := 0; j < len(coefficients[i]); j++ {
			// Print the augmented part with a nice format
			if j == len(coefficients[i])-1 {
				fmt.Print("| ")
			}
			fmt.Printf("%f ", coefficients[i][j])
		}
		fmt.Println("]")
	}

	fmt.Println()
}

func swap(coefficients [][]float64, row1, row2 int) [][]float64 {
	coefficients[row1], coefficients[row2] = coefficients[row2], coefficients[row1]
	return coefficients
}

func add(coefficients [][]float64, row1, row2 int, scalar float64) [][]float64 {
	for i := 0; i < len(coefficients[row1]); i++ {
		coefficients[row1][i] += scalar * coefficients[row2][i]
	}

	return coefficients
}

func mul(coefficients [][]float64, row1 int, scalar float64) [][]float64 {
	for i := 0; i < len(coefficients[row1]); i++ {
		coefficients[row1][i] *= scalar
	}

	return coefficients
}
