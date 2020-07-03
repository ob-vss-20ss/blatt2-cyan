package customer

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/micro/go-micro/v2/logger"
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

type Person struct {
	CustomerID uint32
	Name       string
	Address    string
}

func (c *Customer) InitData() {
	var itemsJSON []Person
	file, _ := ioutil.ReadFile("data/customers.json")
	if err := json.Unmarshal([]byte(file), &itemsJSON); err != nil {
		panic(err)
	}
	for i, item := range itemsJSON {
		fmt.Printf("item from list, %v, %v, %v, %v\n", i, item.CustomerID, item.Name, item.Address)
	}

	c.lastID = itemsJSON[len(itemsJSON)-1].CustomerID

	for j := uint32(0); j < uint32(len(itemsJSON)); j++ {
		c.customers[j+1] = &Person{CustomerID: itemsJSON[j].CustomerID,
			Name: itemsJSON[j].Name, Address: itemsJSON[j].Address}
	}
	for i, item := range c.customers {
		fmt.Printf("item from map, %v, %v, %v, %v\n", i, item.CustomerID, item.Name, item.Address)
	}
}

func (c *Customer) RegisterCustomer(ctx context.Context,
	req *api.RegisterCustomerRequest,
	rsp *api.RegisterCustomerResponse) error {
	name := req.Name
	address := req.Address
	customerID := c.registerCustomer(name, address)

	logger.Infof("Number of registered custommers: %d\n", len(c.customers))

	rsp.CustomerID = customerID
	return nil
}

func (c *Customer) GetCustomer(ctx context.Context,
	req *api.GetCustomerRequest,
	rsp *api.GetCustomerResponse) error {
	customerID := req.CustomerID
	customer, ok := c.customers[customerID]
	if ok {
		rsp.CustomerID = customer.CustomerID
		rsp.Name = customer.Name
		rsp.Address = customer.Address
	} else {
		return fmt.Errorf("customer not registered")
	}
	return nil
}

func (c *Customer) DeleteCustomer(ctx context.Context,
	req *api.DeleteCustomerRequest,
	rsp *api.DeleteCustomerResponse) error {
	customerID := req.CustomerID
	customer, ok := c.customers[customerID]
	if ok {
		delete(c.customers, customerID)
	}

	logger.Infof("Number of registered custommers: %d\n", len(c.customers))

	rsp.CustomerID = customer.CustomerID
	return nil
}

func (c *Customer) registerCustomer(name string, address string) uint32 {
	c.lastID++
	id := c.lastID
	c.customers[id] = &Person{CustomerID: id, Name: name, Address: address}
	return id
}
