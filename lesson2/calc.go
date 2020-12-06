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
		err       error
	)

	fmt.Print("Введите первое число: ")
	_, err = fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Ошибка чтения первого числа")
		os.Exit(4)
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Ошибка чтения второго числа")
		os.Exit(5)
	}

	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^): ")
	_, err = fmt.Scanln(&op)
	if err != nil {
		fmt.Println("Ошибка чтения арифмметической операции")
		os.Exit(6)
	}

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Деление на ноль невозможно")
			os.Exit(3)
		}
		res = a / b
	case "^":
		res = math.Pow(a, b)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %f\n", res)
}
