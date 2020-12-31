package lesson1

import "fmt"

func ExampleHelloWorld() {
	hw, _ := HelloWorld("RU")
	fmt.Printf("%s", hw)
	// Output: Привет, мир!
}

// TODO: How to create more tests without "Example refers to unknown field or method: HelloWorld.EN"?
func ExampleHelloWorld_EN() {
	hw, _ := HelloWorld("EN")
	fmt.Printf("%s", hw)
	// Output: Hello world!
}

// TODO: How to create more tests without "Example refers to unknown field or method: HelloWorld.ZN"?
func ExampleHelloWorld_ZN() {
	hw, _ := HelloWorld("ZN")
	fmt.Printf("%s", hw)
	// Output: 你好，世界!
}
