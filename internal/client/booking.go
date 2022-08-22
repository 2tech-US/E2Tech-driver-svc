package client

import (
	"context"
	"fmt"

	"github.com/lntvan166/e2tech-driver-svc/internal/config"
	"github.com/lntvan166/e2tech-driver-svc/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BookingServiceClient struct {
	Client pb.BookingServiceClient
}

func InitBookingServiceClient(c *config.Config) pb.BookingServiceClient {
	cc, err := grpc.Dial(c.BookingSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewBookingServiceClient(cc)
}

type UpdateResponseRequest struct {
	DriverPhone string
	Latitude    float64
	Longitude   float64
}

func (s *BookingServiceClient) UpdateResponse(ctx context.Context, req *UpdateResponseRequest) (*pb.UpdateResponseResponse, error) {
	return s.Client.UpdateResponse(ctx, &pb.UpdateResponseRequest{
		DriverPhone: req.DriverPhone,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
	})
}
