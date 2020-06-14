package client1

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Client struct {
	catalog  api.CatalogService
	order    api.OrderService
	customer api.CustomerService
	payment  api.PaymentService
}

func New(catalog api.CatalogService, order api.OrderService, customer api.CustomerService, payment api.PaymentService) *Client {
	return &Client{
		catalog:  catalog,
		order:    order,
		customer: customer,
		payment:  payment,
	}
}

func (c *Client) Interact() {

	rsp, err := c.catalog.GetItemsInStock(context.Background(), &api.ItemsInStockRequest{})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received: %+v", rsp.GetItems())
	}
}
