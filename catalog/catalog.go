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
	c.items[1] = &api.Item{ItemID: 1, Name: "Tesla", Price: 100, Available: 3}
	c.items[2] = &api.Item{ItemID: 2, Name: "Falcon9", Price: 1000, Available: 1}
	c.items[3] = &api.Item{ItemID: 3, Name: "FalconHeavy", Price: 1000, Available: 2}
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
	c.AddItems()

	itemID := req.ItemID
	_, ok := c.items[itemID]
	if ok {
		itemInStockRsp, err := c.stock.GetItem(context.Background(),
			&api.ItemRequest{
				ItemID: itemID,
			})

		if err != nil {
			logger.Error(err)
		} else {
			logger.Infof("Received item in stock ID: %+v",
				itemInStockRsp.GetItemID())
			logger.Infof("Received item in stock name: %+v",
				itemInStockRsp.GetName())
			logger.Infof("Received item in stock price: %+v",
				itemInStockRsp.GetPrice())
		}

		rsp.ItemID = itemInStockRsp.GetItemID()
		rsp.Name = itemInStockRsp.GetName()
		rsp.Price = itemInStockRsp.GetPrice()
	}
	//Einzelnes Item mit gegebener Id (req.ItemID) zurückgeben

	return nil
}
