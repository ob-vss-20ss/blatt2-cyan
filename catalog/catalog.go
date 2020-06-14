package catalog

import (
	"github.com/ob-vss-20ss/blatt2-cyan/api"
)

type Catalog struct {
	items []*item
	stock api.StockService
}

func New(stock api.StockService) *Catalog {
	return &Catalog{
		stock: stock,
	}
}

/*func (c *Catalog) GetItemsInStock(ctx context.Context, req *api.ItemsInStockRequest, rsp *api.ItemsInStockResponse) error {
	stockRsp, err := c.stock.GetItemsInStock(context.Background(), &api.GetItemsInStockRequest{})

	if err != nil {
		logger.Error(err)
	}

}*/
