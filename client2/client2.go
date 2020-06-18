package client2

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/store"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Client struct {
	paymentService api.PaymentService
}

func New(paymentService api.PaymentService, store store.Store) *Client {
	return &Client{
		paymentService: paymentService,
	}
}

func (c *Client) Interact() {

	_, err := c.paymentService.ReceivePayment(context.Background(), &api.PaymentRequest{
		OrderID: 1234,
	})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received:")
	}

	/*res, err := c.store.Read("sleep", func(r *store.ReadOptions) { r.Table = "sleeper" })

	sleep := 1000

	if err != nil {
		logger.Errorf("error while reading from store: %+v", err)
	} else {
		sleep, err = strconv.Atoi(string(res[0].Value))
		if err != nil {
			logger.Errorf("error while converting value from store: %+v", err)
		}
	}

	logger.Infof("sleeping %v milliseconds...", sleep)
	time.Sleep(time.Duration(sleep) * time.Millisecond)*/

}
