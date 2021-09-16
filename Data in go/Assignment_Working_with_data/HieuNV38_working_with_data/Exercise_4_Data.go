package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sync"
// 	"time"
// )

// var check int = 0
// var wg sync.WaitGroup

// type Person struct {
// 	FullName    string
// 	Yearofbirth Date
// 	Address     string
// }
// type Date struct {
// 	date  int
// 	month int
// 	year  int
// }
// type Authentication struct {
// 	phone int
// 	email string
// }
// type User struct {
// 	Id             int
// 	ps             string
// 	p              Person
// 	authentication Authentication
// 	balance        int
// 	HS             []History
// }
// type DataTransfer struct {
// 	name1  int
// 	name2  int
// 	number int
// }
// type History struct {
// 	//code=1:sent,code=2:receive,code=3:withdrawal,code=4:deposit
// 	code              int
// 	idSend            int
// 	idreceive         int
// 	amount            int
// 	mamountwithdrawal int
// 	amountdeposit     int
// }

// func Login(user []User) {
// 	var id int
// 	var ps string
// loop:
// 	fmt.Println("Enter ID:")
// 	fmt.Scanln(&id)
// 	fmt.Println("Enter Ps:")
// 	fmt.Scanln(&ps)
// 	var k int = 0
// 	for i := 0; i < len(user); i++ {
// 		if id == user[i].Id && ps == user[i].ps {
// 			fmt.Println("Login OK")
// 		loop1:
// 			Menu1()
// 			var chose int
// 			fmt.Println("Enter chose:")
// 			fmt.Scanln(&chose)
// 			k = 1
// 			switch chose {
// 			case 1:
// 				Send(user, user[i])
// 				if check == 1 {
// 					check = 0
// 					goto loop1
// 				}
// 			case 2:
// 				deposit(user, user[i])
// 				if check == 1 {
// 					check = 0
// 					goto loop1
// 				}
// 			case 3:
// 				withdrawal(user, user[i])
// 				if check == 1 {
// 					check = 0
// 					goto loop1
// 				}
// 			case 4:
// 				os.Exit(0)
// 			}
// 		}
// 	}
// 	if k == 0 {
// 		fmt.Println("Invalid id or password, please re-enter:")
// 		goto loop
// 	}
// }
// func Register() User {
// 	reader := bufio.NewReader(os.Stdin)
// 	var (
// 		id       int
// 		ps       string
// 		FullName string
// 		date     int
// 		month    int
// 		year     int
// 		address  string
// 	)
// 	fmt.Println("Enter ID:")
// 	fmt.Scanln(&id)
// 	fmt.Println("Enter Ps:")
// 	fmt.Scanln(&ps)
// 	fmt.Print("Enter name: ")
// 	FullName, _ = reader.ReadString('\n')
// 	fmt.Println("Enter date:")
// 	fmt.Scanln(&date)
// 	fmt.Println("Enter month:")
// 	fmt.Scanln(&month)
// 	fmt.Println("Enter year:")
// 	fmt.Scanln(&year)
// 	fmt.Println("Enter addres:")
// 	address, _ = reader.ReadString('\n')

// 	check = 1
// 	return User{
// 		Id: id,
// 		ps: ps,
// 		p: Person{
// 			FullName: FullName,
// 			Yearofbirth: Date{
// 				date:  date,
// 				month: month,
// 				year:  year,
// 			},
// 			Address: address,
// 		},
// 		balance: 0,
// 	}

// }

