package main

import "fmt"

func main() {
	// Arrays
	var unicorns [3]string
	unicorns[0] = "Bob"
	unicorns[1] = "Rudolf"
	unicorns[2] = "Ezechiel"

	raindeers := []string{"Coolio", "Sadio", "Funkio", "Samuraio"}

	fmt.Println(unicorns)
	fmt.Println(raindeers)

	fmt.Println(unicorns[1])
	fmt.Println(raindeers[1:3])
	fmt.Println(len(raindeers))
	fmt.Println(len(raindeers[3]))

	// Conditions
	x := 13
	y := 12

	if x > y {
		fmt.Printf("%d is bigger than %d\n", x, y)
	} else {
		fmt.Printf("%d is bigger than %d\n", y, x)
	}

	// Defer
	anotherMain()

	// FizzBuzz
	fizzBuzz()
}

func anotherMain() {
	fmt.Println("Start")

	defer fmt.Println("First!")
	defer fmt.Println("Second!")

	fmt.Println("Done")
}

func fizzBuzz() {
	for i := 0; i <= 100; i++ {
		fmt.Print(i)

		if i%3 == 0 {
			fmt.Print(" Fizz")
		}
		if i%5 == 0 {
			fmt.Print(" Buzz")
		}

		fmt.Print("\n")
	}
}
