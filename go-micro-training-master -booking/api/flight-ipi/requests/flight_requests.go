package requests

import "time"

type CreateFlightRequest struct {
	Id            int64     `json:"id"`
	Name          string    `json:"name" binding:"required"`
	From          string    `json:"from" binding:"required"`
	To            string    `json:"to" binding:"required"`
	Date          time.Time `json:"date" binding:"required"`
	Status        string    `json:"status" binding:"required"`
	AvailableSlot int64     `json:"available_slot" binding:"required"`
}

type SearchFlightRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}
