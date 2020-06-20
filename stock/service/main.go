package main

import (
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
	"github.com/ob-vss-20ss/blatt2-cyan/misc"
	"github.com/ob-vss-20ss/blatt2-cyan/stock"
)

func main() {
	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()

	service := micro.NewService(
		micro.Name("stock"),
		micro.Version("latest"),
		micro.Registry(registry),
	)

	service.Init()

	stockService := stock.New()

	if err := api.RegisterStockHandler(service.Server(),
		stockService); err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
