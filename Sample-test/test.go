package main

import "fmt"

type Point struct {
	a int
	b int
}

func initArr(n int) []int {
	var tem []int
	for i := 0; i < n; i++ {
		tem = append(tem, i+1)
	}
	return tem
}

func checkPath(A, B []int, n int) bool {
	var res bool
	//check given N and length of path can be created
	if n > len(A) {
		res = false
	}

	var k = initArr(n)

	var p []Point
	for i := 0; i < len(A); i++ {
		p = append(p, Point{A[i], B[i]})
	}
	fmt.Println(len(p))
	var lena = len(A)
	//for loop process the 1-> n-1 node.
	//we process last node in other logic
	if ((p[lena-1].a == n) && (p[lena-1].b == n-1)) || ((p[lena-1].a == n-1) && (p[lena-1].b == n)) {
		res = true
	} else {
		res = false
	}
	for j := 0; j < len(k)-2; j++ {
		for _, v := range p {
			if ((v.a == j+1) && (v.b == j+2)) || ((v.a == j+2) && (v.b == j+1)) {
				res = true
				break
			} else {
				res = false
			}
		}
	}

	return res

}
func main() {
	//var A = []int{1, 2, 4, 4, 3}
	//var B = []int{2, 3, 1, 3, 1}
	//var A = []int{1, 2, 1, 3}
	//var B = []int{2, 4, 3, 4}
	var A = []int{2, 4, 5, 3}
	var B = []int{3, 5, 6, 4}
	N := 6
	//var A = []int{1, 3}
	//var B = []int{2, 3}
	fmt.Println(checkPath(A, B, N))
}
