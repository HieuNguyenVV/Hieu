package handlers

import (
	"context"
	"database/sql"
	"gin-training/grpc/grpc-flight/models"
	"gin-training/grpc/grpc-flight/repositories"
	"gin-training/pb"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FlightHandler struct {
	pb.UnimplementedFlightServiceServer
	flightRepository repositories.FlightRepository
}

func NewflightHandler(flightRepository repositories.FlightRepository) (*FlightHandler, error) {
	return &FlightHandler{
		flightRepository: flightRepository,
	}, nil
}
func (h *FlightHandler) CreateFlight(ctx context.Context, in *pb.Flight) (*pb.Flight, error) {
	pRequest := &models.Flight{}

	err := copier.Copy(&pRequest, &in)
	if err != nil {
		return nil, err
	}
	flight, err := h.flightRepository.CreateFlight(ctx, pRequest)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	pResponse := &pb.Flight{}

	err = copier.Copy(&pResponse, &flight)
	if err != nil {
		return nil, err
	}
	return pResponse, nil
}
func (h *FlightHandler) UpdateFlight(ctx context.Context, in *pb.Flight) (*pb.Flight, error) {
	flight, err := h.flightRepository.GetFlightByID(ctx, in.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	if in.Name != "" {
		flight.Name = in.Name
	}
	if in.From != "" {
		flight.From = in.From
	}
	if in.To != "" {
		flight.To = in.To
	}
	if in.Date.String() != "" {
		flight.Date = in.Date.AsTime()
	}
	if in.Status != "" {
		flight.Status = in.Status
	}
	if in.AvailableSlot != 0 {
		flight.AvailableSlot = in.AvailableSlot
	}
	newFlight, err := h.flightRepository.UpdateFlight(ctx, flight)
	if err != nil {
		return nil, err
	}

	pResponse := &pb.Flight{}

	err = copier.Copy(&pResponse, &newFlight)
	if err != nil {
		return nil, err
	}
	return pResponse, nil
}

func (h *FlightHandler) SearchFlight(ctx context.Context, in *pb.SearchFlightRequest) (*pb.ListFlight, error) {
	var (
		flights []*models.Flight
		err     error
	)

	if in.From == "" && in.To == "" {
		return nil, status.Error(codes.InvalidArgument, "Input is empty, please re-enter")
	}

	if in.From != "" {
		flights, err = h.flightRepository.GetFlightsByFrom(ctx, in.From)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, status.Error(codes.NotFound, err.Error())
			}
			return nil, err
		}
	} else if in.To != "" {
		flights, err = h.flightRepository.GetFlightsByDestination(ctx, in.To)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, status.Error(codes.NotFound, err.Error())
			}
			return nil, err
		}
	}

	pResponse := &pb.ListFlight{
		Flight: []*pb.Flight{},
	}

	err = copier.CopyWithOption(&pResponse.Flight, &flights, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
	if err != nil {
		return nil, err
	}

	return pResponse, nil
}
func (h *FlightHandler) FindFlight(ctx context.Context, in *pb.FindResquest) (*pb.Flight, error) {
	customer, err := h.flightRepository.GetFlightByID(ctx, in.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, err
	}

	cRes := pb.Flight{}
	err = copier.Copy(&cRes, &customer)
	if err != nil {
		return nil, err
	}

	return &cRes, nil
}
