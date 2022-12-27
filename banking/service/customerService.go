package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
)

type ICustomerService interface {
	GetAllCustomer() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.ICustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	var cArray []dto.CustomerResponse

	for intr, c := range customers {

		_ = intr

		cArray = append(cArray, c.ToDto())
	}

	return cArray, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repo domain.ICustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
