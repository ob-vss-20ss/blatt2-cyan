package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	nats "github.com/micro/go-plugins/broker/nats/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/store/redis/v2"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
	"github.com/ob-vss-20ss/blatt2-cyan/misc"
	"github.com/ob-vss-20ss/blatt2-cyan/order"
)

func main() {
	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()
	broker := nats.NewBroker()
	store := redis.NewStore()

	service := micro.NewService(
		micro.Name("order"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Broker(broker),
		micro.Store(store),
	)

	service.Init()

	orderService := order.New(
		api.NewCatalogService("catalog", service.Client()),
		api.NewStockService("stock", service.Client()),
		api.NewCustomerService("customer", service.Client()),
	)

	if err := api.RegisterOrderHandler(service.Server(),
		orderService); err != nil {
		log.Fatal(err)
	}

	if err := micro.RegisterSubscriber("payment.*", service.Server(),
		orderService.Process); err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := micro.RegisterSubscriber("shipment.*", service.Server(),
		orderService.Process); err != nil {
		panic(err)
	}

	if err := service.Run(); err == nil {
		log.Fatal(err)
	}

}
