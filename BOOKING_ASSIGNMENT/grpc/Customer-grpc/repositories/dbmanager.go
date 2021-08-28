package repositories

import (
	"context"
	"gin-training/database"
	"gin-training/grpc/Customer-grpc/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error)
	GetCustomerByID(context context.Context, id int64) (*models.Customer, error)
	ChangePassword(ctx context.Context, model *models.Customer) (*models.MessageChangePassword, error)
}
type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (CustomerRepository, error) {
	//db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=1234 dbname=Customer port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Customer{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db.Debug()}, nil
}

func (m *dbmanager) GetCustomerByID(context context.Context, id int64) (*models.Customer, error) {
	customer := models.Customer{}
	if err := m.Where(&models.Customer{Id: id}).First(&customer).Error; err != nil {
		return nil, err
	}
	//fmt.Println(people)
	return &customer, nil
}
func (m *dbmanager) CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}
func (m *dbmanager) UpdateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {

	if err := m.Save(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}
func (m *dbmanager) ChangePassword(ctx context.Context, model *models.Customer) (*models.MessageChangePassword, error) {
	if err := m.Save(model).Error; err != nil {
		return nil, err
	}

	return &models.MessageChangePassword{
		Message: "ChangePassword Success!!",
	}, nil
}
