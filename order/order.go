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
	paymentService  api.PaymentService
}

type Ordering struct {
	customerID  uint32
	articleList []*api.ArticleWithAmount
	paid        bool
	shipped     bool
}

func New(catalogService api.CatalogService, stockService api.StockService,
	customerService api.CustomerService, paymentService api.PaymentService) *Order {
	orderMap := make(map[uint32]Ordering)
	return &Order{orderMap, 0, catalogService, stockService, customerService, paymentService}
}

func (o *Order) Process(ctx context.Context, event *api.Event) error {
	eventMsg := ExtractEventMsg(event.Message)

	msg := fmt.Sprintf("Received %v msg", eventMsg)
	logger.Info(msg)

	orderID, err := shipment.ExtractOrderIDFromMsg(event.Message)

	if err != nil {
		logger.Fatal(err)
	}

	if eventMsg == "paid" {
		//Bestellung finden und auf bezahlt setzen
		tmp := o.orderMap[orderID]
		tmp.paid = true
		o.orderMap[orderID] = tmp
	}

	if eventMsg == "shipped" {
		tmp := o.orderMap[orderID]
		tmp.shipped = true
		//nicht sicher ob dieser Schritt nötig ist
		o.orderMap[orderID] = tmp
	}

	return nil
}

// nolint:lll
func (o *Order) PlaceOrder(ctx context.Context, req *api.PlaceOrderRequest, res *api.PlaceOrderResponse) error {
	msg := fmt.Sprint("Received order request from customer", req.CustomerID)

	logger.Info(msg)

	//Bei CustomerService überprüfen ob Kunde vorhanden
	_, err := o.customerService.GetCustomer(ctx, &api.GetCustomerRequest{
		CustomerID: req.CustomerID,
	})

	if err != nil {
		logger.Info("Kundennummer nicht gefunden. Bestellung abgebrochen")
		message := "Die von Ihnen angegebene Kundennummer ist ungültig.Falls Sie noch kein Konto bei uns haben, registrieren Sie sich bitte zuerst beim Customer-Service\n"
		res.Message = message
		return fmt.Errorf("customer not found")
	}

	//Bei StockService verfügbarkeit prüfen
	if !o.CheckStock(req.ArticleList) {
		logger.Info("Zu wenig Artikel vorhanden. Bestellung abgebrochen")
		res.Message = "Von einem der von Ihnen gewälten Artikel ist nicht mehr genug auf Lager.Reduzieren Sie die Bestellmenge und versuche Sie es nochmal."
		return fmt.Errorf("stock too low")
	}

	//Bei Stock Service Bestand reduzieren
	o.ReduceStock(req.ArticleList)

	//Mit article Service Preis ausrechnen
	price := o.CalculatePrice(req.ArticleList)

	//Bestellung speichern

	ordering := Ordering{req.CustomerID, req.ArticleList, false, false}

	o.orderMap[o.key] = ordering

	//Antwort an Client
	msg = fmt.Sprintf("Bestellung Eingegangen. Sie wird zu Ihnen geliefert sobald sie %d € an unseren Zahlungsdienstleister überwiesen haben", price)
	res.OrderID = o.key
	res.Message = msg

	//Key für die nächste Bestellung erhöhen
	o.key++

	logger.Info("Bestellung erfolgreich abgeschlossen.")

	return nil
}

