package client1

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Client struct {
	customer api.CustomerService
	catalog  api.CatalogService
	order    api.OrderService
	payment  api.PaymentService
	shipment api.ShipmentService
}

func New(customer api.CustomerService,
	catalog api.CatalogService,
	order api.OrderService,
	payment api.PaymentService,
	shipment api.ShipmentService) *Client {
	return &Client{
		customer: customer,
		catalog:  catalog,
		order:    order,
		payment:  payment,
		shipment: shipment,
	}
}

//nolint:mnd
func (c *Client) Interact() {
	time.Sleep(1 * time.Second)
	//Get items in stock-----------------------------------
	//Betrachten des Angebots
	//Dem Kunden werden nur die Artikel angezeigt,
	//die im Lager sind
	itemsInStockRsp, err := c.catalog.GetItemsInStock(context.Background(),
		&api.ItemsInStockRequest{})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received items in Stock: %+v",
			itemsInStockRsp.GetCatalogItems())
	}

	//Get item in stock ID5. Non-existent ID-------------------------------
	//Der Kunde vertippt sich und gibt eine
	//nichtexistierende ID ein
	articleID5 := uint32(5)
	itemInStockRsp, err := c.catalog.GetItem(context.Background(),
		&api.ItemRequest{
			ArticleID: articleID5,
		})

	//logger.Infof("Item in stock response (non-existent ID): %+v", itemInStockRsp)
	//logger.Infof("Error message (non-existent ID): %+v ", err)

	if err != nil {
		logger.Error(err)
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

	//Get item in stock ID2. No items in stock-------------------------------
	//Der Kunde wählt den Artikel,
	//der im Lager nicht vorhanden ist
	articleID2 := uint32(2)
	itemInStockRsp, err = c.catalog.GetItem(context.Background(),
		&api.ItemRequest{
			ArticleID: articleID2,
		})

	//logger.Infof("Item in stock response (not available in stock): %+v", itemInStockRsp)
	//logger.Infof("Error message (not available in stock): %+v", err)

	if err != nil {
		logger.Error(err)
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

	//Get item in stock ID1-------------------------------
	//Der Kunde wählt einen bestimmten Artikel,
	//der im Lager vorhanden ist
	articleID1 := uint32(1)
	itemInStockRsp, err = c.catalog.GetItem(context.Background(),
		&api.ItemRequest{
			ArticleID: articleID1,
		})

	//logger.Infof("Item in stock response (available): %+v", itemInStockRsp)
	//logger.Infof("Error message (not available in stock): %+v", err)

	if err != nil {
		logger.Error(err)
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

	//Place order ID1, ID2-------------------------------
	//Kunde bestellt Artikel mit der ID1 und ID2
	//Kunde ist noch nicht registriert
	customerID := uint32(5)
	articleListOrder := []*api.ArticleWithAmount{
		{
			ArticleID: 1,
			Amount:    2,
		},
		{
			ArticleID: 2,
			Amount:    1,
		},
	}

	placeOrderRsp, err := c.order.PlaceOrder(context.Background(),
		&api.PlaceOrderRequest{
			CustomerID:  customerID,
			ArticleList: articleListOrder,
		})

	//logger.Infof("Order response (customer not registered): %+v", placeOrderRsp)
	//logger.Infof("Error message (customer not registered): %+v ", err)

	if err != nil {
		logger.Infof("Received message: %+v",
			placeOrderRsp.GetMessage())
		logger.Error(err)
	} else {
		logger.Infof("Received order ID: %+v",
			placeOrderRsp.GetOrderID())
	}

	//Register customer ID5------------------------------
	//Kunde registriert sich
	name := "Rebel"
	address := "Grasmeierstraße, 15"

	registerRsp, err := c.customer.RegisterCustomer(context.Background(),
		&api.RegisterCustomerRequest{
			Name:    name,
			Address: address,
		})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received added customerID: %+v",
			registerRsp.GetCustomerID())
	}

	//Place order ID1, ID2-------------------------------
	//Kunde bestellt Artikel mit der ID1 und ID2
	//Kunde ist bereits registriert
	//Artikel mit ID2 ist nicht im stock
	placeOrderRegisteredRsp, err := c.order.PlaceOrder(context.Background(),
		&api.PlaceOrderRequest{
			CustomerID:  customerID,
			ArticleList: articleListOrder,
		})
	var orderID uint32

	//logger.Infof("Order response, item 2 not in stock: %+v", placeOrderRegisteredRsp)
	//logger.Infof("Error message (item 2 not in stock): %+v ", err)

	if err != nil {
		logger.Infof("Received message: %+v",
			placeOrderRegisteredRsp.GetMessage())
		logger.Error(err)
	} else {
		orderID = placeOrderRegisteredRsp.GetOrderID()
		logger.Infof("Received order ID: %+v",
			placeOrderRegisteredRsp.GetOrderID())
	}

	//Place order ID1, ID3-------------------------------
	//Kunde bestellt Artikel mit der ID1 und ID3
	//Kunde ist bereits registriert
	//Beite Artikel sind im Stock vorhanden
	articleListOrder = []*api.ArticleWithAmount{
		{
			ArticleID: 1,
			Amount:    5,
		},
		{
			ArticleID: 3,
			Amount:    5,
		},
	}

	placeOrderRegisteredRsp, err = c.order.PlaceOrder(context.Background(),
		&api.PlaceOrderRequest{
			CustomerID:  customerID,
			ArticleList: articleListOrder,
		})

	//logger.Infof("Order response, all items are in stock: %+v", placeOrderRegisteredRsp)
	//logger.Infof("Error message (all items are in stock): %+v ", err)

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received order ID: %+v",
			placeOrderRegisteredRsp.GetOrderID())
		logger.Infof("Received message: %+v",
			placeOrderRegisteredRsp.GetMessage())
	}

	//Receive payment-----------------------------------
	//Kunde bezahlt die Bestellung
	_, err = c.payment.ReceivePayment(context.Background(),
		&api.PaymentRequest{
			OrderID: orderID,
		})
}
