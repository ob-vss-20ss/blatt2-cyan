package order

import (
	"testing"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	nats "github.com/micro/go-plugins/broker/nats/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
	"github.com/ob-vss-20ss/blatt2-cyan/misc"
)

func articleListEqual(expected []*api.ArticleWithAmount, actual []*api.ArticleWithAmount) bool {
	for i := range expected {
		if expected[i].ArticleID != actual[i].ArticleID || expected[i].Amount != actual[i].Amount {
			return false
		}
	}
	return true
}

func TestExtractOrderIDFromMsg(t *testing.T) {
	msg := "1234 payed"
	actual := ExtractEventMsg(msg)

	expected := "payed"

	if expected != actual {
		t.Errorf("Expected %v, but was %v", expected, actual)
	}
}

func TestCalculatPrice(t *testing.T) {
	//erst mit catalog service möglich
}

func TestCheckStock(t *testing.T) {
	//erst mit stock service möglich
}

func TestReduceStock(t *testing.T) {
	//erst mit Stock service möglich
}

func TestIncreaseStock(t *testing.T) {
	// erst mit stockservice möglich
}

// nolint:funlen
func TestOrderContainsArticles(t *testing.T) {
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

	orderService := New(
		api.NewCatalogService("catalog", service.Client()),
		api.NewStockService("stock", service.Client()),
		api.NewCustomerService("customer", service.Client()),
		api.NewPaymentService("payment", service.Client()),
	)

	//articleListOrder []*api.ArticleWithAmount

	var articleListOrder = []*api.ArticleWithAmount{
		{
			ArticleID: 1,
			Amount:    5,
		},
		{
			ArticleID: 2,
			Amount:    3,
		},
		{
			ArticleID: 3,
			Amount:    8,
		},
	}

	var articleListReturn = []*api.ArticleWithAmount{
		{
			ArticleID: 1,
			Amount:    4,
		},
		{
			ArticleID: 2,
			Amount:    3,
		},
	}

	if !orderService.OrderContainsArticle(articleListOrder, articleListReturn) {
		t.Errorf("Expected true but was false")
	}

	articleListReturn[0].Amount = 6

	if orderService.OrderContainsArticle(articleListOrder, articleListReturn) {
		t.Errorf("Expected false but was true")
	}

	articleListReturn[0].ArticleID = 6
	articleListReturn[0].Amount = 4

	if orderService.OrderContainsArticle(articleListOrder, articleListReturn) {
		t.Errorf("Expected false but was true")
	}
}

// nolint:funlen
func TestShortenOrder(t *testing.T) {
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

	orderService := New(
		api.NewCatalogService("catalog", service.Client()),
		api.NewStockService("stock", service.Client()),
		api.NewCustomerService("customer", service.Client()),
		api.NewPaymentService("payment", service.Client()),
	)

	//articleListOrder []*api.ArticleWithAmount

	var articleListOrder = []*api.ArticleWithAmount{
		{
			ArticleID: 1,
			Amount:    5,
		},
		{
			ArticleID: 2,
			Amount:    3,
		},
		{
			ArticleID: 3,
			Amount:    8,
		},
	}

	var articleListReturn = []*api.ArticleWithAmount{
		{
			ArticleID: 1,
			Amount:    4,
		},
		{
			ArticleID: 2,
			Amount:    3,
		},
	}

	var expected = []*api.ArticleWithAmount{
		{
			ArticleID: 1,
			Amount:    1,
		},
		{
			ArticleID: 3,
			Amount:    8,
		},
	}

	actual := orderService.ShortenOrder(articleListOrder, articleListReturn)

	if !articleListEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}

}
