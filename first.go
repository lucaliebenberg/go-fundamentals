package main

import "fmt"

func main() {
	message := "Hello, world!"
	price := 30
	fmt.Println("Hello world!")

	print(message, price)

	stateTax, _ := calculate(100)
	fmt.Println(stateTax)

}
