package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"gin-training/grpc/booking-grpc/models"
	"gin-training/grpc/booking-grpc/repositories"
	"gin-training/pb"
	"time"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type BookingHandler struct {
	customerClient    pb.CustomerServiceClient
	flightClient      pb.FlightServiceClient
	bookingRepository repositories.BookingRepositories
	pb.UnimplementedBookingsServer
}

func NewBookingHandler(customerClient pb.CustomerServiceClient,
	flightClient pb.FlightServiceClient,
	bookingRepository repositories.BookingRepositories) (*BookingHandler, error) {
	return &BookingHandler{
		customerClient:    customerClient,
		flightClient:      flightClient,
		bookingRepository: bookingRepository,
	}, nil
}

func (h *BookingHandler) Booking(ctx context.Context, in *pb.Info) (*pb.Info, error) {
	fmt.Println("dgfjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj")
	if in.CustomerId == 0 {
		return nil, status.Error(codes.InvalidArgument, "Customer_id is required")
	}

	if in.FlightId == 0 {
		return nil, status.Error(codes.InvalidArgument, "Fligh_id is required")
	}
	fmt.Println("dgfjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj")
	customer, err := h.customerClient.FindCustomer(ctx, &pb.FindRequest{
		Id: in.CustomerId,
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "Customer_id is not found")
			}
		} else {
			return nil, err
		}
	}
	fmt.Println("dgfjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj")
	flight, err := h.flightClient.FindFlight(ctx, &pb.FindResquest{
		Id: in.FlightId,
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "Flight_id is not found")
			}
		} else {
			return nil, err
		}
	}
	fmt.Println("dgfjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj")
	booking := &models.Booking{
		Id:          in.Id,
		BookingCode: in.BookingCode,
		BookingDate: in.BookingDate.AsTime(),
		CustomerId:  customer.Id,
		FlightId:    flight.Id,
		Status:      in.Status,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   gorm.DeletedAt{},
	}
	fmt.Println("dgfjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj")
	res, err := h.bookingRepository.Booking(ctx, booking)
	if err != nil {
		return nil, err
	}
	fmt.Println("dgfjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj")
	pRes := &pb.Info{}
	err = copier.Copy(&pRes, &res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}

func (h *BookingHandler) ViewBooking(ctx context.Context, in *pb.ViewRequest) (*pb.ViewResponse, error) {
	booking, err := h.bookingRepository.GetBookingByCode(ctx, in.BookingCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	customer, err := h.customerClient.FindCustomer(ctx, &pb.FindRequest{
		Id: booking.CustomerId,
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "Customer_id is not found")
			}
		} else {
			return nil, err
		}
	}
	fmt.Println("dgfjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj")
	flight, err := h.flightClient.FindFlight(ctx, &pb.FindResquest{
		Id: booking.FlightId,
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "Flight_id is not found")
			}
		} else {
			return nil, err
		}
	}

	viewbooking := &pb.ViewResponse{
		BookingCode: booking.BookingCode,
		BookingDate: timestamppb.New(booking.BookingDate),
		Customer:    customer,
		Flight:      flight,
	}

	return viewbooking, nil
}
func (h *BookingHandler) CancelBooking(ctx context.Context, in *pb.ViewRequest) (*pb.Cancel, error) {
	booking, err := h.bookingRepository.DeleteBooking(ctx, in.BookingCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	cancelbookimg := &pb.Cancel{
		Id:          booking.Id,
		BookingCode: booking.BookingCode,
		BookingDate: timestamppb.New(booking.BookingDate),
		CancelDate:  &timestamppb.Timestamp{},
		CustomerId:  booking.CustomerId,
		FlightId:    booking.FlightId,
		Status:      booking.Status,
	}

	return cancelbookimg, nil
}
