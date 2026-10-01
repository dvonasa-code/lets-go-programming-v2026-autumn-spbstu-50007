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

	// 1. Проверка первого операнда
	a, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	// 2. Проверка знака операции (должна идти ДО проверки второго операнда)
	op := parts[1]
	if op != "+" && op != "-" && op != "*" && op != "/" {
		fmt.Println("Invalid operation")
		return
	}

	// 3. Проверка второго операнда
	b, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	// 4. Проверка деления на ноль
	if op == "/" && b == 0 {
		fmt.Println("Division by zero")
		return
	}

	// 5. Вычисление
	var result float64
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}

	fmt.Println(result)
}