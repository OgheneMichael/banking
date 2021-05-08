package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Michael", City: "Lagos", ZipCode: "102102", DateOfBirth: "2000-02-02", Status: "1"},
		{Id: "1002", Name: "Wale", City: "Lagos", ZipCode: "103102", DateOfBirth: "2000-02-02", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}
