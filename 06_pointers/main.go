package main

import "fmt"

func main() {
	fmt.Println("Pointeeers")

	a := 6
	b := &a // the memory address of "a"

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	*b = 7 // change the value of "b" to 7 - in other words, change the a's value to 7

	fmt.Println(a, b)

	c := *b
	fmt.Printf("%T\n", c)
}
