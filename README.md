# Go Ticket Booking client (GRPC API for Train Ticket Booking Client)
This Repo will have a golang client for ticket booking grpc server . 

## List of features

### Purchase Ticket PurchaseTicket
* Call Booking for User (firstName, lastName, Section A or B, Price, Source, Destination)
* Store Booking history in memory

### Generate Booking Receipt 
* Get Booking details against ReceiptId

###  Get List of Users for each sections
* Send List of Users and booked seats section wise 
* Fetch and send details of User Seat mapping

###  Modify Seat of User
* Modify the seat of user if it is available
* Update the booking

###  Remove User from Train (Cancel Booking)
* Modify the seat of user if it is available
* Update the booking

# Running Locally using Docker

* git clone https://github.com/AmitKannaujiya/GoClientGrpc.git
```bash
cd <root path>
go run main.go

```

access site on: http://localhost:5001/

# Documentation

* [API Docs](pending)
* [Model Design](pending)
* [Code Structure](pending)

# TODO

- [ ] Update command code with smaller modules
- [ ] Add tests for more cases
- [ ] Add tests for helpers functions
- [ ] Store results of API
- [ ] Add documentation 
- [ ] Add Authentication
- [ ] Update Logging
