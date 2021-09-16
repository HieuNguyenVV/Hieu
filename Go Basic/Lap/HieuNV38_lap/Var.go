// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// Add imports

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var a, b, c int
	// Display the value of those variables.
	fmt.Println("Value of a,b,c :", a, b, c)
	// Declare variables and initialize.
	var a1, b1, c1 string = "Nguyen", "Van", "Hieu"
	// Using the short variable declaration operator.
	x, y, z := "Nguyen", "Van", "Hieu"
	// Display the value of those variables.
	fmt.Println("Value of a1,b1,c1 :", a1, b1, c1)
	fmt.Println("Value of x,y,z :", x, y, z)
	// Perform a type conversion.
	var pi float64 = 3.14
	var float float32 = float32(pi)
	// Display the value of that variable.
	fmt.Println("Value of float :", float)
}
