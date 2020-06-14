package catalog

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Catalog struct {
	items []*Item
	stock api.StockService
}

func New(stock api.StockService) *Catalog {
	return &Catalog{
		stock: stock,
	}
}

func (c *Catalog) GetItemsInStock(ctx context.Context, req *api.ItemsInStockRequest, rsp *api.ItemsInStockResponse) error {
	_, err := c.stock.GetItemsInStock(context.Background(), &api.ItemsInStockRequest{})

	if err != nil {
		logger.Error(err)
	}

	return nil

}
