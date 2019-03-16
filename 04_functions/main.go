package main

import (
	"fmt"
	"strconv"
)

func getAddition(a, b int) int {
	return a + b
}

func main() {
	var result = getAddition(5, 4)

	fmt.Println("muff " + strconv.Itoa(result))
	fmt.Println(result)
}
