package stock

import (
	"context"

	"github.com/ob-vss-20ss/blatt2-cyan/api"
	catalog "github.com/ob-vss-20ss/blatt2-cyan/catalog"
)

type Stock struct {
	items []*catalog.Item
}

func (c *Stock) GetItemsInStock(ctx context.Context, req *api.ItemsInStockRequest, rsp *api.ItemsInStockResponse) error {
	return nil

}
