package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var augmentedLinePos = -1

func main() {
	// initialize an empty coefficients matrix
	coefficients := [][]float64{}
	inputLog := []string{}

	// Take input for a new row of coefficients
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Augmented line position (empty input for no augmented line): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input = strings.TrimSuffix(input, "\n")

	if input != "" {
		augmentedLinePos, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	fmt.Println("Input coefficient matrix row by row, separated by spaces when done, input empty line")

	// Take input for a new row of coefficients. If input is empty, stop taking input
	for {

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")

		if input == "" {
			break
		}

		// Split the input by spaces
		coefficientsStr := strings.Split(input, " ")

		// Convert the string coefficients to float64
		coefficientsFloat := []float64{}
		for _, coefficientStr := range coefficientsStr {
			coefficient, err := strconv.ParseFloat(coefficientStr, 64)
			if err != nil {
				fmt.Println(err)
				continue
			}

			coefficientsFloat = append(coefficientsFloat, coefficient)
		}

		coefficients = append(coefficients, coefficientsFloat)
	}

	for {
		print(coefficients)

		fmt.Print("Row operation: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")
		inputLog = append(inputLog, input)

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
		} else if strings.HasPrefix(input, "swap") {
			// parse the rows
			fmt.Sscanf(input, "swap %d %d", &row1, &row2)
			coefficients = swap(coefficients, row1-1, row2-1)
		} else if strings.HasPrefix(input, "quit") {
			os.Exit(0)
		} else if strings.HasPrefix(input, "print") {
			print(coefficients)
		} else if strings.HasPrefix(input, "div") {
			// parse the row and scalar
			fmt.Sscanf(input, "div %d %f", &row1, &scalar)
			coefficients = mul(coefficients, row1-1, 1/scalar)
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
			// Check if need to print augmented line
			if augmentedLinePos == j {
				fmt.Print("| ")
			}

			// Print coefficient with 2 decimal places
			fmt.Printf("%.2f\t", coefficients[i][j])
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
