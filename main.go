package main

import (
	"fmt"
	l1 "gb_lvl_1/lesson1"
	l2 "gb_lvl_1/lesson2"
)

func main() {
	// LESSON 1

	fmt.Println(l1.HelloWorld("RU"))
	fmt.Println(l1.HelloWorld("EN"))
	fmt.Println(l1.HelloWorld("ZN"))

	// LESSON 2

	l2.Calculator()

	fmt.Println()
	for _, num := range l2.SimpleNumbersV1(1000) {
		fmt.Print(num, " ")
	}

	fmt.Println()
	for _, num := range l2.SimpleNumbersV2(100) {
		fmt.Print(num, " ")
	}
}
