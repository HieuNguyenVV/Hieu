package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func solution(s, t string) int {
	count := 0
	layout := "15:04:05"
	time1, err := time.Parse(layout, s)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(t1)
	time2, err := time.Parse(layout, t)
	if err != nil {
		log.Fatal(err)
	}
	duration := time2.Sub(time1)
	seconds := int(duration.Seconds())
	//fmt.Println(seconds)
	for i := 0; i <= seconds; i++ {
		time_cov := time1.Add(time.Duration(i) * time.Second)
		fmt.Println(time_cov)
		s1 := time_cov.Format(layout)
		fmt.Println(s1)
		map1 := make(map[string]int)
		data := strings.Replace(s1, ":", "", -1)
		//fmt.Println(data)
		newdata := strings.Split(data, "")
		//fmt.Println(newdata)
		for indx, v := range newdata {
			map1[v] = indx
		}
		if len(map1) <= 2 {
			count++
		}
	}
	return count
}
func main() {
	count1 := solution("15:15:00", "15:15:12")
	fmt.Println(count1)
	count2 := solution("22:22:21", "22:22:25")
	fmt.Println(count2)
}
