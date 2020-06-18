package customer

import (
	"context"

	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Customer struct {
}

func New() *Customer {
	return &Customer{}
}

func (c *Customer) RegisterCustomer(ctx context.Context, req *api.RegisterCustomerRequest, rsp *api.RegisterCustomerResponse) error {

	return nil
}

func (c *Customer) GetCustomer(ctx context.Context, req *api.GetCustomerRequest, rsp *api.GetCustomerResponse) error {

	return nil
}

func (c *Customer) DeleteCustomer(ctx context.Context, req *api.DeleteCustomerRequest, rsp *api.DeleteCustomerResponse) error {

	return nil
}
