package main

import (
	"fmt"
	"rsc.io/quote"
	"example.com/greetings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(quote.Go())

	fmt.Println(Hello("osama"))
	fmt.Println(Hello2("osama"))
	fmt.Println(greetings.HelloGreet("osama"))
}

func Hello(name string) string {
	message:= fmt.Sprintf("This is %v", name)
	return message
}

func Hello2(name string) string {
	var message string
	message = fmt.Sprintf("This is %v", name)
	return message
}