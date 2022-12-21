package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			Id:          "1",
			Name:        "John",
			City:        "London",
			Zipcode:     "12345",
			DateofBirth: "12/12/1990"},
		{
			Id:          "2",
			Name:        "Jane",
			City:        "London",
			Zipcode:     "12345",
			DateofBirth: "12/12/1991"},
	}

	return CustomerRepositoryStub{customers}
}
