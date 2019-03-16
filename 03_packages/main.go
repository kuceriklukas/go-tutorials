package main

import (
	"fmt"
	"math"

	"github.com/kuceriklukas/tutorial1/03_packages/utils"
)

func main() {
	var coolFloat = 32.45647
	fmt.Println(math.Ceil(coolFloat))
	fmt.Println(math.Floor(coolFloat))

	var coolGreeting string = utils.GreetMe("Fluffy Unicorn")
	fmt.Println(coolGreeting)

	utils.PrintGreetMe("Girafful")
}
