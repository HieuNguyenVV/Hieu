package responest

import "time"

type User struct {
	Id       int64  `json:"id" `
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Imail    string `json:"imail"`
	Created  time.Time
	Update   time.Time
}
