package main

import "fmt"

func Solution(array []int) int {
	var count int = 0
	// for i := 0; i < len(array); i++ {
	// 	for j := 0; j < len(array); j++ {
	// 		if array[i] == array[j]-1 {
	// 			count++
	// 			break
	// 		}
	// 	}
	// }
	var i, num int
	map1 := make(map[int]int)
	for i, num = range array {

		if _, ok := map1[num+1]; ok {
			count++
		} else {
			map1[num] = i
		}

	}
	if _, ok := map1[array[0]+1]; ok {
		count++
	}
	return count

}
func main() {
	//var A = []int{3, 4, 3, 0, 2, 2, 3, 0, 0}
	var A = []int{3, 4, 4, 3, 1, 0}
	count := Solution(A)
	fmt.Println(count)
}
