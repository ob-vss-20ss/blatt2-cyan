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
	publisher micro.Event
}

func New(publisher micro.Event) *Shipment {
	return &Shipment{publisher: publisher}
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

	msg := fmt.Sprintf("Received payment event for %v", event.Message)

	logger.Info(msg)

	uuid, err := uuid.NewRandom()

	if err != nil {
		logger.Errorf("error creating uuid: %v", err)
	}

	msg = fmt.Sprintf("%d shipped", orderID)

	if err := p.publisher.Publish(context.Background(), &api.Event{
		Id:        uuid.String(),
		Timestamp: time.Now().Unix(),
		Message:   msg,
	}); err != nil {
		logger.Errorf("error while publishing: %v", err)
	}

	return nil
}