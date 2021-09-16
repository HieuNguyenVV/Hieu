// Create two error variables, one called ErrInvalidValue and the other called ErrAmountTooLarge.
// Provide the static message for each variable. Then write a function called checkAmount that accepts a float64 type value and returns an error value.
// Check the value for zero and if it is, return the ErrInvalidValue. Check the value for greater than $1,000 and if it is, return the ErrAmountTooLarge.
// Write a main function to call the checkAmount function and check the return error value. Display a proper message to the screen.

/* Template */
package main

//Add imports.
import (
	"errors"
	"fmt"
)

var (
	ErrInvalidValue error = errors.New("bang 0")

	// Declare an error variable named ErrInvalidValue using the New
	// function from the errors package.

	// Declare an error variable named ErrAmountTooLarge using the New
	// function from the errors package.
	ErrAmountTooLarge error = errors.New("lon hon 1000")
)

// Declare a function named checkAmount that accepts a value of
// type float64 and returns an error interface value.
func checkAmount(a float64 /* parameter */) error /* return arg */ {
	if a == 0 {
		return ErrInvalidValue
	} else if a > 1000 {
		return ErrAmountTooLarge
	} else {
		return nil
	}
	// Is the parameter equal to zero. If so then return
	// the error variable.

	// Is the parameter greater than 1000. If so then return
	// the other error variable.

	// Return nil for the error value.

}

func main() {

	// Call the checkAmount function and check the error. Then
	// use a switch/case to compare the error with each variable.
	// Add a default case. Return if there is an error.
	a := 10.1
	err := checkAmount(a)
	switch err {
	case ErrInvalidValue:
		fmt.Println(err.Error())
		break
	case ErrAmountTooLarge:
		fmt.Println(err.Error())
		break
	default:
		fmt.Println("error value")
	}

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// Display everything is good.
}
