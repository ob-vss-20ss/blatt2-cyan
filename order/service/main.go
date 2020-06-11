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

	service := micro.NewService(
		micro.Name("order"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Broker(broker),
	)

	service.Init()

	if err := api.RegisterOrderHandler(service.Server(),
		order.New(micro.NewEvent("log.order", service.Client()))); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
