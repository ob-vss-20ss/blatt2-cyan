package customer

import (
	"context"
	"fmt"

	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Customer struct {
	customers map[uint32]*Person
	lastID    uint32
}

func New() *Customer {
	return &Customer{
		customers: make(map[uint32]*Person),
	}
}

func (c *Customer) RegisterCustomer(ctx context.Context,
	req *api.RegisterCustomerRequest,
	rsp *api.RegisterCustomerResponse) error {
	name := req.Name
	address := req.Address
	customerID := c.registerCustomer(name, address)
	fmt.Printf("Custommers: %v\n", c.customers)
	fmt.Printf("Custommers: %v\n", len(c.customers))
	rsp.CustomerID = customerID
	return nil
}

func (c *Customer) GetCustomer(ctx context.Context,
	req *api.GetCustomerRequest,
	rsp *api.GetCustomerResponse) error {
	customerID := req.CustomerID
	customer, ok := c.customers[customerID]
	if ok {
		rsp.CustomerID = customer.customerID
		rsp.Name = customer.name
		rsp.Address = customer.address
	} //else {
	//CustomerNotFound?
	//}
	return nil
}

func (c *Customer) DeleteCustomer(ctx context.Context,
	req *api.DeleteCustomerRequest,
	rsp *api.DeleteCustomerResponse) error {
	customerID := req.CustomerID
	_, ok := c.customers[customerID]
	if ok {
		delete(c.customers, customerID)
	}
	fmt.Printf("Custommers: %v\n", c.customers)
	fmt.Printf("Custommers: %v\n", len(c.customers))
	rsp.CustomerID = customerID
	return nil
}

func (c *Customer) registerCustomer(name string, address string) uint32 {
	c.lastID++
	id := c.lastID
	c.customers[id] = &Person{customerID: id, name: name, address: address}
	return id
}
