package payment

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
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
	msg := fmt.Sprintf("Payment request for %v", req.OrderID)

	logger.Info(msg)

	uuid, err := uuid.NewRandom()

	if err != nil {
		logger.Errorf("error creating uuid: %v", err)
	}

	msg = fmt.Sprintf("%d paid", req.OrderID)

	err = p.publisher.Publish(context.Background(), &api.Event{
		Id:        uuid.String(),
		Timestamp: time.Now().Unix(),
		Message:   msg,
	})

	if err != nil {
		logger.Errorf("error while publishing %v", err)
	}

	return nil
}
