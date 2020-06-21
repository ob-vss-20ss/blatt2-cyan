package stock

import (
	"context"

	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Stock struct {
	items map[uint32]*api.Item
}

func New() *Stock {
	return &Stock{
		items: make(map[uint32]*api.Item),
	}
}

func (c *Stock) AddItems() {
	c.items[1] = &api.Item{ItemID: 1, Name: "Tesla", Price: 100, Available: 3}
	c.items[2] = &api.Item{ItemID: 2, Name: "Falcon9", Price: 1000, Available: 1}
}

func (c *Stock) GetItemsInStock(ctx context.Context,
	req *api.ItemsInStockRequest,
	rsp *api.ItemsInStockResponse) error {
	c.AddItems()
	itemList := []*api.Item{}
	for _, value := range c.items {
		itemList = append(itemList, value)
	}
	rsp.Items = itemList
	return nil
}

func (c *Stock) GetItem(ctx context.Context,
	req *api.ItemRequest,
	rsp *api.ItemResponse) error {
	c.AddItems()
	itemID := req.ItemID
	_, ok := c.items[itemID]
	if ok {
		rsp.ItemID = c.items[itemID].ItemID
		rsp.Name = c.items[itemID].Name
		rsp.Price = c.items[itemID].Price
	}
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
