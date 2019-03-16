package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	email     string
	age       int
	gender    string
}

// Greeting method that doesn't change anything (value receiver)
func (p Person) greeting(breakGDPR bool) string {
	result := "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old"

	if breakGDPR {
		result += ", you can reach me at " + p.email + "."
	} else {
		result += "."
	}

	return result
}

// Method that actually modifies properties in the struct (pointer receiver)
func (p *Person) getMarried(spouseName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseName
	}
}

func main() {
	amanda := Person{firstName: "Amanda", lastName: "Amandatis", email: "am@ma.com", gender: "f", age: 22}
	john := Person{"John", "Snow", "youknow@nothing.com", 35, "m"}

	// fmt.Println(amanda)
	// fmt.Println(john)
	fmt.Println(amanda.greeting(false))
	fmt.Println(john.greeting(true))
	john.getMarried("Amandatis")
	fmt.Println(john.greeting(true))
	amanda.getMarried("Muflo")
	fmt.Println(amanda.greeting(true))
}