// nolint:lll
func (o *Order) ReturnItem(ctx context.Context, req *api.ReturnRequest, res *api.ReturnResponse) error {
	msg := fmt.Sprint("Received return request from customer", req.CustomerID)

	logger.Info(msg)

	ordering, ok := o.orderMap[req.OrderID]

	if !ok {
		logger.Info("Unbekannte Bestellnummer. Rückgabe abgebrochen.")
		res.Message = "Die von Ihnen angegebene Bestellnummer ist uns nicht bekannt."
		return fmt.Errorf("order not found")
	}

	//Bestellung prüfen (vorhanden, Kundennummer stimmt überein, artikel mit entsprechender Stückzahl enthalten)
	if ordering.customerID != req.CustomerID {
		logger.Info("Angegebene Kundennummer stimmt nicht. Rückgabe abgebrochen.")
		res.Message = "Die von Ihnen angegebene Kundennummer stimmt nicht mit der Kundennummer der Bestellung überein."
		return fmt.Errorf("wrong customer")
	}

	if !o.OrderContainsArticle(ordering.articleList, req.ArticleList) {
		logger.Info("Rückgabeartikel wa nicht in der Bestellung enthalten. Rückgabe abgebrochen.")
		res.Message = "Die von Ihnen mitgegebene Retourliste enthält mindestens einen Artikel der nicht in der Bestellung enthalten war."
		return fmt.Errorf("order didn't contain article")
	}

	//Preis ausrechnen (catalog Service) und die Summe in der Antwort an den Client schicken
	if !req.Replacement {
		logger.Info("Rückgabe erfolgreich abgeschlossen")
		res.Message = fmt.Sprintf("Hiermit erstatten wir wie gewünscht den Kaufpreis: %d€", o.CalculatePrice(req.ArticleList))
		return nil
	}

	var replacementSuccess bool
	if req.Replacement {
		replacementSuccess = o.CreateReplacement(req)
	}

	if !replacementSuccess && req.Replacement {
		logger.Info("Kaufpreis wurde erstattet, da der Artikel nicht mehr auf Lager ist.")
		res.Message = fmt.Sprintf("Leider konnten wir die Ware nicht ersetzen. Deshalb erstatten wir hiermit den Kaufpreis: %d€", o.CalculatePrice(req.ArticleList))
	} else {
		logger.Info("Rückgabe erfolgreich abgeschlossen")
		res.Message = fmt.Sprintf("Der angeforderte Ersatz ist mit der Bestellnummer %d auf dem Weg zu Ihnen.", o.key)
		o.key++
	}

	//Artikel aus Bestellung bzw. Bestellung löschen
	tmp := o.orderMap[req.OrderID]
	tmp.articleList = o.ShortenOrder(ordering.articleList, req.ArticleList)
	if len(tmp.articleList) == 0 {
		delete(o.orderMap, req.OrderID)
	} else {
		o.orderMap[req.OrderID] = tmp
	}

	return nil
}

// nolint:lll
func (o *Order) CancelOrder(ctx context.Context, req *api.CancelRequest, res *api.CancelResponse) error {
	msg := fmt.Sprintf("Received cancel request from %d (customerID)", req.CustomerID)

	logger.Info(msg)

	ordering, ok := o.orderMap[req.OrderID]

	if !ok {
		logger.Info("Bestellnummer nicht bekannt. Stornierung abgebrochen.")
		res.Message = "Die von Ihnen angegebene Bestellnummer ist uns nicht bekannt."
		return fmt.Errorf("order not found")
	}

	//Bestellung überprüfen(nur Kundennummer)
	if ordering.customerID != req.CustomerID {
		logger.Info("Kundennummer stimmt nicht. Stornierung abgebrochen")
		res.Message = "Die von Ihnen angegebene Kundennummer stimmt nicht mit der Kundennummer, der von Ihnen angegebenen Bestellung überein"
		return fmt.Errorf("wrong customerID")
	}

	//Prüfen ob Bestellung Versandt
	if ordering.shipped {
		logger.Info("Bestellung bereits verschickt. Stornierung abgebrochen.")
		res.Message = "Die Stornierung dieser Bestellung ist leider nicht mehr möglich, da sie bereits versandt wurde"
		return fmt.Errorf("order already shipped")
	}

	//Preis ausrechnen (catalog Service ansprechen), an Antwort an Client senden
	price := o.CalculatePrice(ordering.articleList)

	//Bestände im StockService erhöhen
	o.IncreaseStock(ordering.articleList)

	//Bestellung löschen
	delete(o.orderMap, req.OrderID)

	logger.Info("Bestellung erfolgreich storniert.")
	//Prüfen ob bezahlt
	if !ordering.paid {
		res.Message = "Ihre Bestellung konnte storniert werden."
		return nil
	}

	res.Message = fmt.Sprintf("Ihre Bestellung konnte storniert werden. Hiermit erstatten wir Ihnen den Kaufpreis: %d€", price)
	return nil
}

