package lesson3

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Print("Fizz")
		} else if i%3 != 0 && i%5 == 0 {
			fmt.Print("Buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz")
		} else {
			fmt.Print(i)
		}

		fmt.Println()
	}
}
