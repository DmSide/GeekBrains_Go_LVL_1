package main

import (
	"fmt"
	l1 "gb_lvl_1/lesson1"
	l2 "gb_lvl_1/lesson2"
	l3 "gb_lvl_1/lesson3"
	l4 "gb_lvl_1/lesson4"
)

func main() {
	var (
		err error
		val string
	)
	// LESSON 1
	countries := []string{"RU", "EN", "ZN", "UNKNOWN COUNTRY"}
	for _, country := range countries {
		val, err = l1.HelloWorld(country)
		if err != nil {
			fmt.Printf("Error occurs: %v", err)
		}
		fmt.Println(val)
	}

	// LESSON 2

	if false {
		l2.Calculator()
	}

	fmt.Println()
	for _, num := range l2.SimpleNumbersV1(1000) {
		fmt.Print(num, " ")
	}

	fmt.Println()
	for _, num := range l2.SimpleNumbersV2(100) {
		fmt.Print(num, " ")
	}

	// LESSON 3

	fmt.Println()
	l3.FuzzBuzz()

	arr1 := []int{10, 4, 8, 6, 2, 5, 3, 7, 9}
	a1 := l3.SortBubble(arr1)
	fmt.Println(a1)

	arr2 := []int{10, 4, 8, 6, 2, 5, 3, 7, 9}
	a2 := l3.SortInsert(arr2)
	fmt.Println(a2)

	// LESSON 4
	l4.DeferMagic()

	fs := l4.FibStruct{}
	fs.Init()
	fs.Fib(10)
	fs.Print()
}
