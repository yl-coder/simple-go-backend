package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"database/sql"
)

//go:generate mockgen -source=customerService.go -destination=../mocks/service/mockCustomerService.go -package=service
type CustomerService interface {
	GetAllCustomer(statusFilter sql.NullString)([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func(c DefaultCustomerService) GetAllCustomer(statusFilter sql.NullString) ([]dto.CustomerResponse, *errs.AppError) {
	domainEntities, err := c.repo.FindAll(statusFilter)

	if err != nil {
		return nil, err
	}

	customerResponses := make([]dto.CustomerResponse, len(domainEntities))

	for i, entity := range domainEntities{
		customerResponses[i] = entity.ToDto()
	}
	return customerResponses, nil
}

func(c DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	entity, err := c.repo.ById(id)
	if err != nil {
		return nil, err
	}

	toDto := entity.ToDto()
	return &toDto, nil
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}
