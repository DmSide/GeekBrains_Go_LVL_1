package lesson1

func HelloWorld(country string) string {
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
