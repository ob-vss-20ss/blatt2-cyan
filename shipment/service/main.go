package main

import (
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	nats "github.com/micro/go-plugins/broker/nats/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
	"github.com/ob-vss-20ss/blatt2-cyan/misc"
	"github.com/ob-vss-20ss/blatt2-cyan/shipment"
)

func main() {
	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()
	broker := nats.NewBroker()

	service := micro.NewService(
		micro.Name("shipment"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Broker(broker),
	)

	service.Init()

	if err := api.RegisterShipmentHandler(service.Server(),
		shipment.New(micro.NewEvent("log.shipment", service.Client()))); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
