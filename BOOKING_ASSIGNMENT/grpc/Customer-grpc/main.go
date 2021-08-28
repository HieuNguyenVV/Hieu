package main

import (
	"fmt"
	"gin-training/grpc/Customer-grpc/handlers"
	"gin-training/grpc/Customer-grpc/repositories"
	"gin-training/grpc/helper"
	"gin-training/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := helper.AutoBindConfig("config.yml")
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", ":2223")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	customerRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := handlers.NewCustomerHandler(customerRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterCustomerServiceServer(s, h)

	fmt.Println("Listen at port: 2223")

	s.Serve(listen)
}
