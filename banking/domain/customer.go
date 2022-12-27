package domain

import (
	"banking/dto"
	"banking/errs"
)

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

func (c Customer) statusAsText() string {
	switch {
	case c.Status == "0":
		return "inactive"
	default:
		return "active"
	}
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}

const (
	inactive = 0
	active   = 1
)