// func Send(user []User, user1 User) {
// loop1:
// 	var number int
// 	var id int
// 	fmt.Println("Enter id :")
// 	fmt.Scanln(&id)
// 	var k int = 0
// 	fmt.Println("Enter transfer amount:")
// 	fmt.Scanln(&number)
// 	for i := 0; i < len(user); i++ {
// 		if user1.Id == user[i].Id {
// 			fmt.Println("co tru")
// 			user[i].balance = user[i].balance - number
// 		}
// 	}
// 	for i := 0; i < len(user); i++ {
// 		// if user1.Id == user[i].Id {
// 		// 	user[i].balance = user[i].balance - number
// 		// }
// 		if id == user[i].Id /*&& id != user1.Id*/ {
// 			user1.balance = user1.balance - number
// 			user[i].balance = user[i].balance + number
// 			fmt.Println("Your account has been deducted")
// 			fmt.Println("Balance:", user1.balance)
// 			break
// 		} else {
// 			if i == len(user)-1 {
// 				k = 1
// 			}
// 		}
// 	}
// 	if k == 1 {
// 		k = 0
// 		fmt.Println("Invalid id , please re-enter:")
// 		goto loop1
// 	}
// 	fmt.Println("Do you want to continue the transaction?Y/N")
// 	var str string
// 	fmt.Scanln(&str)
// 	if str == "Y" {
// 		goto loop1
// 	} else {
// 		check = 1
// 	}
// }
// func withdrawal(user []User, user1 User) {
// loop1:
// 	var number int
// 	fmt.Println("Enter withdrawal amount:")
// 	fmt.Scanln(&number)
// 	for i := 0; i < len(user); i++ {
// 		if user1.Id == user[i].Id {
// 			if number >= user[i].balance {
// 				fmt.Println("Withdrawal amount exceeds balance")
// 				break
// 			}
// 			if user[i].balance <= 0 {
// 				fmt.Println("Account not enough")
// 				break
// 			}
// 			//fmt.Println("co")
// 			user[i].balance = user[i].balance - number
// 			//fmt.Println(user[i].balance)
// 			fmt.Println("Your account:")
// 			fmt.Println("Balance:", user[i].balance)
// 		}

// 	}
// 	fmt.Println("Do you want to continue the transaction?Y/N")
// 	var str string
// 	fmt.Scanln(&str)
// 	if str == "Y" {
// 		goto loop1
// 	} else {
// 		check = 1
// 	}
// }
// func deposit(user []User, user1 User) {
// loop1:
// 	var number int
// 	fmt.Println("Enter deposits:")
// 	fmt.Scanln(&number)
// 	for i := 0; i < len(user); i++ {
// 		if user1.Id == user[i].Id {
// 			//fmt.Println("co")
// 			user[i].balance = user[i].balance + number
// 			//fmt.Println(user[i].balance)
// 			fmt.Println("Your account:")
// 			fmt.Println("Balance:", user[i].balance)
// 		}

