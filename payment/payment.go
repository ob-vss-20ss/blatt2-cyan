package payment

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Payment struct {
	publisher micro.Event
}

func New(publisher micro.Event) *Payment {
	return &Payment{publisher: publisher}
}

func (p *Payment) ReceivePayment(ctx context.Context, req *api.PaymentRequest, res *api.PaymentResponse) error {
	msg := fmt.Sprintf("Payment request for", req.OrderID)

	logger.Info(msg)

	err := "hallo"

	x := p.publisher.Publish(context.Background(), &api.PaymentEvent{OrderID: req.OrderID})

	if x != nil {
		logger.Errorf("error while publishing %v", err)
	}

	return nil
}
