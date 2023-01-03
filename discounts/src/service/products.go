package service

import (
	"context"
	"simple/src/models"
)

type ProductService interface {
	Get(context.Context, string) (*models.Product, error)
	// GetList(ctx context.Context, ids []string)
}
