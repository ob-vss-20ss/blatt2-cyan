package apitestorder

/*
import (
	"context"
	"fmt"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	nats "github.com/micro/go-plugins/broker/nats/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
	"github.com/ob-vss-20ss/blatt2-cyan/misc"
	"github.com/ob-vss-20ss/blatt2-cyan/order"
)

type orderTest struct {
	orderService    api.OrderService
	paymentService  api.PaymentService
	shipmentService api.ShipmentService
	stockService    api.StockService
}

func New(orderService api.OrderService, paymentService api.PaymentService,
	shipmentService api.ShipmentService, stockService api.StockService) *orderTest {
	return &orderTest{
		orderService:    orderService,
		paymentService:  paymentService,
		shipmentService: shipmentService,
		stockService:    stockService,
	}
}

func (o *orderTest) TestCalculatPrice() {
	//erst mit catalog service möglich
	placeholder := uint32(0)

	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()
	broker := nats.NewBroker()

	service := micro.NewService(
		micro.Name("order"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Broker(broker),
		//micro.Store(store),
	)

	orderService := order.New(
		api.NewCatalogService("catalog", service.Client()),
		api.NewStockService("stock", service.Client()),
		api.NewCustomerService("customer", service.Client()),
		api.NewPaymentService("payment", service.Client()),
	)

	//articleListOrder []*api.ArticleWithAmount

	var articleList = []*api.ArticleWithAmount{
		{
			ArticleID: placeholder,
			Amount:    placeholder,
		},
		{
			ArticleID: placeholder,
			Amount:    placeholder,
		},
	}

	price := orderService.CalculatePrice(articleList)

	if price != placeholder {
		panic(fmt.Errorf("expected price to be placholder, but was %d", price))
	}

	//amount erhöhen
	articleList[0].Amount = placeholder

	if !orderService.CheckStock(articleList) {
		panic(fmt.Errorf("expected return value to be false, but was true"))
	}
}

func (o *orderTest) TestCheckStock() {
	//erst mit stock service möglich
	placeholder := uint32(0)

	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()
	broker := nats.NewBroker()

	service := micro.NewService(
		micro.Name("order"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Broker(broker),
		//micro.Store(store),
	)

	orderService := order.New(
		api.NewCatalogService("catalog", service.Client()),
		api.NewStockService("stock", service.Client()),
		api.NewCustomerService("customer", service.Client()),
		api.NewPaymentService("payment", service.Client()),
	)

	//articleListOrder []*api.ArticleWithAmount

	var articleList = []*api.ArticleWithAmount{
		{
			ArticleID: placeholder,
			Amount:    placeholder,
		},
		{
			ArticleID: placeholder,
			Amount:    placeholder,
		},
	}

	if !orderService.CheckStock(articleList) {
		panic(fmt.Errorf("expected return value to be true, but was false"))
	}

	//amount erhöhen
	articleList[0].Amount = placeholder

	if !orderService.CheckStock(articleList) {
		panic(fmt.Errorf("expected return value to be false, but was true"))
	}
}

func (o *orderTest) TestReduceStock() {
	// erst mit stockservice möglich
	placeholder := uint32(0)

	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()
	broker := nats.NewBroker()

	service := micro.NewService(
		micro.Name("order"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Broker(broker),
		//micro.Store(store),
	)

	orderService := order.New(
		api.NewCatalogService("catalog", service.Client()),
		api.NewStockService("stock", service.Client()),
		api.NewCustomerService("customer", service.Client()),
		api.NewPaymentService("payment", service.Client()),
	)

	//articleListOrder []*api.ArticleWithAmount

	var articleList = []*api.ArticleWithAmount{
		{
			ArticleID: placeholder,
			Amount:    placeholder,
		},
		{
			ArticleID: placeholder,
			Amount:    placeholder,
		},
	}

	orderService.ReduceStock(articleList)

	resFirstItem, err := o.stockService.GetStockOfItem(
		context.Background(), &api.StockOfItemRequest{
		ArticleID: placeholder,
	})

	if err != nil {
		panic(err)
	}

	resSecondItem, err := o.stockService.GetStockOfItem(
		context.Background(), &api.StockOfItemRequest{
		ArticleID: placeholder,
	})

	if err != nil {
		panic(err)
	}

	if resFirstItem.Amount != placeholder {
		panic(fmt.Errorf("item with articleID %d was expected to have an amount of place holder,
		but had an amount of %d", placeholder, resFirstItem.Amount))
	}

	if resSecondItem.Amount != placeholder {
		panic(fmt.Errorf("item with articleID %d was expected to have an amount of place holder,
		but had an amount of %d", placeholder, resSecondItem.Amount))
	}
}

func (o *orderTest) TestIncreaseStock() {
	// erst mit stockservice möglich
	placeholder := uint32(0)

	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()
	broker := nats.NewBroker()

	service := micro.NewService(
		micro.Name("order"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Broker(broker),
		//micro.Store(store),
	)

	orderService := order.New(
		api.NewCatalogService("catalog", service.Client()),
		api.NewStockService("stock", service.Client()),
		api.NewCustomerService("customer", service.Client()),
		api.NewPaymentService("payment", service.Client()),
	)

	//articleListOrder []*api.ArticleWithAmount

	var articleList = []*api.ArticleWithAmount{
		{
			ArticleID: placeholder,
			Amount:    placeholder,
		},
		{
			ArticleID: placeholder,
			Amount:    placeholder,
		},
	}

	orderService.IncreaseStock(articleList)

	resFirstItem, err := o.stockService.GetStockOfItem(context.Background(), &api.StockOfItemRequest{
		ArticleID: placeholder,
	})

	if err != nil {
		panic(err)
	}

	resSecondItem, err := o.stockService.GetStockOfItem(context.Background(), &api.StockOfItemRequest{
		ArticleID: placeholder,
	})

	if err != nil {
		panic(err)
	}

	if resFirstItem.Amount != placeholder {
		panic(fmt.Errorf("item with articleID %d was expected to have an amount of place holder,
		but had an amount of %d", placeholder, resFirstItem.Amount))
	}

	if resSecondItem.Amount != placeholder {
		panic(fmt.Errorf("item with articleID %d was expected to have an amount of place holder,
		but had an amount of %d", placeholder, resSecondItem.Amount))
	}
}

func (o *orderTest) TestPlaceOrder() {
	var articleList = []*api.ArticleWithAmount{
		{
			ArticleID: 1,
			Amount:    1,
		},
	}

	resPlaceOrder, err := o.orderService.PlaceOrder(context.Background(), &api.PlaceOrderRequest{
		CustomerID:  1,
		ArticleList: articleList,
	})

	if err != nil {
		panic(err)
	}

	resGetOrder, err := o.orderService.GetOrder(context.Background(), &api.GetOrderRequest{
		OrderID: resPlaceOrder.OrderID,
	})

	if err != nil {
		panic(err)
	}

	if resGetOrder.ArticleList[0].ArticleID != 1 ||
		resGetOrder.ArticleList[0].Amount != 1 {
		panic(fmt.Errorf("Expected articleList to be %v, but was %v", articleList, resGetOrder.ArticleList))
	}


	if resGetOrder.CustomerID != 1 {
		panic(fmt.Errorf("Expected customerID to be 1, but was %d", resGetOrder.CustomerID))
	}

	if resGetOrder.Paid {
		panic(fmt.Errorf("expected Paid to be false, but was true"))
	}

	if resGetOrder.Shipped {
		panic(fmt.Errorf("expected Shipped to be false, but was true"))
	}
}

func (o *orderTest) TestProcess() {
	var articleList = []*api.ArticleWithAmount{
		{
			ArticleID: 1,
			Amount:    1,
		},
	}

	resPlaceOrder, err := o.orderService.PlaceOrder(
		context.Background(), &api.PlaceOrderRequest{
		CustomerID:  1,
		ArticleList: articleList,
	})

	if err != nil {
		panic(err)
	}

	_, err = o.paymentService.ReceivePayment(
		context.Background(), &api.PaymentRequest{
		OrderID: resPlaceOrder.OrderID,
	})

	time.Sleep(1000)

	resGetOrder, err := o.orderService.GetOrder(
		context.Background(), &api.GetOrderRequest{
		OrderID: resPlaceOrder.OrderID,
	})

	if err != nil {
		panic(err)
	}

	if !resGetOrder.Paid {
		panic(fmt.Errorf("expected Paid to be true, but was false"))
	}

	if !resGetOrder.Shipped {
		panic(fmt.Errorf("expected Shipped to be true, but was false"))
	}
}
*/
