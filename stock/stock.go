package stock

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Stock struct {
	items map[uint32]*api.StockItem
}

func New() *Stock {
	return &Stock{
		items: make(map[uint32]*api.StockItem),
	}
}

func (c *Stock) AddItems() {
	c.items[1] = &api.StockItem{ArticleID: 1, Amount: 20}
	c.items[2] = &api.StockItem{ArticleID: 2, Amount: 0}
	c.items[3] = &api.StockItem{ArticleID: 3, Amount: 20}
}

func (c *Stock) GetItemsInStock(ctx context.Context,
	req *api.ItemsInStockRequest,
	rsp *api.ItemsInStockResponse) error {
	c.AddItems()
	itemList := []*api.StockItem{}
	for _, value := range c.items {
		if value.Amount != 0 {
			itemList = append(itemList, value)
		}
	}
	rsp.StockItems = itemList
	return nil
}

func (c *Stock) GetItem(ctx context.Context,
	req *api.ItemRequest,
	rsp *api.StockItem) error {

	ArticleID := req.ArticleID
	logger.Info(ArticleID)

	_, ok := c.items[ArticleID]
	if ok && c.items[ArticleID].Amount != 0 {
		rsp.ArticleID = c.items[ArticleID].ArticleID
		rsp.Amount = c.items[ArticleID].Amount
	} else {
		return fmt.Errorf("Item is not available in stock.")
	}

	/*_, ok := c.items[ArticleID]
	if ok && c.items[ArticleID].Amount != 0 {
		rsp.ArticleID = c.items[ArticleID].ArticleID
	} else {
		return fmt.Errorf("Item is not available in stock.")
	}*/

	return nil
}

func (c *Stock) GetStockOfItem(ctx context.Context,
	req *api.StockOfItemRequest,
	rsp *api.StockOfItemResponse) error {

	/*itemID := req.ArticleID
	_, ok := c.items[itemID]
	if ok {
		rsp.Amount = c.items[itemID].Amount
	}*/
	return nil
}

func (c *Stock) ReduceStockOfItem(ctx context.Context,
	req *api.ReduceStockRequest,
	rsp *api.ReduceStockResponse) error {
	/*itemID := req.ArticleID
	reduceBy := req.Amount
	_, ok := c.items[itemID]
	if ok {
		available := c.items[itemID].Amount - reduceBy
		c.items[itemID].Amount = available
		rsp.ArticleID = c.items[itemID].ArticleID
		rsp.Amount = available
	}*/
	return nil
}

func (c *Stock) IncreaseStockOfItem(ctx context.Context,
	req *api.IncreaseStockRequest,
	rsp *api.IncreaseStockResponse) error {
	/*itemID := req.ArticleID
	increaseBy := req.Amount
	_, ok := c.items[itemID]
	if ok {
		available := c.items[itemID].Amount + increaseBy
		c.items[itemID].Amount = available
		rsp.ArticleID = c.items[itemID].ArticleID
		rsp.Amount = available
	}*/
	return nil
}
