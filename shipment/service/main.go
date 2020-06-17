package main

import (
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	nats "github.com/micro/go-plugins/broker/nats/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/ob-vss-20ss/blatt2-cyan/misc"
	"github.com/ob-vss-20ss/blatt2-cyan/shipment"
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

	if err := micro.RegisterSubscriber("payment.*", service.Server(),
		shipment.New(micro.NewEvent("shipment.shipped", service.Client()))); err != nil {
		panic(err)
	}

	if err := service.Run(); err == nil {
		log.Fatal(err)
	}
}
