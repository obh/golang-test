package serviceimpl

import (
	"context"
	"errors"
	"simple/src/models"
	"simple/src/service"
)

type DiscountSvcImpl struct {
	ProductSvc  service.ProductService
	CustomerSvc service.CustomerService
	Discounts   []service.Discount
}

func InitDiscountSvcImpl(ps service.ProductService, cs service.CustomerService) DiscountSvcImpl {
	d := []service.Discount{
		&LoyalCustomerDiscount{},
		&SwitchesDiscount{productSvc: ps},
	}

	return DiscountSvcImpl{
		ProductSvc:  ps,
		CustomerSvc: cs,
		Discounts:   d,
	}
}

func (d *DiscountSvcImpl) Find(ctx context.Context, order models.Order) ([]models.Discount, error) {
	customer, err := d.CustomerSvc.Get(ctx, order.CustomerId)
	if err != nil {
		return nil, errors.New("invalid customer")
	}

	discountsApplied := []models.Discount{}
	for _, discount := range d.Discounts {
		if discount.Eligible(ctx, *customer, order) {
			discountsApplied = append(discountsApplied, discount.Apply(ctx, *customer, order))
		}
	}
	return discountsApplied, nil
}

type LoyalCustomerDiscount struct {
}

type SwitchesDiscount struct {
	productSvc service.ProductService
}

func (ls *LoyalCustomerDiscount) Eligible(c context.Context, customer models.Customer, o models.Order) bool {
	return customer.Revenue > 1000
}

func (ls *LoyalCustomerDiscount) Apply(c context.Context, customer models.Customer, o models.Order) models.Discount {
	return models.Discount{
		Id:                 "1",
		DiscountType:       models.DiscountTypeAmount,
		DiscountAmount:     0.1 * o.OrderTotal,
		DiscountPercentage: 10,
		TotalAmount:        o.OrderTotal,
		TotalPayable:       o.OrderTotal - 0.1*o.OrderTotal,
	}
}

func (ls *SwitchesDiscount) Eligible(c context.Context, customer models.Customer, o models.Order) bool {
	totalCount := 0.0
	for _, item := range o.Items {
		product, err := ls.productSvc.Get(c, item.ProductId)
		if err == nil {
			if product.Id == "2" {
				totalCount += item.Quantity
			}
		}
	}
	return totalCount >= 5
}

func (ls *SwitchesDiscount) Apply(c context.Context, customer models.Customer, o models.Order) models.Discount {
	totalCount := 0.0
	for _, item := range o.Items {
		product, err := ls.productSvc.Get(c, item.ProductId)
		if err == nil {
			if product.Id == "2" {
				totalCount += item.Quantity
			}
		}
	}
	return models.Discount{
		Id:           "2",
		DiscountType: models.DiscountTypeProduct,

		TotalAmount:  o.OrderTotal,
		TotalPayable: o.OrderTotal,

		DiscountProduct:  "2",
		DiscountQuantity: totalCount / 5,
	}
}
