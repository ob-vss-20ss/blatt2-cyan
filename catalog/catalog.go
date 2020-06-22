package catalog

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Catalog struct {
	stock api.StockService
	items map[uint32]*api.Item
}

func New(stock api.StockService) *Catalog {
	return &Catalog{
		stock: stock,
		items: make(map[uint32]*api.Item),
	}
}

func (c *Catalog) AddItems() {
	c.items[1] = &api.Item{ArticleID: 1, Name: "Tesla", Price: 100, Amount: 3}
	c.items[2] = &api.Item{ArticleID: 2, Name: "Falcon9", Price: 1000, Amount: 1}
	c.items[3] = &api.Item{ArticleID: 3, Name: "FalconHeavy", Price: 1000, Amount: 2}
}

func (c *Catalog) GetItemsInStock(ctx context.Context,
	req *api.ItemsInStockRequest,
	rsp *api.ItemsInStockResponse) error {
	c.AddItems()

	itemsInStockRsp, err := c.stock.GetItemsInStock(context.Background(),
		&api.ItemsInStockRequest{})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received items in stock: %+v",
			itemsInStockRsp.GetItems())
	}

	rsp.Items = itemsInStockRsp.GetItems()

	return nil
}

func (c *Catalog) GetItem(ctx context.Context,
	req *api.ItemRequest,
	rsp *api.ItemResponse) error {
	//c.AddItems()

	ArticleID := req.ArticleID
	_, ok := c.items[ArticleID]
	if ok {
		itemInStockRsp, err := c.stock.GetItem(context.Background(),
			&api.ItemRequest{
				ArticleID: ArticleID,
			})

		if err != nil {
			logger.Error(err)
		} else {
			logger.Infof("Received item in stock ID: %+v",
				itemInStockRsp.GetArticleID())
			logger.Infof("Received item in stock name: %+v",
				itemInStockRsp.GetName())
			logger.Infof("Received item in stock price: %+v",
				itemInStockRsp.GetPrice())
		}

		rsp.ArticleID = itemInStockRsp.GetArticleID()
		rsp.Name = itemInStockRsp.GetName()
		rsp.Price = itemInStockRsp.GetPrice()
	}
	return nil
}
