// Write a program that inspects a user's name and greets them in a certain way
// if they are on a list or in a different way if they are not. Also look at
// the user's age and tell them some special secret if they are old enough to
// know it.

package main

import "fmt"

//Add imports

func main() {
	// Change these values and rerun your program.
	name := "Hieu"
	age := 23
	list_name := []string{"Hieu", "Nguyen", "Van"}
	// If the user's name is on a special list then give them a secret greeting.
	for i := 0; i < len(list_name); i++ {
		if name == list_name[i] {
			fmt.Println("Hello Hieu")
		}
	}

	// If the user is old enough then tell them a secret.
	if age < 50 {
		fmt.Println("You are young")
	} else if age >= 50 {
		fmt.Println("You are still young")
	}
}
