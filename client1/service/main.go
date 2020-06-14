package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
	"github.com/ob-vss-20ss/blatt2-cyan/client1"
	"github.com/ob-vss-20ss/blatt2-cyan/misc"
)

func main() {
	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()

	service := micro.NewService(
		micro.Registry(registry),
	)
	service.Init()

	client := client1.New(api.NewCatalogService("catalog", service.Client()),
		api.NewOrderService("order", service.Client()),
		api.NewCustomerService("customer", service.Client()),
		api.NewPaymentService("payment", service.Client()))

	client.Interact()
}
