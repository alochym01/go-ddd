package service

import (
	"alochym/domain"
)

type CustomerServiceRepository interface {
	GetAllCustomers() ([]domain.Customer, error)
	// GetCustomer(id string) (*domain.Customer, error)
	// CreateCustomer(uRequest dto.CustomerUserRequestObject) error
	// UpdateCustomer(id string, uReqUpdate dto.CustomerUserRequestUpdateObject) error
	// DeleteCustomer(id string) error
}

type CustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerServer(repository domain.CustomerRepository) CustomerService {
	return CustomerService{repo: repository}
}

func (c CustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return c.repo.FindAll()
}

// func (c CustomerService) GetCustomer(id string) (*domain.Customer, error) {
// 	return c.repo.FindByID(id)
// }

// func (c CustomerService) CreateCustomer(uRequest dto.CustomerUserRequestObject) error {
// 	temp := domain.Customer{
// 		Name:   uRequest.Name,
// 		City:   uRequest.City,
// 		Status: uRequest.Status,
// 	}
// 	return c.repo.Create(temp)
// }
// func (c CustomerService) UpdateCustomer(id string, uReqUpdate dto.CustomerUserRequestUpdateObject) error {
// 	temp := domain.Customer{
// 		Name:   uReqUpdate.Name,
// 		City:   uReqUpdate.City,
// 		Status: uReqUpdate.Status,
// 	}
// 	return c.repo.Update(id, temp)
// }
// func (c CustomerService) DeleteCustomer(id string) error {
// 	return c.repo.Delete(id)
// }
