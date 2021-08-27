package main

import (
	"gin-training/api/flight-ipi/handlers"
	"gin-training/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	flightClient := pb.NewFlightServiceClient(conn)
	h := handlers.NewFlightHandler(flightClient)
	g := gin.Default()
	gr := g.Group("/v1/ipi")
	gr.POST("/create", h.CreateFlight)
	gr.PUT("/update", h.UpdateFlight)
	gr.PUT("/search", h.SearchFly)
	http.ListenAndServe(":8080", g)
}