func (o *Order) GetOrder(ctx context.Context, req *api.GetOrderRequest, res *api.GetOrderResponse) error {
	msg := fmt.Sprint("Received get order request:", req.OrderID)

	logger.Info(msg)

	ordering, ok := o.orderMap[req.OrderID]

	if !ok {
		logger.Info("Bestellung nicht gefunden. GetOrder abgebrochen")
		return fmt.Errorf("order not found")
	}

	res.CustomerID = ordering.customerID
	res.ArticleList = ordering.articleList
	res.Paid = ordering.paid
	res.Shipped = ordering.shipped

	logger.Info("GetOrder erfolgreich abgeschlossen.")

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
			ArticleID: articleList[i].ArticleID,
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
			ArticleID: articleList[i].ArticleID,
		})

		if err != nil {
			return false
		}

		if stockRsp.Amount < articleList[i].Amount {
			return false
		}
	}

	return true
}

func (o *Order) ReduceStock(articleList []*api.ArticleWithAmount) {
	for i := range articleList {
		_, err := o.stockService.ReduceStockOfItem(context.Background(), &api.ReduceStockRequest{
			ArticleID: articleList[i].ArticleID,
			Amount:    articleList[i].Amount,
		})

		if err != nil {
			panic(err)
		}
	}
}

func (o *Order) IncreaseStock(articleList []*api.ArticleWithAmount) {
	for i := range articleList {
		_, err := o.stockService.IncreaseStockOfItem(context.Background(), &api.IncreaseStockRequest{
			ArticleID: articleList[i].ArticleID,
			Amount:    articleList[i].Amount,
		})

		if err != nil {
			panic(err)
		}
	}
}

func (o *Order) OrderContainsArticle(articleListOrder []*api.ArticleWithAmount,
	articleListReturned []*api.ArticleWithAmount) bool {
	for i := range articleListReturned {
		articleFound := false
		for j := range articleListReturned {
			if articleListReturned[i].ArticleID == articleListOrder[j].ArticleID &&
				articleListReturned[i].Amount <= articleListOrder[j].Amount {
				articleFound = true
				break
			}
		}

		if !articleFound {
			return false
		}
	}

	return true
}

func (o *Order) ShortenOrder(articleListOrder []*api.ArticleWithAmount,
	articleListReturned []*api.ArticleWithAmount) []*api.ArticleWithAmount {
	for i := range articleListReturned {
		for j := range articleListReturned {
			if articleListReturned[i].ArticleID == articleListOrder[j].ArticleID {
				if articleListReturned[i].Amount == articleListOrder[j].Amount {
					articleListOrder[j] = articleListOrder[len(articleListOrder)-1]
					articleListOrder = articleListOrder[:len(articleListOrder)-1]
				} else {
					articleListOrder[j].Amount -= articleListReturned[i].Amount
				}
			}
		}
	}

	return articleListOrder
}

func (o *Order) CreateReplacement(req *api.ReturnRequest) bool {
	if o.CheckStock(req.ArticleList) {
		o.ReduceStock(req.ArticleList)
		o.saveOrder(req.CustomerID, req.ArticleList)
		_, err := o.paymentService.ReceivePayment(context.Background(), &api.PaymentRequest{
			OrderID: o.key,
		})

		return err != nil
	}

	return false
}

func (o *Order) saveOrder(customerID uint32, articleList []*api.ArticleWithAmount) {
	ordering := Ordering{customerID, articleList, false, false}
	o.orderMap[o.key] = ordering
}
