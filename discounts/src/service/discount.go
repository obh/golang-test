package service

import (
	"context"
	"simple/src/models"
)

type DiscountService interface {
	Find(context.Context, models.Order) ([]models.Discount, error)
}

type Discount interface {
	Eligible(context.Context, models.Customer, models.Order) bool
	Apply(context.Context, models.Customer, models.Order) models.Discount
}
