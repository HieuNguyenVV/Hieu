// Declare a struct type to maintain information about a user. Declare a function
// that creates value of and returns pointers of this type and an error value. Call
// this function from main and display the value.
//
// Make a second call to your function but this time ignore the value and just test
// the error value.

package main

import (
	"errors"
	"fmt"
)

// Add imports.

// Declare a type named user.
type User struct {
	Name string
	age  int
}

// Declare a function that creates user type values and returns a pointer
// to that value and an error value of nil.
func funcName(user User) (*User, error) /* (pointer return arg, error return arg) */ {
	//ptr := new(User)
	user.Name = "Hieu"
	user.age = 70
	if user.age > 60 {
		return &user, errors.New("He is so old")
	} else {
		return &user, nil
	}
}

// Create a value of type user and return the proper values.

func main() {

	// Use the function to create a value of type user. Check
	// the error being returned.
	user := new(User)
	arr, err := funcName(*user)
	// Display the value that the pointer points to.
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	fmt.Println(*&arr.Name)
	// Call the function again and just check the error.
	_, err = funcName(*user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
