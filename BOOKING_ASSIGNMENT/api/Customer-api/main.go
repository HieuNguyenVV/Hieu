package main

import (
	"gin-training/api/Customer-api/handlers"
	"gin-training/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	//Create grpc client connect
	customerConn, err := grpc.Dial(":2223", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	customerClient := pb.NewCustomerServiceClient(customerConn)

	//Handler for GIN Gonic
	h := handlers.NewPeopleHandler(customerClient)

	g := gin.Default()
	//Create routes
	gr := g.Group("/v1/api")
	gr.POST("/create", h.CreateCustomer)
	gr.PUT("/update", h.UpdateCustomer)
	gr.PUT("/changePassword", h.ChangePassword)
	//gr.POST("/Find", h.FindPeople)
	//Listen and serve
	http.ListenAndServe(":8080", g)
}
