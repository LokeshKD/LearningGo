// Program to compute displacement after time "t" seconds
// Formula  s' = 1/2.a.t^2 + v.t + s
// Takes 3 float64 inputs for acceleration, inital velocity and initial displacement.
// given a time in seconds (again, input from a user), this computes the final displacement.

package main

import "fmt"

func main() {
	// Let the User know the intent of the program, at runtime.
	fmt.Println("To compute the final distance, we should know the initial values")
	fmt.Println("of acceleration, velocity and distance.\n")

	var acc, vel, dis, tim float64
	// Get the User inputs.
	fmt.Printf("Accelaration:")
	fmt.Scan(&acc)
	fmt.Printf("Velocity:")
	fmt.Scan(&vel)
	fmt.Printf("Initial Distance:")
	fmt.Scan(&dis)

	compute := GenDisplaceFn(acc, vel, dis)

	// Get the time displacement from User
	fmt.Printf("Time in seconds:")
	fmt.Scan(&tim)

	// Print the displacement from the initial point with given values.
	fmt.Printf("\nDisplacement after %f seconds is %f\n", tim, compute(tim))
}

func GenDisplaceFn(acc, vel, dis float64) func (float64) float64 {
	cal := func(tim float64) float64 {
		return 0.5 * acc * tim * tim + vel * tim + dis
	}
	return cal
}
