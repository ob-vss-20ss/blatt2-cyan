package order

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/logger"
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Order struct {
}

func New() *Order {
	return &Order{}
}

func (o *Order) PlaceOrder(ctx context.Context, req *api.PlaceOrderRequest, res *api.PlaceOrderResponse) error {
	msg := fmt.Sprintf("Received Order from %v (customerID)", req.CustomerID)

	logger.Info(msg)

	//Bei CustomerService überprüfen ob Kunce vorhanden

	//Bei StockService verfügbarkeit prüfen

	//Bei Stock Service Besand reduzieren

	//Antwort an Client

	return nil
}

func (o *Order) ReturnItem(ctx context.Context, req *api.ReturnRequest, res *api.ReturnResponse) error {
	msg := fmt.Sprintf("Received return Request from (customerID)", req.CustomerID)

	logger.Info(msg)

	//Bestellung prüfen (vorhanden, Kundennummer stimmt überein, artikel mit entsprechender Stückzahl enthalten)

	//Preis ausrechnen (catalo Service) und die Summe in der Antwort an den Client schicken

	return nil
}

func (o *Order) CancelOrder(ctx context.Context, req *api.CancelRequest, res *api.CancelResponse) error {
	msg := fmt.Sprintf("Received return Request from (customerID)", req.CustomerID)

	logger.Info(msg)

	//Bestellung überprüfen(nur Kundennummer)

	//Prüfen ob Bestellung Versandt

	//Bestände im StockService erhöhen

	//Bestellung löschen

	//Preis ausrechnen (catalog Service ansprechen), an Client in Antwort senden

	return nil
}

func (o *Order) Process(ctx context.Context, event *api.Event) error {
	//abfragen der adresse bei order service -> customerservice

	//msg := fmt.Sprintf("Received payment event for", event.Message)

	//logger.Info(msg)

	//Bestellung finden und auf bezahlt setzen

	return nil
}
