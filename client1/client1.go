package client1

import (
	"context"

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

	logger.Infof("Item in stock response (non-existent ID): %+v", itemInStockRsp)
	logger.Infof("Error message (non-existent ID): %+v ", err)

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

	//Get item in stock ID2. No items in stock-------------------------------
	//Der Kunde wählt den Artikel,
	//der im Lager nicht vorhanden ist
	articleID2 := uint32(2)
	itemInStockRsp, err = c.catalog.GetItem(context.Background(),
		&api.ItemRequest{
			ArticleID: articleID2,
		})

	logger.Info("Item in stock response (not available in stock): %+v", itemInStockRsp)
	logger.Info("Error message (not available in stock): %+v ", err)

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

	//Get item in stock ID1-------------------------------
	//Der Kunde wählt einen bestimmten Artikel,
	//der im Lager vorhanden ist
	articleID1 := uint32(1)
	itemInStockRsp, err = c.catalog.GetItem(context.Background(),
		&api.ItemRequest{
			ArticleID: articleID1,
		})

	logger.Info("Item in stock response (available): %+v", itemInStockRsp)
	logger.Info("Error message (not available in stock): %+v ", err)

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

	//Place order ID1, ID2-------------------------------
	//Kunde bestellt Artikel mit der ID1 und ID2
	//Kunde ist noch nicht registriert
	customerID := uint32(1)
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

	logger.Infof("Order, customer not registered: %+v", placeOrderRsp)
	logger.Infof("Error message (customer not registered): %+v ", err)

	if err != nil {
		logger.Error(err)
		logger.Infof("Received message: %+v",
			placeOrderRsp.GetMessage())
	} else {
		logger.Infof("Received order ID: %+v",
			placeOrderRsp.GetOrderID())
	}

	//Register customer ID1------------------------------
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
	placeOrderRegisteredRsp, err := c.order.PlaceOrder(context.Background(),
		&api.PlaceOrderRequest{
			CustomerID:  customerID,
			ArticleList: articleListOrder,
		})
	var orderID uint32

	logger.Infof("Order response, item 2 not in stock: %+v", placeOrderRegisteredRsp)
	logger.Infof("Error message (item 2 not in stock): %+v ", err)

	if err != nil {
		logger.Error(err)
		logger.Infof("Received message: %+v",
			placeOrderRegisteredRsp.GetMessage())
	} else {
		orderID = placeOrderRegisteredRsp.GetOrderID()
		logger.Infof("Received order ID: %+v",
			placeOrderRegisteredRsp.GetOrderID())
		logger.Infof("Received message: %+v",
			placeOrderRegisteredRsp.GetMessage())
	}

	//Place order ID1, ID3-------------------------------
	//Kunde bestellt Artikel mit der ID1 und ID2
	//Kunde ist bereits registriert
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

	logger.Infof("Order response, all items are in stock: %+v", placeOrderRegisteredRsp)
	logger.Infof("Error message (all items are in stock): %+v ", err)

	if err != nil {
		logger.Error(err)
		logger.Infof("Received message: %+v",
			placeOrderRegisteredRsp.GetMessage())
	} else {
		logger.Infof("Received order ID: %+v",
			placeOrderRegisteredRsp.GetOrderID())
		logger.Infof("Received message: %+v",
			placeOrderRegisteredRsp.GetMessage())
	}

	//Receive payment-----------------------------------
	//Kunde bezahlt die Bestellung
	paymentRsp, err := c.payment.ReceivePayment(context.Background(),
		&api.PaymentRequest{
			OrderID: orderID,
		})

	logger.Infof("Payment response: %+v", paymentRsp)
	logger.Infof("Error message (payment): %+v ", err)

	if err != nil {
		logger.Error(err)
	} else {
		logger.Info(paymentRsp)
	}
}

/*func (c *Client) Interact() {
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
		logger.Infof("Received: %+v", registerRsp.GetCustomerID())
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
			itemsInStockRsp.GetItems())
	}

	//Get item in stock ID1
	articleID := uint32(1)
	itemInStockRsp, err := c.catalog.GetItem(context.Background(),
		&api.ItemRequest{
			ArticleID: articleID,
		})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received item in stock ID1: %+v",
			itemInStockRsp.GetArticleID())
		logger.Infof("Received item in stock name: %+v",
			itemInStockRsp.GetName())
		logger.Infof("Received item in stock price: %+v",
			itemInStockRsp.GetPrice())
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
}*/
