# Learning Web API with GO
## Hexagonal Architecture
### Implementing Code
#### Folder Structure
#### How to Code
1. Requirements
   1. Implementing a web server
      1. User request to get all customers in:
         1. memory database
         2. mysql database
         3. redis database
         4. memcached database
         5. elastic search database
2. Step to code
   1. Define a Domain Object - [`Customer`](domain/customer.go)
      1. Define an Interface to get all customers - [`CustomerRepository`](domain/customer.go)
         1. Define [`a Storage Object`](domain/customerRepositoryStub.go) to implement all methods in the Interface - `CustomerRepository`
            1. A `Storage Object`:
               1. memory database
               2. mysql database
               3. redis database
               4. memcached database
               5. elastic search database
            2. **Define a New function to initial a Storage Object**
   2. Define a Domain Business - [`ServiceCustomer`](service/customer.go)
      1. Define an Interface to get all customer - `ServiceCustomerRepository`
      2. Define [`a Service Object`](service/customer.go):
         1. get all customers from `Storage Object` by link to `CustomerRepository`
         2. Define all methods in the Interface - `ServiceCustomerRepository`
         3. **Define a New function to initial a Service Object**
      3. **Transform Domain Object into Response Object**
   3. Define a Handler that handles user request
      1. Define [`CustomerHandler Object`](app/handler.go):
         1. Define a method which handles user request
         2. get all customers from `Storage Object` by link to `ServiceCustomerRepository`
         3. **Define a New function to initial a CustomerHanddler Object**
      2. **Transform User Request data into Request Object**
   4. Define [`a route`](app/app.go) which route user request to a handler
