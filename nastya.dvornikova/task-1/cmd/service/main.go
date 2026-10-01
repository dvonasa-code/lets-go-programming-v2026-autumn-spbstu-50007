package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	input = strings.TrimSpace(input)
	parts := strings.Fields(input)

	if len(parts) != 3 {
		fmt.Println("Invalid operation")
		return
	}

	// 1. Проверка первого операнда
	a, err1 := strconv.ParseFloat(parts[0], 64)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}

	op := parts[1]

	// 2. Проверка второго операнда
	b, err2 := strconv.ParseFloat(parts[2], 64)
	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}

	// 3. Вычисление
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