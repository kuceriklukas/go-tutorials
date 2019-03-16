package main

import "fmt"

const myConst = "never change"

var isUnicorn = true

func main() {
	name := "Raindeer"
	number := 18
	var myNumber int16 = 45

	// myConst = "asd"

	fmt.Println(name, number)
	fmt.Printf("%T\n", myNumber)
	fmt.Println(myConst)
	fmt.Println(isUnicorn)
}
