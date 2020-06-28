package catalog

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Catalog struct {
	stock api.StockService
	items map[uint32]*api.CatalogItem
}

func New(stock api.StockService) *Catalog {
	return &Catalog{
		stock: stock,
		items: make(map[uint32]*api.CatalogItem),
	}
}

func (c *Catalog) AddItems() {
	c.items[1] = &api.CatalogItem{ArticleID: 1, Name: "Tesla", Price: 1000}
	c.items[2] = &api.CatalogItem{ArticleID: 2, Name: "Mercedes", Price: 100}
	c.items[3] = &api.CatalogItem{ArticleID: 3, Name: "Mini", Price: 500}
}

func (c *Catalog) GetItemsInStock(ctx context.Context,
	req *api.ItemsInStockRequest,
	rsp *api.ItemsInCatalogResponse) error {
	c.AddItems()

	itemsInStockRsp, err := c.stock.GetItemsInStock(context.Background(),
		&api.ItemsInStockRequest{})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received items in stock: %+v",
			itemsInStockRsp.GetStockItems())
	}

	itemList := []*api.CatalogItem{}

	for _, itemInStock := range itemsInStockRsp.StockItems {
		//c.items[itemInStock.ArticleID].Amount = itemInStock.Amount
		_, ok := c.items[itemInStock.ArticleID]
		if ok {
			itemList = append(itemList, c.items[itemInStock.ArticleID])
		}
	}

	rsp.CatalogItems = itemList

	return nil
}

func (c *Catalog) GetItem(ctx context.Context,
	req *api.ItemRequest,
	rsp *api.ItemResponse) error {

	ArticleID := req.ArticleID

	logger.Info(ArticleID)

	_, ok := c.items[ArticleID]
	if ok {
		itemInStockRsp, err := c.stock.GetItem(context.Background(),
			&api.ItemRequest{
				ArticleID: ArticleID,
			})

		logger.Info(itemInStockRsp)
		logger.Info(err)

		if err != nil {
			return fmt.Errorf("Item is not available in stock")
		} else {
			logger.Infof("Received item in stock ID: %+v",
				itemInStockRsp.GetArticleID())
			logger.Infof("Received available: %+v",
				itemInStockRsp.GetAmount())
			rsp.ArticleID = c.items[itemInStockRsp.ArticleID].ArticleID
			rsp.Name = c.items[itemInStockRsp.ArticleID].Name
			rsp.Price = c.items[itemInStockRsp.ArticleID].Price
			rsp.Amount = itemInStockRsp.Amount
		}
	} else {
		return fmt.Errorf("Item is not available. Non-existent ID.")
	}
	return nil
}
