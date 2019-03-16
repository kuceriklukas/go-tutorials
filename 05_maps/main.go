package main

import "fmt"

func main() {
	var people = make(map[string]string)

	people["Fufo"] = "fuffel@f.com"
	people["Mufo"] = "mu@f.com"
	people["Adds"] = "asd@asws.com"

	fmt.Println(people)

	persons := map[string]string{"as-85": "Lukas Muflon", "ers-85": "Tafi Dafi", "api-6": "Armadillo Ezechielus"}

	fmt.Println(persons)

	for id, name := range persons {
		fmt.Printf("%s \t\t %s\n", id, name)
	}
}
