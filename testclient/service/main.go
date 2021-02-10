package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
	"github.com/ob-vss-20ss/blatt2-cyan/misc"
	"github.com/ob-vss-20ss/blatt2-cyan/testclient"
)

func main() {
	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()

	service := micro.NewService(
		micro.Registry(registry),
	)
	service.Init()

	client := testclient.New(api.NewCustomerService("customer", service.Client()),
		api.NewCatalogService("catalog", service.Client()),
		api.NewStockService("stock", service.Client()))

	client.Interact()
}
