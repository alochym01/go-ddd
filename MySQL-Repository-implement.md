## How to Code

### Domain Object

1. Create a Client connection Object
2. Define `CustomerRepositoryMySQL Object`
   1. **Define an attribute which has a mysql Client connection object**
   2. Implement all methods in `CustomerRepository`
      1. FindAll() ([]Customer, error)
      2. FindByID(id string) (*Customer, error)
      3. Create(c Customer) error
      4. Update(id string, c Customer) error
      5. Delete(id string) error
   3. **Define a New function to initial a CustomerRepositoryMySQL Object**

### Business Logic - Service Object

1. Define `ServiceRepository Object`
2. Validation `UserRequestObject` input // TODO check again in Handlers Object
3. Implement all methods in `ServiceRepository Object`
   1. GetAllCustomers() ([]domain.Customer, error)
   2. GetCustomer(id string) (*domain.Customer, error)
   3. CreateCustomer(uRequest UserRequestObject) error
   4. UpdateCustomer(id string, uRequest UserRequestUpdateObject) error
   5. DeleteCustomer(id string) error
4. **Transform Domain Object into Response Object - DTO**
5. **Define a New function to initial a CustomerRepositoryMySQL Object**

### Validation User Request - Handlers Object

1. Get all Customers
2. Get Customer by ID
   1. URL should only accept numberic ids
3. Create Customer
   1. Method is POST
   2. User Request body
      1. Name
      2. City
      3. Status
   3. **Transform user request body into UserRequestObject - DTO**
4. Update Customer
   1. Method is PUT
   2. URL should only accept numberic ids
   3. User Request body
      1. Name
      2. City
      3. Status
   4. **Transform user request body into UserRequestUpdateObject - DTO**
5. Delete Customer
   1. Method is DELETE
   2. URL should only accept numberic ids

### Error Handling

1. Return 404 if customer record is `not found`
2. Return 500 if server internal error

### Logging Handling

1. zap logging - structure logging
2. log standard library