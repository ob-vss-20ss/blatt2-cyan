package testclient

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Client struct {
	customer api.CustomerService
	catalog  api.CatalogService
	stock    api.StockService
}

func New(customer api.CustomerService,
	catalog api.CatalogService,
	stock api.StockService) *Client {
	return &Client{
		customer: customer,
		catalog:  catalog,
		stock:    stock,
	}
}

//nolint:mnd
func (c *Client) Interact() {
	//Register customer ID1
	//customerID := uint32(1)
	name := "Rebel"
	address := "Grasmeierstraße, 15"

	registerRsp, err := c.customer.RegisterCustomer(context.Background(),
		&api.RegisterCustomerRequest{
			//CustomerID: customerID,
			Name:    name,
			Address: address,
		})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received added customerID: %+v",
			registerRsp.GetCustomerID())
	}
	//Register customer ID2
	name = "Toska"
	address = "Klarastraße, 8"

	registerRsp, err = c.customer.RegisterCustomer(context.Background(),
		&api.RegisterCustomerRequest{
			//CustomerID: customerID,
			Name:    name,
			Address: address,
		})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received added customerID: %+v", registerRsp.GetCustomerID())
	}
	//Get customer ID1
	customerID := uint32(1)

	getCustomerRsp, err := c.customer.GetCustomer(context.Background(),
		&api.GetCustomerRequest{
			CustomerID: customerID,
		})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received customerID: %+v",
			getCustomerRsp.GetCustomerID())
		logger.Infof("Received customer name: %+v",
			getCustomerRsp.GetName())
		logger.Infof("Received customer address: %+v",
			getCustomerRsp.GetAddress())
	}
	//Delete customer ID1
	deleteCustomerRsp, err := c.customer.DeleteCustomer(context.Background(),
		&api.DeleteCustomerRequest{
			CustomerID: customerID,
		})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received deleted customerID: %+v",
			deleteCustomerRsp.GetCustomerID())
	}
	//Get items in stock
	itemsInStockRsp, err := c.catalog.GetItemsInStock(context.Background(),
		&api.ItemsInStockRequest{})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received items in Stock: %+v",
			itemsInStockRsp.GetCatalogItems())
	}
	//Get item in stock ID1
	articleID := uint32(1)
	itemInStockRsp, err := c.catalog.GetItem(context.Background(),
		&api.ItemRequest{
			ArticleID: articleID,
		})

	if err != nil {
		logger.Infof("Item is not available")
	} else {
		logger.Infof("Received item in stock ID1: %+v",
			itemInStockRsp.GetArticleID())
		logger.Infof("Received item in stock name: %+v",
			itemInStockRsp.GetName())
		logger.Infof("Received item in stock price: %+v",
			itemInStockRsp.GetPrice())
		logger.Infof("Received item in stock available: %+v",
			itemInStockRsp.GetAmount())
	}
	//Get stock of item
	stockOfItemRsp, err := c.stock.GetStockOfItem(context.Background(),
		&api.StockOfItemRequest{
			ArticleID: articleID,
		})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received stock of item ID1: %+v",
			stockOfItemRsp.GetAmount())
	}
	//Reduce stock of item
	reduceBy := uint32(1)
	stockReduceRsp, err := c.stock.ReduceStockOfItem(context.Background(),
		&api.ReduceStockRequest{
			ArticleID: articleID,
			Amount:    reduceBy,
		})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received item ID1: %+v",
			stockReduceRsp.GetArticleID())
		logger.Infof("Received reduced stock of item ID1: %+v",
			stockReduceRsp.GetAmount())
	}
	//Increase stock of item
	increaseBy := uint32(1)
	stockIncreaseRsp, err := c.stock.IncreaseStockOfItem(context.Background(),
		&api.IncreaseStockRequest{
			ArticleID: articleID,
			Amount:    increaseBy,
		})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received item ID1: %+v",
			stockIncreaseRsp.GetArticleID())
		logger.Infof("Received increased stock of item ID1: %+v",
			stockIncreaseRsp.GetAmount())
	}
}
