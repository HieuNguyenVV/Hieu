package requests

import "time"

type CreateBookingRequest struct {
	Id           int64     `json:"id"`
	BookingCode  int64     `json:"code"`
	Booking_date time.Time `json:"date"`
	CustomerId   int64     `json:"customerid"`
	FlightId     int64     `json:"flightid"`
	Status       string    `json:"status" binding:"required"`
}
type ViewRequest struct {
	Bookingcode int64 `json:"code"`
}
