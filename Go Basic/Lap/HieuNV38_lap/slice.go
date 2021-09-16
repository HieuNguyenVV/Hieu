// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

import "fmt"

// Add imports.

func main() {

	// Declare a nil slice of integers.
	var slice1 []int
	// Append numbers to the slice.
	for i := 0; i < 10; i++ {
		slice1 = append(slice1, i)
	}
	// Display each value in the slice.
	fmt.Println("Value of slice1:", slice1)
	// Declare a slice of strings and populate the slice with names.
	slice2 := []string{"Nguyen", "Van", "Hieu", "FPT", "SW"}
	// Display each index position and slice value.
	fmt.Println("Value of slice2:", slice2)
	// Take a slice of index 1 and 2 of the slice of strings.
	slice3 := slice2[1:2]
	slice4 := slice2[1:]
	slice5 := slice2[:3]
	slice6 := slice2[:]
	// Display each index position and slice values for the new slice.
	fmt.Println("Value of slice3:", slice3)
	fmt.Println("Value of slice4:", slice4)
	fmt.Println("Value of slice5:", slice5)
	fmt.Println("Value of slice6:", slice6)
	//Length og slice3
	fmt.Println("Length of slice3:", len(slice3))
	fmt.Println("Length of slice4:", len(slice4))
	fmt.Println("Length of slice5:", len(slice5))
	fmt.Println("Length of slice6:", len(slice6))
	//Capacity of slice3
	fmt.Println("Capacity of slice3:", cap(slice3))
	fmt.Println("Capacity of slice4:", cap(slice4))
	fmt.Println("Capacity of slice5:", cap(slice5))
	fmt.Println("Capacity of slice6:", cap(slice6))
}
