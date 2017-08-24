package main

import "fmt"

// add finds the sum of all multiples of 3 or 5 below 1000
func add() int {
	sum := 0

	for i := 3; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(add())
}
