package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lntvan166/e2tech-driver-svc/internal/db"
	"github.com/lntvan166/e2tech-driver-svc/internal/pb"
	"github.com/lntvan166/e2tech-driver-svc/internal/utils"
)

func (s *Server) CreateAddress(context context.Context, req *pb.CreateAddressRequest) (*pb.CreateAddressResponse, error) {
	arg := db.CreateAddressParams{
		DriverID:    req.DriverId,
		Detail:      req.Detail,
		HouseNumber: req.HouseNumber,
		Street:      req.Street,
		Ward:        req.Ward,
		District:    req.District,
		City:        req.City,
	}
	addressArg := utils.Address{
		HouseNumber: req.HouseNumber,
		Street:      req.Street,
		Ward:        req.Ward,
		District:    req.District,
		City:        req.City,
	}
	location, err := utils.AddressToLocation(addressArg)
	if err != nil {
		return &pb.CreateAddressResponse{
			Status: http.StatusBadRequest,
			Error:  fmt.Sprintf("invalid address: %s", err),
		}, nil
	}

	arg.Latitude = location.Latitude
	arg.Longitude = location.Longitude

	address, err := s.DB.CreateAddress(context, arg)
	if err != nil {
		return &pb.CreateAddressResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("create address error: %s", err),
		}, nil
	}

	dataRsp := &pb.Address{
		Id:          address.ID,
		DriverId:    address.DriverID,
		Details:     address.Detail,
		HouseNumber: address.HouseNumber,
		Street:      address.Street,
		Ward:        address.Ward,
		District:    address.District,
		City:        address.City,
		Latitude:    address.Latitude,
		Longitude:   address.Longitude,
	}

	return &pb.CreateAddressResponse{
		Status:  http.StatusCreated,
		Address: dataRsp,
	}, nil
}

func (s *Server) GetAddress(context context.Context, req *pb.GetAddressRequest) (*pb.GetAddressResponse, error) {
	address, err := s.DB.GetAddressByDriverID(context, req.DriverId)
	if err != nil {
		return &pb.GetAddressResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("get address error: %s", err),
		}, nil
	}

	dataRsp := &pb.Address{
		Id:          address.ID,
		DriverId:    address.DriverID,
		Details:     address.Detail,
		HouseNumber: address.HouseNumber,
		Street:      address.Street,
		Ward:        address.Ward,
		District:    address.District,
		City:        address.City,
		Latitude:    address.Latitude,
		Longitude:   address.Longitude,
	}

	return &pb.GetAddressResponse{
		Status:  http.StatusOK,
		Address: dataRsp,
	}, nil
}

func (s *Server) GetLocation(context context.Context, req *pb.GetLocationRequest) (*pb.GetLocationResponse, error) {
	address, err := s.DB.GetAddressByDriverID(context, req.DriverId)
	if err != nil {
		return &pb.GetLocationResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("get address error: %s", err),
		}, nil
	}

	dataRsp := &pb.Location{
		Latitude:  address.Latitude,
		Longitude: address.Longitude,
	}

	return &pb.GetLocationResponse{
		Status:   http.StatusOK,
		Location: dataRsp,
	}, nil
}

func (s *Server) UpdateAddress(context context.Context, req *pb.UpdateAddressRequest) (*pb.UpdateAddressResponse, error) {
	arg := db.UpdateAddressParams{
		ID:          req.Id,
		Detail:      req.Detail,
		HouseNumber: req.HouseNumber,
		Street:      req.Street,
		Ward:        req.Ward,
		District:    req.District,
		City:        req.City,
	}
	addressArg := utils.Address{
		HouseNumber: req.HouseNumber,
		Street:      req.Street,
		Ward:        req.Ward,
		District:    req.District,
		City:        req.City,
	}
	location, err := utils.AddressToLocation(addressArg)
	if err != nil {
		return &pb.UpdateAddressResponse{
			Status: http.StatusBadRequest,
			Error:  fmt.Sprintf("invalid address: %s", err),
		}, nil
	}

	arg.Latitude = location.Latitude
	arg.Longitude = location.Longitude

	address, err := s.DB.UpdateAddress(context, arg)
	if err != nil {
		return &pb.UpdateAddressResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("update address error: %s", err),
		}, nil
	}

	dataRsp := &pb.Address{
		Id:          address.ID,
		DriverId:    address.DriverID,
		Details:     address.Detail,
		HouseNumber: address.HouseNumber,
		Street:      address.Street,
		Ward:        address.Ward,
		District:    address.District,
		City:        address.City,
		Latitude:    address.Latitude,
		Longitude:   address.Longitude,
	}

	return &pb.UpdateAddressResponse{
		Status:  http.StatusOK,
		Address: dataRsp,
	}, nil
}

func (s *Server) DeleteAddress(context context.Context, req *pb.DeleteAddressRequest) (*pb.DeleteAddressResponse, error) {
	err := s.DB.DeleteAddress(context, req.Id)
	if err != nil {
		return &pb.DeleteAddressResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("delete address error: %s", err),
		}, nil
	}

	return &pb.DeleteAddressResponse{
		Status: http.StatusOK,
	}, nil
}
