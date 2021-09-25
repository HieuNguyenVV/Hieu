package main

import "fmt"

func Solution(n, k int) int {
	if n == 1 && k > 1 {
		return -1
	}
	if n == 1 && k == 1 {
		return 1
	}
	sum := n
	count := 1
	for i := n - 1; i > 0; i-- {
		if sum+i > k {
			continue
		}
		if sum+i == k {
			count++
			break
		}
		if sum+i < k {
			count++
			sum = sum + i
		}

	}
	return count
}
func main() {
	count := Solution(5, 10)
	if count == -1 {
		return
	}
	fmt.Println(count)
}
