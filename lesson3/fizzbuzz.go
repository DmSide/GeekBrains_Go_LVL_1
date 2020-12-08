package lesson3

import (
	"fmt"
)

func FuzzBuzz() {
	fmt.Println("Hello, playground")
	for i := 1; i < 100; i++ {
		if i%5 != 0 && i%3 != 0 {
			fmt.Print(i)
		} else if i%3 == 0 {
			fmt.Print("Fizz")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		}

		fmt.Println()
	}
}
