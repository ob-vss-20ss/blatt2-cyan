package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
	"github.com/ob-vss-20ss/blatt2-cyan/catalog"
	"github.com/ob-vss-20ss/blatt2-cyan/misc"
)

func main() {
	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()

	service := micro.NewService(
		micro.Name("catalog"),
		micro.Version("latest"),
		micro.Registry(registry),
	)

	service.Init()

	stock := micro.NewService()
	stock.Init()

	catalogService := catalog.New(api.NewStockService("stock", stock.Client()))

	catalogService.InitData()

	if err := api.RegisterCatalogHandler(service.Server(), catalogService); err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
