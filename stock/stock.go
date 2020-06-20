package stock

import (
	"context"

	"github.com/ob-vss-20ss/blatt2-cyan/api"
	catalog "github.com/ob-vss-20ss/blatt2-cyan/catalog"
)

type Stock struct {
	items map[uint32]*catalog.Item
}

func New() *Stock {
	return &Stock{
		items: make(map[uint32]*catalog.Item),
	}
}

func (c *Stock) GetItemsInStock(ctx context.Context,
	req *api.ItemsInStockRequest,
	rsp *api.ItemsInStockResponse) error {
	return nil
}

func (c *Stock) GetStockOfItem(ctx context.Context,
	req *api.StockOfItemRequest,
	rsp *api.StockOfItemResponse) error {
	return nil
}

func (c *Stock) ReduceStockOfItem(ctx context.Context,
	req *api.ReduceStockRequest,
	rsp *api.ReduceStockResponse) error {
	return nil
}

func (c *Stock) IncreaseStockOfItem(ctx context.Context,
	req *api.IncreaseStockRequest,
	rsp *api.IncreaseStockResponse) error {
	return nil
}
