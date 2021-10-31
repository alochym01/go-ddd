package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

// func (c CustomerRepositoryStub) FindByID(id string) (*Customer, error) {
// 	for i, v := range c.customers {
// 		if strings.Contains(v.Id, id) {
// 			return &c.customers[i], nil
// 		}
// 	}
// 	return nil, errors.New("Record Not Found")
// }
// func (c CustomerRepositoryStub) Create(temp Customer) error {
// 	return nil
// }
// func (c CustomerRepositoryStub) Update(id string, temp Customer) error {
// 	return nil
// }
// func (c CustomerRepositoryStub) Delete(id string) error {
// 	return nil
// }
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	c := []Customer{
		{Id: "1", Name: "Do Nguyen Ha", City: "Ho Chi Minh", Status: "true"},
		{Id: "2", Name: "Do Thi Hiep", City: "Ho Chi Minh", Status: "true"},
	}
	return CustomerRepositoryStub{customers: c}
}
