package requests

import "time"

type CreateBookingRequest struct {
	Id           int64     `json:"id"`
	Booking_code int64     `json:"code"`
	Booking_date time.Time `json:"date"`
	Customer_id  int64     `json:"customer_id"`
	Flight_id    int64     `json:"flight_id"`
	Status       string    `json:"status" binding:"required"`
}
type ViewRequest struct {
	Booking_code int64 `json:"code"`
}
