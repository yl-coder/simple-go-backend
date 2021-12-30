package domain

import (
	"banking/dto"
	"banking/errs"
	"database/sql"
)

type Customer struct {
	Id string
	Name string
	City string
	Zipcode string
	DateOfBirth string `db:"date_of_birth"`
	Status string
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id: c.Id,
		Name: c.Name,
		City: c.City,
		Zipcode: c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status: c.Status,
	}
}

type CustomerRepository interface {
	FindAll(statusFilter sql.NullString) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

type CustomerRepositoryStub struct {
	customers []Customer
}

func(s CustomerRepositoryStub) FindAll() ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{{"1001", "Daniel", "Singapore", "12345", "13-11-2020", "Ok"}}
	return CustomerRepositoryStub{customers: customers}
}
