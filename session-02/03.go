package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int

	fmt.Print("Enter n: ")
	fmt.Scan(&n)

	a1 := 1
	a2 := 2

	outputString := ""

	if n >= 1 {
		outputString += strconv.Itoa(a1)
		outputString += " "
	}

	if n >= 2 {
		outputString += strconv.Itoa(a2)
		outputString += " "
	}

	if n >= 3 {
		for range n - 2 {
			a3 := a1 + a2
			outputString += strconv.Itoa(a3)
			outputString += " "

			a1 = a2
			a2 = a3
		}
	}

	fmt.Println(outputString)
}
