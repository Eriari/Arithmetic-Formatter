package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO Problem 1: Buat fungsi penjumlahan
func Add(a int, b int) int {
	return a + b
}

// TODO Problem 2: Buat fungsi nilai maksimum
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// TODO Problem 3: Buat fungsi pencarian angka
func getNumber(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func ArithmeticArranger(problems []string, val ...bool) string {

	// ! DON'T CHANGE THIS CODE
	the_solutions := false
	if len(val) > 0 && val[0] {
		the_solutions = true
	}

	// TODO Problem and Testing 1: Buat batasan dengan 5 input data
	arrangedProblems := ""
	if len(problems) > 5 {
		return "Error: Too many problems."
	}

	// TODO Problem dan Testing 2: Buat batasan hanya menggunakan operator "+" dan "-"
	operations := make([]string, len(problems))
	for i, problem := range problems {
		parts := strings.Fields(problem)
		operations[i] = parts[1]
	}
	uniqueOperations := make(map[string]bool)
	for _, op := range operations {
		uniqueOperations[op] = true
	}

	for i, problem := range problems {
		parts := strings.Fields(problem)
		if parts[1] != "+" && parts[1] != "-" {
			return "Error: Operator must be '+' or '-'."
		}
		operations[i] = parts[1]
	}

	// TODO Problem dan Testing 3: Buat batasan dimana hanya terdapat angka dan tidak boleh lebih dari 4 angka
	numbers := []string{}
	for _, problem := range problems {
		parts := strings.Fields(problem)
		numbers = append(numbers, parts[0], parts[2])
	}
	for _, number := range numbers {
		if _, err := strconv.Atoi(number); err != nil {
			return "Error: Numbers must only contain digits."
		}

		if len(number) > 4 {
			return "Error: Numbers cannot be more than four digits."
		}

	}

	// TODO Problem dan Testing 4: Buat output sesuai dengan format yang berlaku. Lihat di README.md
	topRow := ""
	dashes := ""
	values := []int{}
	solutions := ""
	for i := 0; i < len(numbers); i += 2 {
		width := max(len(numbers[i]), len(numbers[i+1])) + 2
		topRow += fmt.Sprintf("%*s", width, numbers[i])
		dashes += strings.Repeat("-", width)
		value, _ := strconv.Atoi(numbers[i])
		if operations[i/2] == "+" {
			value += getNumber(numbers[i+1])
		} else {
			value -= getNumber(numbers[i+1])
		}

		values = append(values, value)
		solutions += fmt.Sprintf("%*d", width, value)
		if i != len(numbers)-2 {
			topRow += "    "
			dashes += "    "
			solutions += "    "
		}
	}

	bottomRow := ""
	for i := 1; i < len(numbers); i += 2 {
		width := max(len(numbers[i-1]), len(numbers[i])) + 1
		bottomRow += fmt.Sprintf("%s%*s", operations[i/2], width, numbers[i])
		if i != len(numbers)-1 {
			bottomRow += "    "
		}
	}

	if the_solutions {
		arrangedProblems = topRow + "\n" + dashes + "\n" + bottomRow + "\n" + solutions

	} else {
		arrangedProblems = topRow + "\n" + dashes + "\n" + bottomRow
	}

	return arrangedProblems
}
