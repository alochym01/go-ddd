package domain

type Customer struct {
	Id     string
	Name   string
	City   string
	Status string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	// FindByID(id string) (*Customer, error)
	// Create(c Customer) error
	// Update(id string, c Customer) error
	// Delete(id string) error
}
