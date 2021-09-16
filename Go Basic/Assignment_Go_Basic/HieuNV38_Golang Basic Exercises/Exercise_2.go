package main

// import (
// 	"fmt"
// 	"regexp"
// 	"strconv"
// )

// func main() {
// 	var a int = 100
// 	b := countBit1(a)
// 	fmt.Println("so bit 1 :", b)

// 	var Number uint = 100
// 	Number1 := SetBit(Number, 5)
// 	fmt.Println("bien a sau khi doi bit thu 5 len 1 :", Number1)
// 	fmt.Printf("%b", Number1)
// 	fmt.Println()

// 	Number2 := ClearBit(100, 6)
// 	fmt.Println("bien a sau khi doi bit thu 6 len 0 :", Number2)
// 	fmt.Printf("%b", Number2)
// 	fmt.Println()

// 	Number3 := Togger(100, 3)
// 	fmt.Println("bien a sau khi dao bit thu 3 ve 0 :", Number3)
// 	fmt.Printf("%b", Number3)
// 	fmt.Println()
// 	fmt.Printf("%b", 13)
// 	fmt.Println()

// 	Number4 := Daobit(13)
// 	fmt.Println("bien a sau khi thay doi vi tri bit ox1234->ox3412 :", Number4)
// 	fmt.Printf("%b", Number4)
// 	fmt.Println()

// 	//fmt.Printf("%b", Number4)
// 	fmt.Println()
// 	Number5 := Daobit1(Number4)
// 	fmt.Println("bien a sau khi thay doi vi tri bit 0x3412->0x4321 :", Number5)
// 	fmt.Printf("%b", Number5)
// 	fmt.Println()
// }

// //??m s? bit c? gi? tr? 1 v? in ra m?n h?nh
// func countBit1(b int) int {
// 	c := strconv.FormatInt(int64(b), 2)
// 	fmt.Println(c)
// 	re := regexp.MustCompile("1")
// 	oneBit := re.FindAll([]byte(c), -1)
// 	return len(oneBit)
// }

// //Vi?t ch??ng tr?nh thay ??i gi? tr? bit th? 5 t? ph?i qua c?a a v? 1 v? hi?n th? ra m?n h?nh gi? tr? c?a bi?n a.
// //Funtion set bit:0|1=1   1|1=1
// func SetBit(a, b uint) uint {
// 	return (a | (1 << (b - 1)))
// }

// //Vi?t ch??ng tr?nh thay ??i gi? tr? bit th? 4 t? ph?i qua c?a a v? 0 v? hi?n th? ra m?n h?nh gi? tr? c?a bi?n a.
// func ClearBit(a, b uint) uint {
// 	return (a & (a ^ (1 << (b - 1))))
// }

// //Vi?t ch??ng tr?nh ??o gi? tr? bit th? 3 t? ph?i qua c?a a v? 0 v? hi?n th? ra m?n h?nh gi? tr? c?a bi?n a.
// func Togger(a, b uint) uint {
// 	return (a ^ (1 << (b - 1)))
// }

// //Ch? s? d?ng Bitwise ?? thay ??i s? 0x1234 th?nh 0x3412 v? 0x4321.
// func Daobit(b uint) uint {
// 	return (b&0xCC)>>2 | (b&0x33)<<2
// }
// func Daobit1(b uint) uint {
// 	return (b&0xAA)>>1 | (b&0x55)<<1
// }
//func bitwise(a unit16) {

//}
