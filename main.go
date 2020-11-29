package main

import "fmt"

//func helloWorld(страна string ) string{	// it works )
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
	// привет := "Привет" // Just for fun
	fmt.Println(helloWorld("RU"))
	fmt.Println(helloWorld("EN"))
	fmt.Println(helloWorld("ZN"))
}
