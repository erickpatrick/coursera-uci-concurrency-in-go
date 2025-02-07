package main

import "fmt"

// this program has 2 goroutines which try to increment the number variable by
// different values, +1 and +2 respectively. As each goroutine works like a
// thread, they get to be executed in an non-deterministic order. As such
// if the order of the goroutines commands is set to be after the last
// statement of the program, they will not be printed out nor will
// change the values of the variable number, thus causing a race
// condition between assigning new values to the variable and
// printing its current value to the consumer of the program
func main() {
	var number int = 0

	go func() {
		number += 1
		fmt.Println("first number ", number)
	}()

	fmt.Println("second number ", number)

	go func() {
		number += 2
		fmt.Println("third number ", number)
	}()

	fmt.Println("last number ", number)
}
