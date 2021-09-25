package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution3(str string) int {
	map1 := make([]string, 0)
	map1 = append(map1, str)
	s1 := strings.Split(str, "")
	n := len(s1)

	for j := 0; j < n; j++ {
		for i := 0; i <= 9; i++ {
			s := strconv.Itoa(i)
			if s != s1[j] {
				if j == 0 {
					data := strings.Replace(str, s1[j], s, 1)
					map1 = append(map1, data)
				} else {
					map1 = append(map1, str[:j]+s+str[j+1:])
				}

			}

		}
	}
	count := 0
	for _, value := range map1 {
		number, _ := strconv.Atoi(value)
		if number%3 == 0 {
			count++
		}

	}
	return count
}
func main() {
	str := "022"
	count := Solution3(str)
	fmt.Println(count)

}
