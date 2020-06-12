package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/store/redis/v2"
	"github.com/vesose/example-micro/api"
	"github.com/vesose/example-micro/client"
	"github.com/vesose/example-micro/misc"
)

func main() {
	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()
	store := redis.NewStore()

	service := micro.NewService(
		micro.Registry(registry),
		micro.Store(store),
	)
	service.Init()

	client := client.New(api.NewCatalogService("catalog", service.Client()),
		api.NewOrderService("order", service.Client()),
		api.NewCustomerService("customer", service.Client()),
		service.Options().Store)

	client.Interact()
}
