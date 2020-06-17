package payment

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Payment struct {
	publisher micro.Event
}

func New(publisher micro.Event) *Payment {
	return &Payment{publisher: publisher}
}

func (p *Payment) ReceivePayment(ctx context.Context, req *api.PaymentRequest,*api.PaymentResponse) error {
	msg := fmt.Sprintf("Payment request for", req.OrderID)

	logger.info(msg)

	if err := p.publisher.Publish(context.Background(), &api.PaymentEvent) {
		OrderID: req.OrderID
	}; err != nil {
		logger.Errorf("error while publishing %v", err)
	}
}