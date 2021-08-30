package handlers

import (
	"context"
	"database/sql"
	"gin-training/grpc/Customer-grpc/models"
	"gin-training/grpc/Customer-grpc/repositories"
	"gin-training/pb"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomerHandler struct {
	pb.UnimplementedCustomerServiceServer
	customerRepository repositories.CustomerRepository
}

func NewCustomerHandler(customerRepository repositories.CustomerRepository) (*CustomerHandler, error) {
	return &CustomerHandler{
		customerRepository: customerRepository,
	}, nil
}
func (h *CustomerHandler) CreateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	// pRequest := &models.Customer{
	// 	Id:           in.Id,
	// 	CustomerName: in.CustomerName,
	// 	Address: sql.NullString{
	// 		String: in.Address,
	// 		Valid:  true,
	// 	},
	// 	License:   in.License,
	// 	Phone:     in.Phone,
	// 	Email:     in.Email,
	// 	Password:  in.Password,
	// 	Active:    false,
	// 	CreatedAt: time.Time{},
	// 	UpdatedAt: time.Time{},
	// 	DeletedAt: gorm.DeletedAt{},
	// }
	pRequest := &models.Customer{}
	err := copier.Copy(&pRequest, &in)
	if err != nil {
		return nil, err
	}
	customer, err := h.customerRepository.CreateCustomer(ctx, pRequest)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	pResponse := &pb.Customer{}
	err = copier.Copy(&pResponse, &customer)
	if err != nil {
		return nil, err
	}
	return pResponse, nil
}

func (h *CustomerHandler) UpdateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	customer, err := h.customerRepository.GetCustomerByID(ctx, in.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	if in.Address != "" {
		customer.Address = sql.NullString{
			Valid:  true,
			String: in.Address,
		}
	}

	if in.CustomerName != "" {
		customer.CustomerName = in.CustomerName
	}

	if in.Phone != "" {
		customer.Phone = in.Phone
	}
	if in.License != "" {
		customer.License = in.License
	}
	if in.Email != "" {
		customer.Email = in.Email
	}
	if in.Password != "" {
		customer.Password = in.Password
	}
	if in.Active != customer.Active {
		customer.Active = in.Active
	}
	newCustomer, err := h.customerRepository.UpdateCustomer(ctx, customer)
	if err != nil {
		return nil, err
	}
	pResponse := &pb.Customer{}
	err = copier.Copy(&pResponse, &newCustomer)

	if err != nil {
		return nil, err
	}
	return pResponse, nil
}
func (h *CustomerHandler) ChangePassword(ctx context.Context, in *pb.ChangePasswordRequest) (*pb.MessageChangePassword, error) {
	customer, err := h.customerRepository.GetCustomerByID(ctx, in.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	if in.Oldpassword != customer.Password {
		return nil, status.Error(codes.PermissionDenied, "Your old password is incorrect")
	}

	if in.Newpassword != in.Confirmpassword {
		return nil, status.Error(codes.Canceled, "Makes sure your password match")
	} else {
		customer.Password = in.Newpassword
	}

	Message, err := h.customerRepository.ChangePassword(ctx, customer)
	if err != nil {
		return nil, err
	}

	pResponse := &pb.MessageChangePassword{}
	err = copier.Copy(&pResponse, &Message)
	if err != nil {
		return nil, err
	}
	return pResponse, nil
}
func (h *CustomerHandler) FindCustomer(ctx context.Context, in *pb.FindRequest) (*pb.FindResponse, error) {
	customer, err := h.customerRepository.GetCustomerByID(ctx, in.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, err
	}

	cRes := pb.FindResponse{
		Id:      customer.Id,
		Name:    customer.CustomerName,
		Address: customer.Address.String,
		Email:   customer.Email,
	}

	return &cRes, nil
}
