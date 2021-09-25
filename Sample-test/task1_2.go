package main

import "fmt"

func Solution2(arr []int) int {
	min := arr[0]
	max := arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	fmt.Println(min, max)
	n := (max + min) / 2
	fmt.Println(n)
	count := 0
	for _, value := range arr {
		if value > n {
			count = count + value - n
		}
		if value < n {
			count = count + n - value
		}
		if value == n {
			continue
		}
	}
	return count
}
func main() {
	var a = []int{3, 1, 4, 1}
	//var a = []int{3, 2, 1, 1, 2, 3, 1}
	fmt.Println("x = ", Solution2(a))

}
