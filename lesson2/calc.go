package lesson2

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

func Calculator(stdin io.Reader) (float64, error) {
	var (
		a, b, res float64
		op        string
		err       error
	)

	reader := bufio.NewReader(stdin)

	fmt.Print("Введите первое число: ")
	_, err = fmt.Fscan(reader, &a)
	if err != nil {
		return 0, fmt.Errorf("ой, Ошибка чтения первого числа")
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Fscan(reader, &b)
	if err != nil {
		return 0, fmt.Errorf("ой, Ошибка чтения второго числа")
	}

	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^): ")
	_, err = fmt.Fscan(reader, &op)
	if err != nil {
		return 0, fmt.Errorf("ой, Ошибка чтения арифмметической операции")
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
			return 0, fmt.Errorf("ой, Деление на ноль невозможно")
		}
		res = a / b
	case "^":
		res = math.Pow(a, b)
	default:
		return 0, fmt.Errorf("ой, Операция выбрана неверно")
	}
	return res, nil
}
