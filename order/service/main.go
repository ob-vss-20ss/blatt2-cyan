package main

import (
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	nats "github.com/micro/go-plugins/broker/nats/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
	"github.com/ob-vss-20ss/blatt2-cyan/misc"
	"github.com/ob-vss-20ss/blatt2-cyan/order"
)

func main() {
	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()
	broker := nats.NewBroker()
	//store := redis.NewStore()

	service := micro.NewService(
		micro.Name("shipment"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Broker(broker),
		//micro.Store(store),
	)

	service.Init()

	orderService := order.New()

	if err := micro.RegisterSubscriber("payment.*", service.Server(),
		orderService); err != nil {
		panic(err)
	}

	if err := micro.RegisterSubscriber("payment.*", service.Server(),
		orderService); err != nil {
		panic(err)
	}

	if err := api.RegisterOrderHandler(service.Server(),
		orderService); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err == nil {
		log.Fatal(err)
	}
}
