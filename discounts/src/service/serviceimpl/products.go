package serviceimpl

import (
	"context"
	"encoding/json"
	"errors"
	"simple/src/models"
)

const productsJson = `
[
  {
    "id": "A101",
    "description": "Screwdriver",
    "category": "1",
    "price": "9.75"
  },
  {
    "id": "A102",
    "description": "Electric screwdriver",
    "category": "1",
    "price": "49.50"
  },
  {
    "id": "B101",
    "description": "Basic on-off switch",
    "category": "2",
    "price": "4.99"
  },
  {
    "id": "B102",
    "description": "Press button",
    "category": "2",
    "price": "4.99"
  },
  {
    "id": "B103",
    "description": "Switch with motion detector",
    "category": "2",
    "price": "12.95"
  }
]`

type ProductServiceImpl struct {
	products []models.Product
}

func (p *ProductServiceImpl) Get(ctx context.Context, id string) (*models.Product, error) {
	for _, product := range p.products {
		if product.Id == id {
			return &product, nil
		}
	}
	return nil, errors.New("cannot find product")
}

// func (p *ProductServiceImpl) GetList(ctx context.Context, ids []string) {

// }

func InitProductSvcImpl() ProductServiceImpl {
	products := []models.Product{}
	json.Unmarshal([]byte(productsJson), &products)
	return ProductServiceImpl{products: products}
}
