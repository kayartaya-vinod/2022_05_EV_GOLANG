# Mini project

Create a new Go project that allows the user to perform CRUD operations on customer data. It is expected to be a console based menu driven application. When the application is run, the user is presented with the following menu options:

1. List all customers
2. Search customer by id
3. Search customers by name
4. Search customer by email address
5. Search customer by phone number
6. Search customers by city
7. Add a new customer
8. Update customer details
9. Delete customer details
10. Exit

The customer type has the following fields:

1. Id
2. Name
3. Email
4. Phone
5. Address
    1. Street
    2. City
    3. State
    4. Pincode
    5. Country

Create an interface `CustomerManager` with the following methods:

1. add(customer Customer)
2. update(customer Customer)
3. delete(id int)
4. findById(id int)
5. findByName(name string)
6. findByEmail(email string)
7. findByPhone(phone string)
8. findByCity(city string)

Create a struct `JsonCustomerManager` and provide implementations of the above methods. The customer data should be preserved in a JSON file.

Try to keep functions well organized and keep them short. Handle all possible erros such that the application does not break at any point in time.
