package repositories

import (
	"context"
	"gin-training/database"
	"gin-training/grpc/booking-grpc/models"
	"time"

	"gorm.io/gorm"
)

type BookingRepositories interface {
	Booking(ctx context.Context, booking *models.Booking) (*models.Booking, error)
	GetBookingByCode(context context.Context, BookingCode int64) (*models.Booking, error)
	DeleteBooking(ctx context.Context, BookingCode int64) (*models.CancelBooking, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (BookingRepositories, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Booking{})

	if err != nil {
		return nil, err
	}

	return &dbmanager{db.Debug()}, nil
}

func (m *dbmanager) Booking(ctx context.Context, booking *models.Booking) (*models.Booking, error) {
	if err := m.Save(booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}
func (m *dbmanager) GetBookingByCode(context context.Context, BookingCode int64) (*models.Booking, error) {
	booking := models.Booking{}
	if err := m.Where(&models.Booking{Id: BookingCode}).First(&booking).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}
func (m *dbmanager) DeleteBooking(ctx context.Context, BookingCode int64) (*models.CancelBooking, error) {
	booking, err := m.GetBookingByCode(ctx, BookingCode)
	if err != nil {
		return nil, err
	}

	if err := m.Unscoped().Delete(&booking).Error; err != nil {
		return nil, err
	}
	return &models.CancelBooking{
		Id:          booking.Id,
		BookingCode: booking.BookingCode,
		BookingDate: booking.BookingDate,
		Cancel_date: time.Time{},
		CustomerId:  booking.CustomerId,
		FlightId:    booking.FlightId,
		Status:      booking.Status,
	}, nil
}
