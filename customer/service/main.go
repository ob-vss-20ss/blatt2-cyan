package main

import (
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
	"github.com/ob-vss-20ss/blatt2-cyan/customer"
	"github.com/ob-vss-20ss/blatt2-cyan/misc"
)

func main() {
	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()
	//broker := nats.NewBroker()
	//store := redis.NewStore()

	service := micro.NewService(
		micro.Name("order"),
		micro.Version("latest"),
		micro.Registry(registry),
		//micro.Broker(broker),
		//micro.Store(store),
	)

	service.Init()

	customerService := customer.New()

	if err := api.RegisterCustomerHandler(service.Server(),
		customerService); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err == nil {
		log.Fatal(err)
	}

}
