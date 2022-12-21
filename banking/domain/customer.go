package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type ICustomerRepository interface {
	FindAll() ([]Customer, error)
}
