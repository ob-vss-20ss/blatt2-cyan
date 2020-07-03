package client2

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Client struct {
	orderService api.OrderService
}

func New(orderService api.OrderService) *Client {
	return &Client{
		orderService: orderService,
	}
}

//nolint:mnd
func (c *Client) Interact() {
	time.Sleep(2 * time.Second)

	var articleListOrder = []*api.ArticleWithAmount{
		{
			ArticleID: 1,
			Amount:    2,
		},
	}

	resOrder, err := c.orderService.PlaceOrder(context.Background(), &api.PlaceOrderRequest{
		CustomerID:  2,
		ArticleList: articleListOrder,
	})

	if err != nil {
		panic(err)
	}

	logger.Info(resOrder)

	resCancel, err := c.orderService.CancelOrder(context.Background(), &api.CancelRequest{
		CustomerID: 2,
		OrderID:    resOrder.OrderID,
	})

	if err != nil {
		panic(err)
	}

	logger.Info(resCancel)
}
