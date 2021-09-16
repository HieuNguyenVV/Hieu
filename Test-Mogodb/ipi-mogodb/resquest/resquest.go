package resquest

import "time"

type User struct {
	Id       int64  `json:"id" uri:"id"`
	Name     string `json:"name"  binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone"  binding:"required"`
	Imail    string `json:"imail"  binding:"required"`
	Created  time.Time
	Update   time.Time
}
type Userudate struct {
	Id       int64  `json:"id" uri:"id"`
	Name     string `json:"name"  `
	Password string `json:"password" `
	Phone    string `json:"phone"  `
	Imail    string `json:"imail"  `
	Created  time.Time
	Update   time.Time
}
type UeserResquestID struct {
	Id int64 `json:"id" uri:"id"`
}
type UserResquestName struct {
	Name string `json:"name"  uri:"name"`
}
