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

	service := micro.NewService(
		micro.Name("customer"),
		micro.Version("latest"),
		micro.Registry(registry),
	)

	service.Init()

	if err := api.RegisterCustomerHandler(service.Server(),
		customer.New(micro.NewEvent("log.customer", service.Client()))); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
