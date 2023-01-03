package service

import (
	"context"
	"simple/src/models"
)

type CustomerService interface {
	Get(ctx context.Context, id string) (*models.Customer, error)
	// GetList(ctx context.Context, id []string)
}
