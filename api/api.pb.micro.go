// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api/api.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Catalog service

type CatalogService interface {
	GetItemsInStock(ctx context.Context, in *ItemsInStockRequest, opts ...client.CallOption) (*ItemsInStockResponse, error)
}

type catalogService struct {
	c    client.Client
	name string
}

func NewCatalogService(name string, c client.Client) CatalogService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "api"
	}
	return &catalogService{
		c:    c,
		name: name,
	}
}

func (c *catalogService) GetItemsInStock(ctx context.Context, in *ItemsInStockRequest, opts ...client.CallOption) (*ItemsInStockResponse, error) {
	req := c.c.NewRequest(c.name, "Catalog.GetItemsInStock", in)
	out := new(ItemsInStockResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Catalog service

type CatalogHandler interface {
	GetItemsInStock(context.Context, *ItemsInStockRequest, *ItemsInStockResponse) error
}

func RegisterCatalogHandler(s server.Server, hdlr CatalogHandler, opts ...server.HandlerOption) error {
	type catalog interface {
		GetItemsInStock(ctx context.Context, in *ItemsInStockRequest, out *ItemsInStockResponse) error
	}
	type Catalog struct {
		catalog
	}
	h := &catalogHandler{hdlr}
	return s.Handle(s.NewHandler(&Catalog{h}, opts...))
}

type catalogHandler struct {
	CatalogHandler
}

func (h *catalogHandler) GetItemsInStock(ctx context.Context, in *ItemsInStockRequest, out *ItemsInStockResponse) error {
	return h.CatalogHandler.GetItemsInStock(ctx, in, out)
}

// Client API for Customer service

type CustomerService interface {
}

type customerService struct {
	c    client.Client
	name string
}

func NewCustomerService(name string, c client.Client) CustomerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "api"
	}
	return &customerService{
		c:    c,
		name: name,
	}
}

// Server API for Customer service

type CustomerHandler interface {
}

func RegisterCustomerHandler(s server.Server, hdlr CustomerHandler, opts ...server.HandlerOption) error {
	type customer interface {
	}
	type Customer struct {
		customer
	}
	h := &customerHandler{hdlr}
	return s.Handle(s.NewHandler(&Customer{h}, opts...))
}

type customerHandler struct {
	CustomerHandler
}

// Client API for Order service

type OrderService interface {
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "api"
	}
	return &orderService{
		c:    c,
		name: name,
	}
}

// Server API for Order service

type OrderHandler interface {
}

func RegisterOrderHandler(s server.Server, hdlr OrderHandler, opts ...server.HandlerOption) error {
	type order interface {
	}
	type Order struct {
		order
	}
	h := &orderHandler{hdlr}
	return s.Handle(s.NewHandler(&Order{h}, opts...))
}

type orderHandler struct {
	OrderHandler
}

// Client API for Payment service

type PaymentService interface {
}

type paymentService struct {
	c    client.Client
	name string
}

func NewPaymentService(name string, c client.Client) PaymentService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "api"
	}
	return &paymentService{
		c:    c,
		name: name,
	}
}

// Server API for Payment service

type PaymentHandler interface {
}

func RegisterPaymentHandler(s server.Server, hdlr PaymentHandler, opts ...server.HandlerOption) error {
	type payment interface {
	}
	type Payment struct {
		payment
	}
	h := &paymentHandler{hdlr}
	return s.Handle(s.NewHandler(&Payment{h}, opts...))
}

type paymentHandler struct {
	PaymentHandler
}

// Client API for Shipment service

type ShipmentService interface {
}

type shipmentService struct {
	c    client.Client
	name string
}

func NewShipmentService(name string, c client.Client) ShipmentService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "api"
	}
	return &shipmentService{
		c:    c,
		name: name,
	}
}

// Server API for Shipment service

type ShipmentHandler interface {
}

func RegisterShipmentHandler(s server.Server, hdlr ShipmentHandler, opts ...server.HandlerOption) error {
	type shipment interface {
	}
	type Shipment struct {
		shipment
	}
	h := &shipmentHandler{hdlr}
	return s.Handle(s.NewHandler(&Shipment{h}, opts...))
}

type shipmentHandler struct {
	ShipmentHandler
}

// Client API for Stock service

type StockService interface {
}

type stockService struct {
	c    client.Client
	name string
}

func NewStockService(name string, c client.Client) StockService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "api"
	}
	return &stockService{
		c:    c,
		name: name,
	}
}

// Server API for Stock service

type StockHandler interface {
}

func RegisterStockHandler(s server.Server, hdlr StockHandler, opts ...server.HandlerOption) error {
	type stock interface {
	}
	type Stock struct {
		stock
	}
	h := &stockHandler{hdlr}
	return s.Handle(s.NewHandler(&Stock{h}, opts...))
}

type stockHandler struct {
	StockHandler
}
