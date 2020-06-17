package payment

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/store"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Payment struct {
	publisher micro.Event
}

func New(publisher micro.Event) *Payment {
	return &Payment{publisher: publisher}
}

func (p *Payment) RecievePayment(ctx context.Context, req *api.PaymentRequest, rsp *api.PaymentResponse) error {
	msg := fmt.Sprintf("Request payment for", req.orderID)

	logger.Info(msg)

	if err := p.publisher.Publish(context.Background(), &api.PaymentEvent{
		orderID: req.orderID,
	}); err != nil {
		logger.Errorf("error while publishing: %v", err)
	}

	return nil
}