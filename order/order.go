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
	payed       bool
	shipped     bool
}

func New(catalogService api.CatalogService, stockService api.StockService, customerService api.CustomerService) *Order {
	orderMap := make(map[uint32]Ordering)
	return &Order{orderMap, 0, catalogService, stockService, customerService}
}

func (o *Order) Process(ctx context.Context, event *api.Event) error {
	eventMsg := ExtractEventMsg(event.Message)

	msg := fmt.Sprintf("Received %v msg", eventMsg)
	logger.Info(msg)

	orderID, err := shipment.ExtractOrderIDFromMsg(event.Message)

	if err != nil {
		logger.Fatal(err)
	}

	if eventMsg == "payed" {
		//Bestellung finden und auf bezahlt setzen
		tmp := o.orderMap[orderID]
		tmp.payed = true
		//nicht sicher ob dieser Schritt nötig ist
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

func (o *Order) PlaceOrder(ctx context.Context, req *api.PlaceOrderRequest, res *api.PlaceOrderResponse) error {
	msg := fmt.Sprintf("Received order request from %v (customerID)", req.CustomerID)

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
	msg := fmt.Sprintf("Received return request from %d (customerID)", req.CustomerID)

	logger.Info(msg)

	ordering, ok := o.orderMap[req.OrderID]

	if !ok {
		res.Message = "Die von Ihnen angegebene Bestellnummer ist uns nicht bekannt."
		return fmt.Errorf("Order not Found")
	}

	//Bestellung prüfen (vorhanden, Kundennummer stimmt überein, artikel mit entsprechender Stückzahl enthalten)
	if ordering.customerID != req.CustomerID {
		res.Message = "Die von Ihnen angegebene Kundennummer stimmt nicht mit der Kundennummer der Bestellung überein."
		return fmt.Errorf("Wrong Customer")
	}

	if !o.OrderContainsArticle(ordering.articleList, req.ArticleList) {
		res.Message = "Die von Ihnen mitgegebene Retourliste enthält mindestens einen Artikel der nicht in der Bestellung enthalten war."
		return fmt.Errorf("Order didn't contain article")
	}

	//Preis ausrechnen (catalog Service) und die Summe in der Antwort an den Client schicken
	var replacementSuccess bool
	if req.Replacement {
		replacementSuccess = o.CreateReplacement(*req)
	}

	if !replacementSuccess && req.Replacement {
		res.Message = fmt.Sprint("Leider konnten wir die Ware nicht ersetzen. Deshalb erstatten wir hiermit den Kaufpreis:\n", o.CalculatePrice(req.ArticleList))
	} else {
		res.Message = fmt.Sprint("Der angeforderte Ersatz ist auf dem mit der Bestellnummer %d auf dem Weg", o.key)
		o.key++
	}

	if !req.Replacement {
		res.Message = fmt.Sprint("Hier mit erstatten wir wie gewünscht den Kaufpreis:", o.CalculatePrice(req.ArticleList))
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

func (o *Order) CancelOrder(ctx context.Context, req *api.CancelRequest, res *api.CancelResponse) error {
	msg := fmt.Sprintf("Received cancel request from %d (customerID)", req.CustomerID)

	logger.Info(msg)

	ordering, ok := o.orderMap[req.OrderID]

	if !ok {
		res.Message = "Die von Ihnen angegebene Bestellnummer ist uns nicht bekannt."
		return fmt.Errorf("Order not Found")
	}

	//Bestellung überprüfen(nur Kundennummer)
	if ordering.customerID != req.CustomerID {
		res.Message = "Die von Ihnen angegebene Kundennummer stimmt nicht mit der Kundennummer, der von Ihnen angegebenen Bestellung überein"
		return fmt.Errorf("Wrong CutomerID")
	}

	//Prüfen ob Bestellung Versandt
	if ordering.shipped {
		res.Message = "Die Stornierung dieser Bestellung ist leider nicht mehr möglich, da sie bereits versandt wurde"
		return fmt.Errorf("Order shipped")
	}

	//Preis ausrechnen (catalog Service ansprechen), an Client in Antwort senden
	o.CalculatePrice(ordering.articleList)

	//Bestände im StockService erhöhen
	o.IncreaseStock(ordering.articleList)

	//Bestellung löschen
	delete(o.orderMap, req.OrderID)

	return nil
}

func (o *Order) GetOrder(ctx context.Context, req *api.GetOrderRequest, res *api.GetOrderResponse) error {
	msg := fmt.Sprint("Received get order request:", req.OrderID)

	logger.Info(msg)

	ordering, ok := o.orderMap[req.OrderID]

	if !ok {
		return fmt.Errorf("Order not Found")
	}

	res.CustomerID = ordering.customerID
	res.ArticleList = ordering.articleList

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

func (o *Order) OrderContainsArticle(articleListOrder []*api.ArticleWithAmount, articleListReturned []*api.ArticleWithAmount) bool {
	for i := range articleListReturned {
		articleFound := false
		for j := range articleListReturned {
			if articleListReturned[i].ArticleID == articleListOrder[j].ArticleID && articleListReturned[i].Amount <= articleListOrder[j].Amount {
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

func (o *Order) ShortenOrder(articleListOrder []*api.ArticleWithAmount, articleListReturned []*api.ArticleWithAmount) []*api.ArticleWithAmount {
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

func (o *Order) CreateReplacement(req api.ReturnRequest) bool {
	if o.CheckStock(req.ArticleList) {
		o.ReduceStock(req.ArticleList)
		o.saveOrder(req.CustomerID, req.ArticleList)
		_, err := o.paymentService.ReceivePayment(context.Background(), &api.PaymentRequest{
			OrderID: o.key,
		})

		if err != nil {
			return false
		}
		return true
	}

	return false
}

func (o *Order) saveOrder(customerID uint32, articleList []*api.ArticleWithAmount) {
	ordering := Ordering{customerID, articleList, false, false}
	o.orderMap[o.key] = ordering
}
