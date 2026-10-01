package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	input := strings.TrimSpace(string(inputBytes))
	if input == "" {
		fmt.Println("Invalid operation")
		return
	}

	parts := strings.Fields(input)
	if len(parts) != 3 {
		fmt.Println("Invalid operation")
		return
	}

	a, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	op := parts[1]

	b, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var result float64
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = a / b
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(result)
}
