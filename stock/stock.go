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
	c.items[1] = &api.Item{ArticleID: 1, Name: "Tesla", Price: 100, Amount: 20}
	c.items[2] = &api.Item{ArticleID: 2, Name: "Falcon9", Price: 1000, Amount: 20}
	c.items[3] = &api.Item{ArticleID: 3, Name: "Falcon Heavy", Price: 1000, Amount: 20}
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
	//c.AddItems()
	ArticleID := req.ArticleID
	_, ok := c.items[ArticleID]
	if ok {
		rsp.ArticleID = c.items[ArticleID].ArticleID
		rsp.Name = c.items[ArticleID].Name
		rsp.Price = c.items[ArticleID].Price
	}
	return nil
}

func (c *Stock) GetStockOfItem(ctx context.Context,
	req *api.StockOfItemRequest,
	rsp *api.StockOfItemResponse) error {
	//c.AddItems()
	itemID := req.ArticleID
	_, ok := c.items[itemID]
	if ok {
		rsp.Amount = c.items[itemID].Amount
	}
	return nil
}

func (c *Stock) ReduceStockOfItem(ctx context.Context,
	req *api.ReduceStockRequest,
	rsp *api.ReduceStockResponse) error {
	itemID := req.ArticleID
	reduceBy := req.Amount
	_, ok := c.items[itemID]
	if ok {
		available := c.items[itemID].Amount - reduceBy
		c.items[itemID].Amount = available
		rsp.ArticleID = c.items[itemID].ArticleID
		rsp.Amount = available
	}
	return nil
}

func (c *Stock) IncreaseStockOfItem(ctx context.Context,
	req *api.IncreaseStockRequest,
	rsp *api.IncreaseStockResponse) error {
	itemID := req.ArticleID
	increaseBy := req.Amount
	_, ok := c.items[itemID]
	if ok {
		available := c.items[itemID].Amount + increaseBy
		c.items[itemID].Amount = available
		rsp.ArticleID = c.items[itemID].ArticleID
		rsp.Amount = available
	}
	return nil
}
