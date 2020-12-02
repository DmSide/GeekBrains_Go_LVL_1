package lesson2

import (
	"fmt"
	"math"
	"os"
)

func Calculator() {
	var (
		a, b, res float64
		op        string
	)

	fmt.Print("Введите первое число: ")
	_, _ = fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	_, _ = fmt.Scanln(&b)

	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^): ")
	_, _ = fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "^":
		res = math.Pow(a, b)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %f\n", res)
}
