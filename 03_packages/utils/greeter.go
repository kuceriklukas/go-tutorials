package utils

import "fmt"

const greeting = "Hello, "

func GreetMe(name string) string {
	return greeting + name
}

func PrintGreetMe(name string) {
	fmt.Println(greeting + name)
}
