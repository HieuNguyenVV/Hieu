package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Declare an error variable named ErrHexFileFormat using the New
var (
	ErrHexFileFormat error = errors.New("Not the format of the hex file")
)

func main() {
	FilePath := "nrf52832_xxaa.hex"
	data, err := ioutil.ReadFile(FilePath)
	if err != nil {
		panic(err)
	}
	err = CheckFileFormat(data)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("file is format .hex")
	}
	OutPut(data)
	// //Cheksum:if error break
	err = CheckSum(data)
	if err != nil {
		return
	}
	data1 := ChangeData(data)
	// write  data to file.bin
	file, err := os.Create("data.bin")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var bin_buf bytes.Buffer
	binary.Write(&bin_buf, binary.BigEndian, data1)
	WriteBytetoFileBin(file, bin_buf.Bytes())
	//Read file .bin and display
	fileParth2 := "data.bin"
	databin, err := ioutil.ReadFile(fileParth2)
	if err != nil {
		panic(err)
	}
	OutPut(databin)
}

//Change data to hex data
func ChangeData(k []byte) []byte {
	data := string(k)
	dataofline := strings.Split(data, "\n")
	rawData := []string{}
	for _, line := range dataofline {
		if len(line) == 0 {
			break
		}
		l, _ := strconv.ParseInt(line[1:3], 16, 64)
		rawData = append(rawData, line[9:9+l*2])
	}
	//noi chuoi
	data1 := strings.Join(rawData, "")
	//data hex
	dataHex, _ := hex.DecodeString(data1)
	return dataHex
}

// writing byte data to file bin
func WriteBytetoFileBin(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}

//Xuat du lieu ra man hinh
func OutPut(data []byte) {
	//tao slice chua ma data ascii cua data hex
	var data1 = make([]byte, len(data))
	_ = copy(data1, data)
	for index, value := range data1 {
		if value <= 0x0f {
			data1[index] = 0x2e
		}
	}
	//adress fist=0x0
	//0x0+0x10
	address := 0x0
	for i := 2; i <= len(data); i += 16 {
		if i <= len(data)-16 {
			fmt.Printf("%08x % X %s", address, data[i:i+16], data1[i:i+16])
			fmt.Println()
		} else {
			//truong hop data empty
			empty := strings.Repeat("FF ", (len(data) - i))
			fmt.Printf("%08[1]x % [2]X %[3]s%[4]s", address, data[i:], empty, data1[i:])
			fmt.Println()
		}
		address += 0x10
	}
}

//checsum ca file
func CheckSum(data []byte) error {
	//Convert Hex data to String data
	dataString := string(data)
	//split string through '/n'
	dataOfLine := strings.Split(dataString, "\n")
	for _, dataOfline1 := range dataOfLine {
		dataOfline1 = strings.ToUpper(dataOfline1)
		//TrimSpace returns a slice of the string s, with all leading and trailing white space removed
		dataOfline1 = strings.TrimSpace(dataOfline1)
		if len(dataOfline1) == 0 {
			break
		}
		cs := Checksumofline(dataOfline1)
		checksum1, _ := strconv.ParseInt(dataOfline1[len(dataOfline1)-2:], 16, 32)
		checksum2, _ := strconv.ParseInt(cs, 16, 32)
		if checksum1 != checksum2 {
			return ErrHexFileFormat
		}
	}
	return nil
}

//Funtion check format file
func CheckFileFormat(data []byte) error {
	//Convert Hex data to String data
	dataString := string(data)
	//split string through '/n'
	dataOfLine := strings.Split(dataString, "\n")
	for _, dataOfline1 := range dataOfLine {
		dataOfline1 = strings.ToUpper(dataOfline1)
		//TrimSpace returns a slice of the string s, with all leading and trailing white space removed
		dataOfline1 = strings.TrimSpace(dataOfline1)
		//check empty line
		if len(dataOfline1) == 0 {
			break
		}
		//Check start byte
		if dataOfline1[:1] != ":" {
			return ErrHexFileFormat
		}
		//Check length of line
		if len(dataOfline1) < 11 {
			return ErrHexFileFormat
		}
		//Check length of data
		lenght, err := strconv.ParseInt(dataOfline1[1:3], 16, 64)
		//fmt.Println(lenght)
		if err != nil {
			panic(err)
		}
		//fmt.Println(len(dataOfline1[9:len(dataOfline1)-2]) / 2)
		//fmt.Println(int64(len(dataOfline1[9:len(dataOfline1)-2]) / 2))
		if lenght != int64(len(dataOfline1[9:len(dataOfline1)-2])/2) {
			return ErrHexFileFormat
		}
		cs := Checksumofline(dataOfline1)
		checksum1, _ := strconv.ParseInt(dataOfline1[len(dataOfline1)-2:], 16, 32)
		checksum2, _ := strconv.ParseInt(cs, 16, 32)
		if checksum1 != checksum2 {
			return ErrHexFileFormat
		}
	}
	return nil
}

//checksum tung dong
func Checksumofline(data string) string {
	data1, err := hex.DecodeString(data[1:])
	if err != nil {
		panic(err)
	}
	sum := data1[0]
	for _, b := range data1[1 : len(data1)-1] {
		sum = sum + b
	}
	checsum := 0x01 + ^sum
	return fmt.Sprintf("%x", checsum)
}
