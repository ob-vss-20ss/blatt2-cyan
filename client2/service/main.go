package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
	"github.com/ob-vss-20ss/blatt2-cyan/misc"
)

func main() {
	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()
	//store := redis.NewStore()

	service := micro.NewService(
		micro.Registry(registry),
		//micro.Store(store),
	)
	service.Init()

	client := apitestOrder.New(
		api.NewOrderService("order", service.Client()),
		api.NewPaymentService("payment", service.Client()),
		api.NewShipmentService("shipment", service.Client()),
		api.NewStockService("stock", service.Client()),
	)

	client.TestCalculatePrice()
	client.TestCheckStock()
	client.TestReduceStock()
	client.TestIncreaseStock()
	client.TestPlaceOrder()
	client.TestProcess()
}
