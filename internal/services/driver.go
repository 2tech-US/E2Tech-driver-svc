package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/lntvan166/e2tech-driver-svc/internal/db"
	"github.com/lntvan166/e2tech-driver-svc/internal/pb"
	"github.com/lntvan166/e2tech-driver-svc/internal/utils"
)

func (s *Server) CreateDriver(context context.Context, req *pb.CreateDriverRequest) (*pb.CreateDriverResponse, error) {
	_, err := s.DB.GetDriverByPhone(context, req.Phone)
	if err != sql.ErrNoRows {
		return &pb.CreateDriverResponse{
			Status: http.StatusBadRequest,
			Error:  "driver already exists",
		}, nil
	}

	arg := db.CreateDriverParams{
		Phone: req.Phone,
		Name:  req.Name,
	}

	driver, err := s.DB.CreateDriver(context, arg)
	if err != nil {
		return &pb.CreateDriverResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("svc: create driver error: %s", err),
		}, nil
	}

	dataRsp := &pb.Driver{
		Id:          driver.ID,
		Phone:       driver.Phone,
		Name:        driver.Name,
		DateOfBirth: utils.ParsedDateToString(driver.DateOfBirth.Time),
	}

	return &pb.CreateDriverResponse{
		Status: http.StatusCreated,
		Driver: dataRsp,
	}, nil
}

func (s *Server) GetDriverByPhone(context context.Context, req *pb.GetDriverByPhoneRequest) (*pb.GetDriverByPhoneResponse, error) {
	driver, err := s.DB.GetDriverByPhone(context, req.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.GetDriverByPhoneResponse{
				Status: http.StatusBadRequest,
				Error:  "user not found",
			}, nil
		}

		return &pb.GetDriverByPhoneResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to get user",
		}, nil
	}

	dataRsp := &pb.Driver{
		Id:          driver.ID,
		Phone:       driver.Phone,
		Name:        driver.Name,
		DateOfBirth: utils.ParsedDateToString(driver.DateOfBirth.Time),
	}

	return &pb.GetDriverByPhoneResponse{
		Status: http.StatusOK,
		Driver: dataRsp,
	}, nil
}

func (s *Server) GetLocation(context context.Context, req *pb.GetLocationRequest) (*pb.GetLocationResponse, error) {
	driver, err := s.DB.GetDriverByPhone(context, req.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.GetLocationResponse{
				Status: http.StatusBadRequest,
				Error:  "user not found",
			}, nil
		}

		return &pb.GetLocationResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to get user",
		}, nil
	}

	return &pb.GetLocationResponse{
		Status:    http.StatusOK,
		Longitude: utils.Float64FromNull(driver.Longitude),
		Latitude:  utils.Float64FromNull(driver.Latitude),
	}, nil
}

func (s *Server) ListDrivers(context context.Context, req *pb.ListDriversRequest) (*pb.ListDriversResponse, error) {
	arg := db.ListDriversParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	drivers, err := s.DB.ListDrivers(context, arg)
	if err != nil {
		return &pb.ListDriversResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to list drivers",
		}, nil
	}

	dataRsp := make([]*pb.Driver, len(drivers))
	for i, driver := range drivers {
		dataRsp[i] = &pb.Driver{
			Id:          driver.ID,
			Phone:       driver.Phone,
			Name:        driver.Name,
			DateOfBirth: utils.ParsedDateToString(driver.DateOfBirth.Time),
		}
	}

	return &pb.ListDriversResponse{
		Status: http.StatusOK,
		Driver: dataRsp,
	}, nil
}

func (s *Server) UpdateDriver(context context.Context, req *pb.UpdateDriverRequest) (*pb.UpdateDriverResponse, error) {
	driver, err := s.DB.GetDriverByPhone(context, req.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.UpdateDriverResponse{
				Status: http.StatusBadRequest,
				Error:  "user not found",
			}, nil
		}

		return &pb.UpdateDriverResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to get user",
		}, nil
	}

	strDate, err := utils.ParseStringToDate(req.DateOfBirth)
	sqlDate := sql.NullTime{Time: strDate, Valid: true}
	if err != nil {
		return &pb.UpdateDriverResponse{
			Status: http.StatusBadRequest,
			Error:  "invalid date of birth",
		}, nil
	}

	arg := db.UpdateDriverParams{
		ID:          req.Id,
		Phone:       req.Phone,
		Name:        req.Name,
		DateOfBirth: sqlDate,
	}

	driver, err = s.DB.UpdateDriver(context, arg)
	if err != nil {
		return &pb.UpdateDriverResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to update driver",
		}, nil
	}

	dataRsp := &pb.Driver{
		Id:          driver.ID,
		Phone:       driver.Phone,
		Name:        driver.Name,
		DateOfBirth: utils.ParsedDateToString(driver.DateOfBirth.Time),
	}

	return &pb.UpdateDriverResponse{
		Status: http.StatusOK,
		Driver: dataRsp,
	}, nil
}

func (s *Server) DeleteDriver(context context.Context, req *pb.DeleteDriverRequest) (*pb.DeleteDriverResponse, error) {
	driver, err := s.DB.GetDriverByPhone(context, req.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.DeleteDriverResponse{
				Status: http.StatusBadRequest,
				Error:  "user not found",
			}, nil
		}

		return &pb.DeleteDriverResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to get user",
		}, nil
	}

	err = s.DB.DeleteDriver(context, driver.Phone)
	if err != nil {
		return &pb.DeleteDriverResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to delete driver",
		}, nil
	}

	return &pb.DeleteDriverResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) UpdateLocation(context context.Context, req *pb.UpdateLocationRequest) (*pb.UpdateLocationResponse, error) {
	_, err := s.DB.GetDriverByPhone(context, req.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.UpdateLocationResponse{
				Status: http.StatusBadRequest,
				Error:  "user not found",
			}, nil
		}

		return &pb.UpdateLocationResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to get user",
		}, nil
	}

	arg := db.UpdateLocationParams{
		Phone:     req.Phone,
		Latitude:  utils.NullFloat64(req.Latitude),
		Longitude: utils.NullFloat64(req.Longitude),
	}

	_, err = s.DB.UpdateLocation(context, arg)
	if err != nil {
		return &pb.UpdateLocationResponse{
			Status: http.StatusInternalServerError,
			Error:  "failed to update location",
		}, nil
	}

	return &pb.UpdateLocationResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) UpdateStatus(context context.Context, req *pb.UpdateStatusRequest) (*pb.UpdateStatusResponse, error) {
	_, err := s.DB.GetDriverByPhone(context, req.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.UpdateStatusResponse{
				Status: http.StatusBadRequest,
				Error:  "driver: user not found",
			}, nil
		}

		return &pb.UpdateStatusResponse{
			Status: http.StatusInternalServerError,
			Error:  "driver: failed to get user",
		}, nil
	}

	arg := db.UpdateStatusParams{
		Phone:  req.Phone,
		Status: req.Status,
	}

	_, err = s.DB.UpdateStatus(context, arg)
	if err != nil {
		return &pb.UpdateStatusResponse{
			Status: http.StatusInternalServerError,
			Error:  "driver: failed to update status",
		}, nil
	}

	return &pb.UpdateStatusResponse{
		Status: http.StatusOK,
	}, nil
}
