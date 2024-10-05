package main

import (
	"fmt"
)

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	var a, vo, so, t float64

	// Prompt user for input
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)
	
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&vo)
	
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&so)
	
	fn := GenDisplaceFn(a, vo, so)

	fmt.Print("Enter time: ")
	fmt.Scan(&t)

	displacement := fn(t)
	fmt.Printf("Displacement after %.2f seconds: %.2f\n", t, displacement)
}
