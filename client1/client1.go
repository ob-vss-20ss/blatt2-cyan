package client1

import (
	"context"

	"github.com/coreos/etcd/store"
	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Client1 struct {
	catalog  api.CatalogService
	order    api.OrderService
	customer api.CustomerService
	payment  api.PaymentService
}

func New(catalog api.CatalogService, order api.OrderService, customer api.CustomerService, payment api.PaymentService, store store.Store) *Client1 {
	return &Client1{
		catalog:  catalog,
		order:    order,
		customer: customer,
		payment:  payment,
	}
}

func (c *Client1) Interact() {

	rsp, err := c.catalog.GetItemsInStock(context.Background(), &api.ItemsInStockRequest{})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received: %+v", rsp.GetItems())
	}
}
