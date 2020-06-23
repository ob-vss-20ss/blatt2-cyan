package client4

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Client struct {
	orderService   api.OrderService
	paymentService api.PaymentService
}

func New(orderService api.OrderService, paymentService api.PaymentService) *Client {
	return &Client{
		orderService:   orderService,
		paymentService: paymentService,
	}
}

//nolint:mnd
func (c *Client) Interact() {
	var articleListOrder = []*api.ArticleWithAmount{
		{
			ArticleID: 1,
			Amount:    2,
		},
		{
			ArticleID: 3,
			Amount:    3,
		},
	}

	resOrder, err := c.orderService.PlaceOrder(context.Background(), &api.PlaceOrderRequest{
		CustomerID:  1,
		ArticleList: articleListOrder,
	})

	if err != nil {
		panic(err)
	}

	logger.Info(resOrder)

	resPay, err := c.paymentService.ReceivePayment(context.Background(), &api.PaymentRequest{
		OrderID: resOrder.OrderID,
	})

	if err != nil {
		panic(err)
	}

	logger.Info(resPay)

	var articleListReturn = []*api.ArticleWithAmount{
		{
			ArticleID: 1,
			Amount:    2,
		},
	}

	resReturn, err := c.orderService.ReturnItem(context.Background(), &api.ReturnRequest{
		CustomerID:  1,
		OrderID:     resOrder.OrderID,
		Replacement: false,
		ArticleList: articleListReturn,
	})

	if err != nil {
		panic(err)
	}

	logger.Info(resReturn)
}
