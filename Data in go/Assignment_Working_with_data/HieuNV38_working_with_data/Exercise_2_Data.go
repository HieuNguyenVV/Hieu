package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"strings"
// 	"sync"
// )

// // var wg sync.WaitGroup
// // var count int = 1
// // var mutex = sync.RWMutex{}

// type NameFile struct {
// 	name string
// }
// type Data struct {
// 	data []string
// 	//name []string
// }

// func ReadFiles(filename []NameFile) Data {
// 	output := make(chan Data)
// 	input := make(chan string)
// 	var wg sync.WaitGroup
// 	defer close(output)
// 	for i := 0; i < len(filename); i++ {
// 		wg.Add(1)
// 		go ReadFile(filename[i], input)
// 	}
// 	go Output(input, output, &wg)
// 	wg.Wait()
// 	close(input)
// 	return <-output
// }
// func Output(input chan string, output chan Data, wg *sync.WaitGroup) {
// 	var data Data
// 	for resutl := range input {
// 		data.data = append(data.data, resutl)

// 		wg.Done()
// 	}

// 	output <- data
// }
// func Writefiles( /*filepath []NameFile,*/ data Data) {
// 	var wg sync.WaitGroup
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go Writefile( /*filepath[i],*/ data.data[i], &wg)
// 	}
// 	wg.Wait()
// }
// func Writefile( /*filepath NameFile,*/ data string, wg *sync.WaitGroup) {
// 	dataofline := strings.Split(data, " ")
// 	name := "copy" + dataofline[0]
// 	file, err := os.Create(name)

// 	if err != nil {
// 		//handle the error here return
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	var number int = 0
// 	var data1 string
// 	for i := 1; i < len(dataofline); i++ {
// 		data1 = data1 + " " + dataofline[i]
// 	}
// 	//fmt.Println(data1)
// 	//file.WriteString(data)
// 	for i := 1; i < len(data1); i++ {
// 		//fmt.Println(data[i])
// 		number++
// 		a := data1[i]
// 		b := string(a)
// 		//fmt.Println(a)
// 		//fmt.Println(b)
// 		file.WriteString(b)
// 		x := 100 * number / (len(data1) - 1)
// 		fmt.Printf("File %s writing x = %v%%\n", name, x)
// 	}

// 	wg.Done()

// }

// func ReadFile(filepath NameFile, input chan string) {
// 	file, err := os.Open(filepath.name)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fileinfo, err := file.Stat()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	filesize := int(fileinfo.Size())
// 	number := filesize
// 	var k int = 0
// 	var data string
// 	for i := 1; i <= number; i = i + 1 {
// 		byteSlice := make([]byte, i)
// 		minBytes := 1
// 		numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		k = k + numBytesRead
// 		x := 100 * k / filesize
// 		fmt.Printf("File %s reading x = %v%%\n", filepath.name, x)
// 		a := string(byteSlice)
// 		//fmt.Println(byteSlice)
// 		data = data + a
// 		//fmt.Println(a)
// 		if k == filesize {
// 			break
// 		}
// 	}
// 	//fmt.Println(data)
// 	//time.Sleep(time.Second * 10)
// 	//wg.Done()
// 	data = filepath.name + " " + data
// 	input <- data
// }

// func main() {
// 	nameFile := []NameFile{
// 		NameFile{name: "a1.txt"},
// 		NameFile{name: "a2.txt"},
// 		NameFile{name: "a3.txt"},
// 		NameFile{name: "a4.txt"},
// 		NameFile{name: "a5.txt"},
// 		NameFile{name: "a6.txt"},
// 		NameFile{name: "a7.txt"},
// 		NameFile{name: "a8.txt"},
// 		NameFile{name: "a9.txt"},
// 		NameFile{name: "a10.txt"},
// 	}

// 	reuslt := ReadFiles(nameFile)
// 	Writefiles( /*nameFile2,*/ reuslt)

// }
