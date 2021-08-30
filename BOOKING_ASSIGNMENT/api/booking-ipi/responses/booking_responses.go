package responses

import "time"

type CreateBookingResponses struct {
	Id           int64     `json:"id"`
	BookingCode  int64     `json:"code"`
	Booking_date time.Time `json:"date"`
	CustomerId   int64     `json:"customer_id"`
	FlightId     int64     `json:"flight_id"`
	Status       string    `json:"status" binding:"required"`
}
type ViewResponse struct {
	BookingCode int64 `json:"code"`
	Customer    Customer
	Flight      Flight
}

type Customer struct {
	Id      int64  `json:"id"`
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"max=256,min=6"`
	Email   string `json:"email"binding:"max=256,min=6"`
}
type Flight struct {
	Id            int64     `json:"id"`
	Name          string    `json:"name" binding:"required"`
	From          string    `json:"from" binding:"required"`
	To            string    `json:"to" binding:"required"`
	Date          time.Time `json:"date" binding:"required"`
	Status        string    `json:"status" binding:"required"`
	AvailableSlot int64     `json:"available_slot" binding:"required"`
}

type Cancel struct {
	Id           int64     `json:"id"`
	BookingCode  int64     `json:"code"`
	Booking_date time.Time `json:"booking_date"`
	Cancel_date  time.Time `json:"cancel_date"`
	CustomerId   int64     `json:"customer_id"`
	FlightId     int64     `json:"flight_id"`
	Status       string    `json:"status" binding:"required"`
}
