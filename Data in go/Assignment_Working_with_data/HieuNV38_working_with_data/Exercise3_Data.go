package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex = sync.RWMutex{}

type Tweet struct {
	name string
	sp   string
}

func Check1(str1 string, str string) {
	if strings.Contains(str, "Golang") || strings.Contains(str, "golang") {
		//fmt.Printf(st1)
		fmt.Println(str1, "        Twite about golan")
	} else {
		//fmt.Printf(st1)
		fmt.Println(str1, "         Does not Twite about golan")
	}
	wg.Done()
}
func Check2(str []Tweet) {
	for i := 0; i < len(str); i++ {
		if strings.Contains(str[i].sp, "Golang") || strings.Contains(str[i].sp, "golang") {
			fmt.Printf(str[i].name)
			fmt.Println("        Twite about golan")
		} else {
			fmt.Printf(str[i].name)
			fmt.Println("         Does not Twite about golan")
		}

	}
}
func main() {

	var str = []Tweet{
		{
			"anhlv",
			"#golang top tip: if your unit tes import any other package you wrote, including themselves, they're not unit test",
		}, {
			"beertoco",
			"Backend developer, doing frontend featuring the eternal struge of centering something. #codi",
		}, {
			"zz00",
			"Re: Polarity of Golang in China: My thinking nowadays is that it had a lot to do with this book and author https://github.com/astaxie/build-web-application-with-gola",
		}, {
			"Lm",
			"Loong forward to the #gopher meetup in Hsinchu tonight with @ironze",
		}, {
			"vampirewalk6",
			"I just ote a golang slack bot! It reports the state of github repository. #Slack #gola",
		}, {
			"tndoan26",
			"I just wrota golang slack bot! It reports the state of github repository. #Slack #gola",
		}, {
			"aloalo12",
			"Iust wrote a java slack bot! It reports the state of github repository. #Sla",
		}, {
			"tlatoi12",
			"lang!",
		},
	}
	t := time.Now()
	for i := 0; i < len(str); i++ {
		wg.Add(1)
		go Check1(str[i].name, str[i].sp)
	}
	wg.Wait()
	t1 := time.Now()
	fmt.Println(t1.Sub(t))

	t4 := time.Now()
	Check2(str)
	t2 := time.Now()
	fmt.Println(t2.Sub(t4))
}
