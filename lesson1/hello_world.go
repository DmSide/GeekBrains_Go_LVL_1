package lesson1

type HelloWorldError struct{}

func (*HelloWorldError) Error() string {
	// Some additional logic could be here
	return "I don't know the entered country. Country must contain only 2 letters"
}

func HelloWorld(country string) (string, error) {
	switch country {
	case "RU":
		return "Привет, мир!", nil
	case "EN":
		return "Hello world!", nil
	case "ZN":
		return "你好，世界!", nil
	default:
		return "", &HelloWorldError{}
	}
}
