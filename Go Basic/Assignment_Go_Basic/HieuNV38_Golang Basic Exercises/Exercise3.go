package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// )

// type Input interface {
// 	Nhap()
// 	Xuat()
// 	DiemTB()
// 	PhanLoai()
// }
// type SinhVien struct {
// 	ten      string
// 	gioiting string
// 	tuoi     int
// 	diemtoan float32
// 	diemli   float32
// 	diemhoa  float32
// 	dirmtb   float32
// }

// func (sv *SinhVien) Nhap() {
// 	fmt.Println("Nhap ten")
//var string
// 	fmt.Scanln(&sv.ten)
// 	fmt.Println("Nhap gioi tinh")
// 	fmt.Scanln(&sv.gioiting)
// 	fmt.Println("Nhap tuoi")
// 	fmt.Scanln(&sv.tuoi)
// 	fmt.Println("Nhap diem toan")
// 	fmt.Scanln(&sv.diemtoan)
// 	fmt.Println("Nhap diem li")
// 	fmt.Scanln(&sv.diemli)
// 	fmt.Println("Nhap diem hoa")
// 	fmt.Scanln(&sv.diemhoa)
// }
// func (sv SinhVien) Xuat() {
// 	fmt.Println("Ten", sv.ten)
// 	fmt.Println("gioi tinh", sv.gioiting)
// 	fmt.Println("tuoi", sv.tuoi)
// 	fmt.Println("diem toan", sv.diemtoan)
// 	fmt.Println("diem li", sv.diemli)
// 	fmt.Println("diem hoa", sv.diemhoa)

// }
// func (sv SinhVien) DiemTB() {
// 	sv.dirmtb = (sv.diemtoan + sv.diemli + sv.diemhoa) / 3
// 	fmt.Println("Diem trung binh:", sv.dirmtb)
// }
// func (sv SinhVien) PhanLoai() {
// 	if sv.dirmtb >= 8.5 {
// 		fmt.Println("Gioi")
// 	} else if sv.dirmtb >= 6.0 && sv.dirmtb < 8.5 {
// 		fmt.Println("Kha")
// 	} else if sv.dirmtb < 6.0 && sv.dirmtb >= 4.0 {
// 		fmt.Println("Trung binh")
// 	} else {
// 		fmt.Println("Yeu")
// 	}
// }

// func Sort(sv [100]SinhVien, n int) {
// 	for i := 0; i < n; i++ {
// 		for j := (i + 1); j < n; j++ {
// 			if (sv[i].dirmtb) > (sv[j].dirmtb) {
// 				swap(&sv[i], &sv[j])
// 			}
// 		}
// 	}
// }
// func swap(a *SinhVien, b *SinhVien) {
// 	*a, *b = *b, *a
// }
// func main() {
// 	var n int
// 	fmt.Scanln(&n)
// 	var sv1 [100]SinhVien
// 	for i := 0; i < n; i++ {
// 		sv1[i].Nhap()

// 	}
// 	var chon int
// 	fmt.Println("Nhap lua chon")
// 	fmt.Scanln(&chon)
// 	switch chon {
// 	case 1:
// 		for i := 0; i < n; i++ {
// 			sv1[i].Xuat()
// 		}
// 		break
// 	case 2:
// 		for i := 0; i < n; i++ {
// 			sv1[i].Xuat()
// 			sv1[i].DiemTB()
// 		}
// 		break
// 	case 3:
// 		for i := 0; i < n; i++ {
// 			Sort(sv1, n)
// 			sv1[i].Xuat()
// 		}
// 		break
// 	case 4:
// 		for i := 0; i < n; i++ {
// 			sv1[i].Xuat()
// 			sv1[i].PhanLoai()
// 		}
// 		break
// 	case 5:
// 		f, err := os.Create("data.txt")

// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer f.Close()
// 		for i := 0; i < n; i++ {
// 			data := "Ten :" + sv1[i].ten + "  Gioi tinh :" + sv1[i].gioiting + "    Tuoi" + strconv.Itoa(sv1[i].tuoi) +
// 				"    Diem Hoa :" + strconv.Itoa(int(sv1[i].diemhoa)) + "    Diem li :" + strconv.Itoa(int(sv1[i].diemli)) + "    Diem Toan :" + strconv.Itoa(int(sv1[i].diemtoan))
// 			fmt.Println(data)
// 			_, err2 := f.WriteString(data)
// 			if err2 != nil {
// 				log.Fatal(err2)
// 			}
// 		}
// 		break
// 	case 0:
// 		return
// 	}
// }
