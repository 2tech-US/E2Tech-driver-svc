package services

import (
	"github.com/lntvan166/e2tech-driver-svc/internal/client"
	"github.com/lntvan166/e2tech-driver-svc/internal/config"
	"github.com/lntvan166/e2tech-driver-svc/internal/db"
	"github.com/lntvan166/e2tech-driver-svc/internal/pb"
)

type Server struct {
	DB     *db.Queries
	Config *config.Config
	pb.UnimplementedDriverServiceServer

	BookingSvc *client.BookingServiceClient
}
