package order

import (
	"context"
	"fmt"
	"strings"

	"github.com/micro/go-micro/v2/logger"

	"github.com/ob-vss-20ss/blatt2-cyan/api"
	"github.com/ob-vss-20ss/blatt2-cyan/shipment"
)

type Order struct {
	orderMap        map[uint32]Ordering
	key             uint32
	catalogService  api.CatalogService
	stockService    api.StockService
	customerService api.CustomerService
}

type Ordering struct {
	customerID  uint32
	articleList []*api.ArticleWithAmount
	payed       bool
	shipped     bool
}

func New(catalogService api.CatalogService, stockService api.StockService, customerService api.CustomerService) *Order {
	orderMap := make(map[uint32]Ordering)
	return &Order{orderMap, 0, catalogService, stockService, customerService}
}

func (o *Order) Process(ctx context.Context, event *api.Event) error {
	eventMsg := ExtractEventMsg(event.Message)

	msg := fmt.Sprintf("Received return %v msg", eventMsg)
	logger.Info(msg)

	orderID, err := shipment.ExtractOrderIDFromMsg(event.Message)

	if err != nil {
		logger.Fatal(err)
	}

	if eventMsg == "payed" {
		//Bestellung finden und auf bezahlt setzen
		fmt.Printf("%d", orderID)
	}

	if eventMsg == "shipped" {

	}

	//abfragen der adresse bei order service -> customerservice

	//msg := fmt.Sprintf("Received payment event for", event.Message)

	//logger.Info(msg)

	//Bestellung finden und auf bezahlt setzen

	return nil
}

func (o *Order) PlaceOrder(ctx context.Context, req *api.PlaceOrderRequest, res *api.PlaceOrderResponse) error {
	msg := fmt.Sprintf("Received Order from %v (customerID)", req.CustomerID)

	logger.Info(msg)

	//Bei CustomerService überprüfen ob Kunde vorhanden
	_, err := o.customerService.GetCustomer(ctx, &api.GetCustomerRequest{
		CustomerID: req.CustomerID,
	})

	if err != nil {
		res.Message = "Die von Ihnen angegebene Kundennummer ist ungültig.\nFalls Sie noch kein Konto bei uns haben, registrieren Sie sich bitte zuerst beim Customer-Service\n"
		return fmt.Errorf("Customer not found.")
	}

	//Bei StockService verfügbarkeit prüfen
	if !o.CheckStock(req.ArticleList) {
		res.Message = "Von einem der von Ihnen gewälten Artikel ist nicht mehr genug auf Lager.\nReduzieren Sie die Bestellmenge und versuche Sie es nochmal.\n"
		return fmt.Errorf("Stock to low")
	}

	//Bei Stock Service Bestand reduzieren
	o.ReduceStock(req.ArticleList)

	//Mit article Service Preis ausrechnen
	price := o.CalculatePrice(req.ArticleList)

	//Bestellung speichern

	ordering := Ordering{req.CustomerID, req.ArticleList, false, false}

	o.orderMap[o.key] = ordering

	//Antwort an Client
	msg = fmt.Sprintf("Bestellung Eingegangen. Sie wird zu Ihnen geliefert sobald sie %v an unseren Zahlungsdienstleister überwiesen haben", price)
	res.OrderID = o.key
	res.Message = msg

	//Key für die nächste Bestellung erhöhen
	o.key++

	return nil
}

func (o *Order) ReturnItem(ctx context.Context, req *api.ReturnRequest, res *api.ReturnResponse) error {
	msg := fmt.Sprintf("Received return Request from %d (customerID)", req.CustomerID)

	logger.Info(msg)

	//Bestellung prüfen (vorhanden, Kundennummer stimmt überein, artikel mit entsprechender Stückzahl enthalten)

	//Preis ausrechnen (catalog Service) und die Summe in der Antwort an den Client schicken

	return nil
}

func (o *Order) CancelOrder(ctx context.Context, req *api.CancelRequest, res *api.CancelResponse) error {
	msg := fmt.Sprintf("Received return Request from %d (customerID)", req.CustomerID)

	logger.Info(msg)

	//Bestellung überprüfen(nur Kundennummer)
	if o.orderMap[req.OrderID].customerID != req.CustomerID {
		res.Message = "Die von Ihnen angegebene Kundennummer stimmt nicht mit der Kundennummer, der von Ihnen angegebenen Bestellung überein"
		return fmt.Errorf("Wrong CutomerID")
	}
	//Prüfen ob Bestellung Versandt
	if o.orderMap[req.OrderID].shipped {
		res.Message = "Die Stornierung dieser Bestellung ist leider nicht mehr möglich, da sie bereits versandt wurde"
		return fmt.Errorf("Order shipped")
	}

	//Preis ausrechnen (catalog Service ansprechen), an Client in Antwort senden
	o.CalculatePrice(o.orderMap[req.OrderID].articleList)

	//Bestände im StockService erhöhen
	o.IncreaseStock(o.orderMap[req.OrderID].articleList)

	//Bestellung löschen
	delete(o.orderMap, req.OrderID)

	return nil
}

func ExtractEventMsg(msg string) string {
	tmp := strings.Split(msg, " ")
	return tmp[1]
}

func (o *Order) CalculatePrice(articleList []*api.ArticleWithAmount) uint32 {
	price := uint32(0)
	for i := range articleList {

		catalogRsp, err := o.catalogService.GetItem(context.Background(), &api.ItemRequest{
			ItemID: articleList[i].ArticleID,
		})

		if err != nil {
			panic(err)
		}

		price += catalogRsp.Price * articleList[i].Amount
	}

	return price
}

func (o *Order) CheckStock(articleList []*api.ArticleWithAmount) bool {
	for i := range articleList {

		stockRsp, err := o.stockService.GetStockOfItem(context.Background(), &api.StockOfItemRequest{
			ItemID: articleList[i].ArticleID,
		})

		if err != nil {
			panic(err)
		}

		if stockRsp.Stock < articleList[i].Amount {
			return false
		}
	}

	return true
}

func (o *Order) ReduceStock(articleList []*api.ArticleWithAmount) {
	for i := range articleList {
		_, err := o.stockService.ReduceStockOfItem(context.Background(), &api.ReduceStockRequest{
			ItemID: articleList[i].ArticleID,
			Amount: articleList[i].Amount,
		})

		if err != nil {
			panic(err)
		}
	}
}

func (o *Order) IncreaseStock(articleList []*api.ArticleWithAmount) {
	for i := range articleList {
		_, err := o.stockService.IncreaeStockOfItem(context.Background(), &api.IncreaseStockRequest{
			ItemID: articleList[i].ArticleID,
			Amount: articleList[i].Amount,
		})

		if err != nil {
			panic(err)
		}
	}
}
