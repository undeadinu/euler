package main

import "fmt"

// fibo calculates the sum of even-valued terms in a fibonacci sequence until some number
func fibo(limit int) int {
	// variables:  a, b = 1, c, sum = 0
	a, b, c, sum := 1, 1, 0, 0

	for {

		if c > limit {
			break
		}

		// if number is even, add it to the sum
		if c%2 == 0 {
			sum += c
		}

		// adds two previous values to c
		c = a + b

		// assigns c to a, a to b, enabling the loop to progress
		b = a
		a = c
	}
	return sum
}

func main() {
	fmt.Println(fibo(4000000))
}
