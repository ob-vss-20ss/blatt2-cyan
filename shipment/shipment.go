package shipment

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Shipment struct {
	publisher micro.Event
}

func New(publisher micro.Event) *Shipment {
	return &Shipment{publisher: publisher}
}

func (p *Shipment) Process(ctx context.Context, event *api.PaymentEvent) error {
	//abfragen der adresse bei order service -> customerservice

	msg := fmt.Sprintf("Received payment event for", event.orderID)

	logger.Info(msg)

	if err := p.publisher.Publish(context.Background(), &api.ShipmentEvent{
		orderID: event.orderID,
	}); err != nil {
		logger.Errorf("error while publishing: %v", err)
	}

	return nil
}