package main

import (
	"gin-training/api/booking-ipi/handlers"
	"gin-training/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	bookingConn, err := grpc.Dial(":2224", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	bookingClient := pb.NewBookingsClient(bookingConn)

	//Handler for GIN Gonic
	h := handlers.NewBookingHandler(bookingClient)

	g := gin.Default()
	//Create routes
	gr := g.Group("/v3/api")
	gr.POST("/booking", h.Booking)
	gr.POST("/viewbooking", h.ViewBooking)
	gr.PUT("/cancel", h.CancelBooking)
	//gr.POST("/Find", h.FindPeople)
	//Listen and serve
	http.ListenAndServe(":8080", g)

}
