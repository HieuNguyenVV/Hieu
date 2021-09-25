package main

import "fmt"

func Solution(str string) int {
	allStrings := make([]string, 0)
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			allStrings = append(allStrings, str[i:j+1])
		}
	}
	fmt.Println(allStrings)
	mapString := make(map[string]int)
	for _, value := range allStrings {
		if value1, ok := mapString[value]; ok {
			mapString[value] = value1 + 1
		} else {
			mapString[value] = 1
		}
	}
	fmt.Println(mapString)
	mapString1 := make(map[int]int)
	for key, value := range mapString {
		if value == 1 {
			mapString1[len(key)] = 1
		}
	}
	fmt.Println(mapString1)
	count := 2
	for i := 0; i < len(str); i++ {
		if value1, ok := mapString1[i]; ok && value1 == 1 {
			return i
		}
	}
	return count
}
func main() {
	str := "aabbbabaa"
	//str := "zyzyzyz"
	//str := "abaaba"
	fmt.Println(Solution(str))
}