// 	}
// 	fmt.Println("Do you want to continue the transaction?Y/N")
// 	var str string
// 	fmt.Scanln(&str)
// 	if str == "Y" {
// 		goto loop1
// 	} else {
// 		check = 1
// 	}
// }
// func deposituserGoroutine(user []User, user1 User) {
// 	var number int = 999999
// 	var data2 History
// 	data2.code = 4
// 	data2.amountdeposit = number
// 	for i := 0; i < len(user); i++ {
// 		if user1.Id == user[i].Id {
// 			user[i].balance = user[i].balance + number
// 			user[i].HS = append(user[i].HS, data2)
// 		}
// 	}
// 	fmt.Println("successful deposit!!!!!!!!")
// 	wg.Done()
// }
// func withdrawaluserGoroutine(user []User, user1 User) {
// 	var number int = 100000
// 	var data2 History
// 	data2.code = 3
// 	data2.mamountwithdrawal = number
// 	for i := 0; i < len(user); i++ {
// 		if user1.Id == user[i].Id {
// 			if number >= user[i].balance {
// 				fmt.Println("Withdrawal amount exceeds balance")
// 				break
// 			}
// 			if user[i].balance <= 0 {
// 				fmt.Println("Account not enough")
// 				break
// 			}
// 			user[i].balance = user[i].balance - number
// 			user[i].HS = append(user[i].HS, data2)
// 		}
// 	}
// 	fmt.Println("successful Withdrawal!!!!!!!!")
// 	wg.Done()
// }
// func worker(user []User) {
// 	input := make(chan DataTransfer)
// 	for i := 0; i < len(user); i++ {
// 		if i+1 == len(user) {
// 			break
// 		}
// 		wg.Add(4)
// 		go SenduserGoroutine(user, user[i], user[i+1], input)
// 		go ReceiveuserGoroutine(user, input)
// 		go withdrawaluserGoroutine(user, user[i])
// 		go deposituserGoroutine(user, user[i])
// 	}
// 	check = 1
// 	wg.Wait()
// }
// func SenduserGoroutine(user2 []User, user User, user1 User, input chan DataTransfer) {
// 	var data DataTransfer
// 	var a int = 1000000
// 	for i := 0; i < len(user2); i++ {
// 		if user.Id == user2[i].Id {
// 			user2[i].balance = user2[i].balance - a
// 			fmt.Println("Successful transfer to ", user1.p.FullName)
// 		}
// 	}
// 	data.name1 = user.Id
// 	data.name2 = user1.Id
// 	data.number = a
// 	var data2 History
// 	data2.code = 1
// 	data2.idSend = user.Id
// 	data2.idreceive = user1.Id
// 	data2.amount = a
// 	for i := 0; i < len(user2); i++ {
// 		if user.Id == user2[i].Id {
// 			user2[i].HS = append(user2[i].HS, data2)
// 		}
// 	}
// 	user.HS = append(user.HS, data2)
// 	wg.Done()
// 	input <- data
// }
// func ReceiveuserGoroutine(user []User, input chan DataTransfer) {
// 	data := <-input
// 	var data2 History
// 	data2.code = 2
// 	data2.idSend = data.name1
// 	data2.idreceive = data.name2
// 	data2.amount = data.number
// 	for i := 0; i < len(user); i++ {
// 		if data.name2 == user[i].Id {
// 			user[i].balance = user[i].balance + data.number
// 			user[i].HS = append(user[i].HS, data2)
// 		}
// 	}
// 	for i := 0; i < len(user); i++ {
// 		if data.name1 == user[i].Id {
// 			fmt.Println("Account received ", data.number, "From", user[i].p.FullName)
// 		}
// 	}
// 	wg.Done()
// }
// func Output(user []User) {
// 	for i := 0; i < len(user); i++ {
// 		fmt.Println("Id:", user[i].Id)
// 		fmt.Println("FullName:", user[i].p.FullName)
// 		fmt.Println("Date of birth :", user[i].p.Yearofbirth.date, "/", user[i].p.Yearofbirth.month, "/", user[i].p.Yearofbirth.year)
// 		fmt.Println("Address:", user[i].p.Address)
// 		fmt.Println("Balance :", user[i].balance)
// 		fmt.Println("Email:", user[i].authentication.email)
// 		fmt.Println("Phone:", user[i].authentication.phone)
// 		fmt.Println("Transaction history:")
// 		for j := 0; j < len(user[i].HS); j++ {
// 			if user[i].HS[j].code == 1 {
// 				fmt.Println("Send to ID:", user[i].HS[j].idreceive)
// 				fmt.Println("Transfer amount:", user[i].HS[j].amount)
// 				fmt.Println()
// 			}
// 			if user[i].HS[j].code == 2 {
// 				fmt.Println("Receive from ID:", user[i].HS[j].idSend)
// 				fmt.Println("Amount received:", user[i].HS[j].amount)
// 				fmt.Println()
// 			}
// 			if user[i].HS[j].code == 3 {
// 				fmt.Println("Amount withdrawn:", user[i].HS[j].mamountwithdrawal)
// 				fmt.Println()
// 			}
// 			if user[i].HS[j].code == 4 {
// 				fmt.Println("Amount deposit :", user[i].HS[j].amountdeposit)
// 				fmt.Println()
// 			}
// 		}
// 		fmt.Println()
// 	}
// }
// func Contructer() []User {
// 	user := []User{
// 		User{
// 			Id: 1,
// 			ps: "1234",
// 			p: Person{
// 				FullName: "Nguyen Van A",
// 				Yearofbirth: Date{
// 					date:  16,
// 					month: 02,
// 					year:  1998,
// 				},
// 				Address: "Ha Noi",
// 			},
// 			authentication: Authentication{
// 				phone: 942104802,
// 				email: "NguyenVanA@gmail.com",
// 			},
// 			balance: 100000000,
// 			HS:      []History{},
// 		},
// 		User{
// 			Id: 2,
// 			ps: "1234",
// 			p: Person{
// 				FullName: "Nguyen Van B",
// 				Yearofbirth: Date{
// 					date:  16,
// 					month: 02,
// 					year:  1998,
// 				},
// 				Address: "Ha Noi",
// 			},
// 			authentication: Authentication{
// 				phone: 942104802,
// 				email: "NguyenVanA@gmail.com",
// 			},
// 			balance: 100000000,
// 			HS:      []History{},
// 		},
// 		User{
// 			Id: 3,
// 			ps: "1234",
// 			p: Person{
// 				FullName: "Nguyen Van C",
// 				Yearofbirth: Date{
// 					date:  16,
// 					month: 02,
// 					year:  1998,
// 				},
// 				Address: "Ha Noi",
// 			},
// 			authentication: Authentication{
// 				phone: 942104802,
// 				email: "NguyenVanA@gmail.com",
// 			},
// 			balance: 100000000,
// 			HS:      []History{},
// 		},
// 		User{
// 			Id: 4,
// 			ps: "1234",
// 			p: Person{
// 				FullName: "Nguyen Van D",
// 				Yearofbirth: Date{
// 					date:  16,
// 					month: 02,
// 					year:  1998,
// 				},
// 				Address: "Ha Noi",
// 			},
// 			authentication: Authentication{
// 				phone: 942104802,
// 				email: "NguyenVanA@gmail.com",
// 			},
// 			balance: 100000000,
// 			HS:      []History{},
// 		},
// 		User{
// 			Id: 5,
// 			ps: "1234",
// 			p: Person{
// 				FullName: "Nguyen Van E",
// 				Yearofbirth: Date{
// 					date:  16,
// 					month: 02,
// 					year:  1998,
// 				},
// 				Address: "Ha Noi",
// 			},
// 			authentication: Authentication{
// 				phone: 942104802,
// 				email: "NguyenVanA@gmail.com",
// 			},
// 			balance: 100000000,
// 			HS:      []History{},
// 		},
// 		User{
// 			Id: 6,
// 			ps: "1234",
// 			p: Person{
// 				FullName: "Nguyen Van F",
// 				Yearofbirth: Date{
// 					date:  16,
// 					month: 02,
// 					year:  1998,
// 				},
// 				Address: "Ha Noi",
// 			},
// 			authentication: Authentication{
// 				phone: 942104802,
// 				email: "NguyenVanA@gmail.com",
// 			},
// 			balance: 100000000,
// 			HS:      []History{},
// 		},
// 		User{
// 			Id: 7,
// 			ps: "1234",
// 			p: Person{
// 				FullName: "Nguyen Van G",
// 				Yearofbirth: Date{
// 					date:  16,
// 					month: 02,
// 					year:  1998,
// 				},
// 				Address: "Ha Noi",
// 			},
// 			authentication: Authentication{
// 				phone: 942104802,
// 				email: "NguyenVanA@gmail.com",
// 			},
// 			balance: 100000000,
// 			HS:      []History{},
// 		},
// 		User{
// 			Id: 8,
// 			ps: "1234",
// 			p: Person{
// 				FullName: "Nguyen Van H",
// 				Yearofbirth: Date{
// 					date:  16,
// 					month: 02,
// 					year:  1998,
// 				},
// 				Address: "Ha Noi",
// 			},
// 			authentication: Authentication{
// 				phone: 942104802,
// 				email: "NguyenVanA@gmail.com",
// 			},
// 			balance: 100000000,
// 			HS:      []History{},
// 		}, User{
// 			Id: 9,
// 			ps: "1234",
// 			p: Person{
// 				FullName: "Nguyen Van K",
// 				Yearofbirth: Date{
// 					date:  16,
// 					month: 02,
// 					year:  1998,
// 				},
// 				Address: "Ha Noi",
// 			},
// 			authentication: Authentication{
// 				phone: 942104802,
// 				email: "NguyenVanA@gmail.com",
// 			},
// 			balance: 100000000,
// 			HS:      []History{},
// 		}, User{
// 			Id: 10,
// 			ps: "1234",
// 			p: Person{
// 				FullName: "Nguyen Van L",
// 				Yearofbirth: Date{
// 					date:  16,
// 					month: 02,
// 					year:  1998,
// 				},
// 				Address: "Ha Noi",
// 			},
// 			authentication: Authentication{
// 				phone: 942104802,
// 				email: "NguyenVanA@gmail.com",
// 			},
// 			balance: 100000000,
// 			HS:      []History{},
// 		},
// 	}
// 	return user
// }
// func Menu() {
// 	fmt.Println("1:Login.")
// 	fmt.Println("2:Register.")
// 	fmt.Println("3:Multiple concurrent.")
// 	fmt.Println("4:Output.")
// 	fmt.Println("5:Quit.")
// }
// func Menu1() {
// 	fmt.Println("1:Transfers money.")
// 	fmt.Println("2:Deposit account .")
// 	fmt.Println("3:Withdrawal.")
// 	fmt.Println("4:Quit.")
// }

// func main() {
// 	user := Contructer()
// 	// worker(user)
// 	var chose int
// loop:
// 	Menu()
// 	fmt.Println("Enter chose:")
// 	fmt.Scanln(&chose)
// 	switch chose {
// 	case 1:
// 		Login(user)
// 	case 2:
// 		user = append(user, Register())
// 		if check == 1 {
// 			fmt.Println("successful registration")
// 			time.Sleep(time.Second * 2)
// 			goto loop
// 		}
// 	case 3:
// 		worker(user)
// 		//Output(user)
// 		if check == 1 {
// 			time.Sleep(time.Second * 2)
// 			goto loop
// 		}
// 	case 4:
// 		Output(user)
// 	case 5:
// 		os.Exit(0)
// 	}

// }
