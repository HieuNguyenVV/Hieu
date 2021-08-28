package main

import (
	"fmt"
	"gin-training/grpc/grpc-flight/handlers"
	"gin-training/grpc/grpc-flight/repositories"
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
	listen, err := net.Listen("tcp", ":2222")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	flightRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := handlers.NewflightHandler(flightRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterFlightServiceServer(s, h)

	fmt.Println("Listen at port: 2222")

	s.Serve(listen)
}
