package catalog

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

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

type CatalogItem struct {
	ArticleID uint32
	Name      string
	Price     uint32
}

func (c *Catalog) InitData() {
	var itemsJSON []CatalogItem
	file, _ := ioutil.ReadFile("data/catalog.json")
	if err := json.Unmarshal([]byte(file), &itemsJSON); err != nil {
		panic(err)
	}
	for i, item := range itemsJSON {
		fmt.Printf("item from list, %v, %v, %v, %v\n", i, item.ArticleID, item.Name, item.Price)
	}
	for j := uint32(0); j < uint32(len(itemsJSON)); j++ {
		c.items[j+1] = &api.CatalogItem{ArticleID: itemsJSON[j].ArticleID, Name: itemsJSON[j].Name, Price: itemsJSON[j].Price}
	}
	for i, item := range c.items {
		fmt.Printf("item from map, %v, %v, %v, %v\n", i, item.ArticleID, item.Name, item.Price)
	}
}

/*func (c *Catalog) AddItems() {
	c.items[1] = &api.CatalogItem{ArticleID: 1, Name: "Tesla", Price: 1000}
	c.items[2] = &api.CatalogItem{ArticleID: 2, Name: "Mercedes", Price: 100}
	c.items[3] = &api.CatalogItem{ArticleID: 3, Name: "Mini", Price: 500}
}*/

func (c *Catalog) GetItemsInStock(ctx context.Context,
	req *api.ItemsInStockRequest,
	rsp *api.ItemsInCatalogResponse) error {
	//c.AddItems()
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

	logger.Infof("ArticleID from client: %d\n", ArticleID)

	_, ok := c.items[ArticleID]
	if ok {
		itemInStockRsp, err := c.stock.GetItem(context.Background(),
			&api.ItemRequest{
				ArticleID: ArticleID,
			})

		logger.Infof("Got item from stock: %+v", itemInStockRsp)
		logger.Infof("Got error from stock: %+v", err)

		if err != nil {
			return fmt.Errorf("item is not available in stock")
		}

		logger.Infof("Received item in stock ID: %+v",
			itemInStockRsp.GetArticleID())
		logger.Infof("Received available from stock: %+v",
			itemInStockRsp.GetAmount())
		rsp.ArticleID = c.items[itemInStockRsp.ArticleID].ArticleID
		rsp.Name = c.items[itemInStockRsp.ArticleID].Name
		rsp.Price = c.items[itemInStockRsp.ArticleID].Price
		rsp.Amount = itemInStockRsp.Amount
	} else {
		return fmt.Errorf("item is not available, non-existent ID")
	}
	return nil
}
