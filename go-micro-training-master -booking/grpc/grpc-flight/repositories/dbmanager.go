package repositories

import (
	"context"
	"gin-training/database"
	"gin-training/grpc/grpc-flight/models"

	"gorm.io/gorm"
)

type FlightRepository interface {
	CreateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error)
	UpdateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error)
	GetFlightByID(ctx context.Context, id int64) (*models.Flight, error)
	GetFlightsByFrom(ctx context.Context, from string) ([]*models.Flight, error)
	GetFlightsByDestination(ctx context.Context, destination string) ([]*models.Flight, error)
}
type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (FlightRepository, error) {
	//db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=1234 dbname=Flight port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Flight{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db.Debug()}, nil
}
func (m *dbmanager) GetFlightByID(context context.Context, id int64) (*models.Flight, error) {
	flight := models.Flight{}
	if err := m.Where(&models.Flight{Id: id}).First(&flight).Error; err != nil {
		return nil, err
	}
	//fmt.Println(people)
	return &flight, nil
}
func (m *dbmanager) GetFlightsByFrom(ctx context.Context, from string) ([]*models.Flight, error) {
	flights := []*models.Flight{}
	if err := m.Where(&models.Flight{From: from}).Find(&flights).Error; err != nil {
		return nil, err
	}
	return flights, nil
}

func (m *dbmanager) GetFlightsByDestination(ctx context.Context, destination string) ([]*models.Flight, error) {
	flights := []*models.Flight{}
	if err := m.Where(&models.Flight{To: destination}).Find(&flights).Error; err != nil {
		return nil, err
	}
	return flights, nil
}
func (m *dbmanager) CreateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}
func (m *dbmanager) UpdateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error) {
	// if err := m.Where(&models.Flight{Id: model.Id}).Updates(&models.Flight{}).Error; err != nil {
	// 	return nil, err
	// }
	if err := m.Save(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

// UpdateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error)
// 	GetFlightByID(context context.Context, id uuid.UUID) (*models.Flight, error)
// }
