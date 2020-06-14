package stock

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
	catalog "github.com/ob-vss-20ss/blatt2-cyan/catalog"
)

type Stock struct {
	items []*catalog.Item
	stock api.StockService
}

func (c *Stock) GetItemsInStock(ctx context.Context, req *api.ItemsInStockRequest, rsp *api.ItemsInStockResponse) error {
	_, err := c.stock.GetItemsInStock(context.Background(), &api.ItemsInStockRequest{})

	if err != nil {
		logger.Error(err)
	}

	return nil

}
