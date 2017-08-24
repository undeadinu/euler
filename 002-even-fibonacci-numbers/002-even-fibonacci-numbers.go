package main

import (
	"fmt"
	"time"
)

// timeTrack calculates time a function takes to run in seconds
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start).Seconds()
	fmt.Printf("%s took %.8f seconds", name, elapsed)
	fmt.Println()
}

// fibo calculates the sum of even-valued terms in a fibonacci sequence until some number
func fibo(limit int) int {
	a, b, c, sum := 1, 1, 0, 0

	for {

		// break once we reach our limit
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
	fmt.Println(fibo(40000))
}
