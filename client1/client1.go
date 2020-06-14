package client1

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Client struct {
	catalog api.CatalogService
}

func New(catalog api.CatalogService) *Client {
	return &Client{
		catalog: catalog,
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
