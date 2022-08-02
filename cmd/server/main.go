package main

import (
	"fmt"
	"log"
	"net"

	"github.com/lntvan166/e2tech-driver-svc/internal/config"
	"github.com/lntvan166/e2tech-driver-svc/internal/db"
	"github.com/lntvan166/e2tech-driver-svc/internal/pb"
	"github.com/lntvan166/e2tech-driver-svc/internal/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	DB := db.Connect(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Driver Svc on", c.Port)

	s := services.Server{
		DB:     DB,
		Config: &c,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterDriverServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
