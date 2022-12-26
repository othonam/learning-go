package domain

import "banking/errs"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

type ICustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError)
}
