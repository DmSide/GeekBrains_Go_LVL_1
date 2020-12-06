package main

import "fmt"
import l3 "gb_lvl_1/lesson3"

func helloWorld(country string) string {
	switch country {
	case "RU":
		return "Привет, мир!"
	case "EN":
		return "Hello world!"
	case "ZN":
		return "你好，世界!"
	default:
		return "I don't know the entered country. Country must contain only 2 letters"

	}
}

func main() {
	fmt.Println(helloWorld("RU"))
	fmt.Println(helloWorld("EN"))
	fmt.Println(helloWorld("ZN"))

	// LESSON 3

	arr := []int{10, 4, 8, 6, 2, 5, 3, 7, 9}
	a := l3.SortInsert(arr)
	fmt.Println(a)
}
