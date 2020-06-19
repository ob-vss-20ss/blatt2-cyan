package client1

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Client struct {
	customer api.CustomerService
}

func New(customer api.CustomerService) *Client {
	return &Client{
		customer: customer,
	}
}

func (c *Client) Interact() {
	//customerID := uint32(1)
	name := "Rebel"
	address := "Grasmeierstraße, 15"

	registerRsp, err := c.customer.RegisterCustomer(context.Background(), &api.RegisterCustomerRequest{
		//CustomerID: customerID,
		Name:    name,
		Address: address,
	})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received added customerID: %+v", registerRsp.GetCustomerID())
	}

	name = "Toska"
	address = "Klarastraße, 8"

	registerRsp, err = c.customer.RegisterCustomer(context.Background(), &api.RegisterCustomerRequest{
		//CustomerID: customerID,
		Name:    name,
		Address: address,
	})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received: %+v", registerRsp.GetCustomerID())
	}

	customerID := uint32(1)

	getCustomerRsp, err := c.customer.GetCustomer(context.Background(), &api.GetCustomerRequest{
		CustomerID: customerID,
	})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received customerID: %+v", getCustomerRsp.GetCustomerID())
		logger.Infof("Received customer name: %+v", getCustomerRsp.GetName())
		logger.Infof("Received customer address: %+v", getCustomerRsp.GetAddress())
	}

	deleteCustomerRsp, err := c.customer.DeleteCustomer(context.Background(), &api.DeleteCustomerRequest{
		CustomerID: customerID,
	})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received deleted customerID: %+v", deleteCustomerRsp.GetCustomerID())
	}
}
