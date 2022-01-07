package domain

// Adapter

// this interface should implement FindAll()
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Nicholas", City: "New York", Zipcode: "111001", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "1001", Name: "Anthony", City: "New York", Zipcode: "111001", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "1001", Name: "Tommy", City: "New York", Zipcode: "111001", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "1001", Name: "Rob", City: "New York", Zipcode: "111001", DateofBirth: "2000-01-01", Status: "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}
