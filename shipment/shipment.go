package shipment

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Shipment struct {
	publisher       micro.Event
	orderService    api.OrderService
	customerService api.CustomerService
}

func New(publisher micro.Event, orderService api.OrderService, customerService api.CustomerService) *Shipment {
	return &Shipment{
		publisher:       publisher,
		orderService:    orderService,
		customerService: customerService,
	}
}

func ExtractOrderIDFromMsg(msg string) (uint32, error) {
	s := strings.Split(msg, " ")

	tmp, err := strconv.ParseInt(s[0], 10, 32)

	orderID := uint32(tmp)

	return orderID, err
}

func (p *Shipment) Process(ctx context.Context, event *api.Event) error {
	//abfragen der adresse bei order service -> customerservice

	orderID, err := ExtractOrderIDFromMsg(event.Message)

	if err != nil {
		return fmt.Errorf("error extracting orderID from event message")
	}

	msg := fmt.Sprintf("Received payment event for %v", orderID)

	logger.Info(msg)

	res, err := p.orderService.GetOrder(ctx, &api.GetOrderRequest{
		OrderID: orderID,
	})

	if err != nil {
		msg = fmt.Sprintf("%d orderID not fount", orderID)
		logger.Info(msg)
		return fmt.Errorf("orderID not found")
	}

	customerRes, err := p.customerService.GetCustomer(ctx, &api.GetCustomerRequest{
		CustomerID: res.CustomerID,
	})

	if err != nil {
		msg = fmt.Sprintf("%d customerID not fount", orderID)
		logger.Info(msg)
		return fmt.Errorf("customerID not found")
	}
	msg = fmt.Sprintf("Sending order to\n%v\n%v", customerRes.Name, customerRes.Address)
	logger.Info(msg)

	uuid, err := uuid.NewRandom()

	if err != nil {
		logger.Errorf("error creating uuid: %v", err)
	}

	msg = fmt.Sprintf("%d shipped", orderID)

	logger.Info(msg)

	if err := p.publisher.Publish(context.Background(), &api.Event{
		ID:        uuid.String(),
		Timestamp: time.Now().Unix(),
		Message:   msg,
	}); err != nil {
		logger.Errorf("error while publishing: %v", err)
	}

	return nil
}
