package main

import "fmt"

func Solution1(A, B []int) bool {
	if len(A) != len(B) {
		return false
	}
	for i := 0; i < len(A)-1; i++ {
		if A[i] == A[i+1] || B[i] == B[i+1] {
			return false
		}
	}
	for i := 0; i < len(A)-1; i++ {
		if A[len(A)-1] != B[0] {
			return false
		}
		if A[i] != B[i+1] {
			return false
		}
	}
	return true
	// if len(A) != len(B) || A[len(A)-1] != B[0] {
	// 	return false
	// }

	// for i := 0; i < len(A)-1; i++ {
	// 	if A[i] == A[i+1] || B[i] == B[i+1] {
	// 		return false
	// 	}
	// }

	// for i := 0; i < len(A); i++ {
	// 	if A[i] == B[i] {
	// 		return false
	// 	}
	// }

	// return true

}
func main() {
	var A = []int{1, 3, 2, 4}
	var B = []int{4, 1, 3, 2}
	// var A = []int{1, 2, 3, 4}
	// var B = []int{2, 1, 4, 3}
	//
	// var A = []int{1, 2, 1}
	// var B = []int{2, 3, 3}
	fmt.Println(Solution1(A, B))
}
