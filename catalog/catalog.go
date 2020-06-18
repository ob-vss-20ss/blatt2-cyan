package catalog

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Catalog struct {
	stock api.StockService
	items []*Item
}

func New(stock api.StockService) *Catalog {
	return &Catalog{
		stock: stock,
		items: make([]*Item, 5),
	}
}

func (c *Catalog) GetItemsInStock(ctx context.Context, req *api.ItemsInStockRequest, rsp *api.ItemsInStockResponse) error {
	_, err := c.stock.GetItemsInStock(context.Background(), &api.ItemsInStockRequest{})

	if err != nil {
		logger.Error(err)
	}

	return nil

}

func (c *Catalog) GetItem(ctx context.Context, req *api.ItemRequest, rsp *api.ItemResponse) error {
	//Einzelnes Item mit gegebener Id (req.ItemID) zurückgeben

	return nil

}
