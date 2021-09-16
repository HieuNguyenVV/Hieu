// Declare an interface named speaker with a method named speak. Declare a struct
// named english that represents a person who speaks english and declare a struct named
// france for someone who speaks france. Implement the speaker interface for each
// struct using a value receiver and these literal strings "Hello World" and "Bonjour le monde".
// Declare a variable of type speaker and assign the address of a value of type english
// and call the method. Do it again for a value of type france.
//
// Add a new function named sayHello that accepts a value of type speaker.
// Implement that function to call the speak method on the interface value. Then create
// new values of each type and use the function.
package main

import "fmt"

// Declare the speaker interface with a single method called speak.
type Speaker interface {
	NameSpeak()
}

// Declare an empty struct type named english.
type Enlish struct {
	name string
}

// Declare a method named speak for the english type
// using a value receiver. "Hello World"

func (ptr1 Enlish) NameSpeak() {
	fmt.Println("Hello word")
}

// Declare an empty struct type named france.
type France struct {
	name string
}

// Declare a method named speak for the france type
// using a pointer receiver. "Bonjour le monde"
func (ptr2 *France) NameSpeak() {
	fmt.Println("Bonjour le monde")
}

// sayHello accepts values of the speaker type.
func sayHello(sp Speaker /* Declare parameter */) {
	sp.NameSpeak()
	// Call the speak method from the speaker parameter.
}
func main() {
	// Declare a variable of the interface speaker type
	// set to its zero value.
	var a Speaker
	// Declare a variable of type english
	en := Enlish{}
	// Assign the english value to the speaker variable.
	a = en
	// Call the speak method against the speaker variable.
	a.NameSpeak()
	// Declare a variable of type france.
	fr := France{}
	// Assign the france pointer to the speaker variable.
	a = &fr
	// Call the speak method against the speaker variable.
	a.NameSpeak()
	// Call the sayHello function with new values and pointers
	// of english and france.
	sayHello(en)
	sayHello(&fr)
}
