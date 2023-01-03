package serviceimpl

import (
	"context"
	"errors"
	"simple/src/models"
)

type CustomerSvcImpl struct {
	customers []models.Customer
}

func (cs *CustomerSvcImpl) Get(ctx context.Context, id string) (*models.Customer, error) {
	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].Id == id {
			return &cs.customers[i], nil
		}
	}
	return nil, errors.New("customer not found")
}

func (cs *CustomerSvcImpl) GetList(ctx context.Context, id []string) {

}

func InitCustomerSvcImpl() CustomerSvcImpl {
	c1 := models.Customer{
		Id:      "1",
		Name:    "Coca Cola",
		Since:   "2014-06-28",
		Revenue: 492.12,
	}
	c2 := models.Customer{
		Id:      "2",
		Name:    "Teamleader",
		Since:   "2015-01-15",
		Revenue: 1505.95,
	}
	c3 := models.Customer{
		Id:      "3",
		Name:    "Jeroen De Wit",
		Since:   "2016-02-11",
		Revenue: 0.0,
	}
	customers := []models.Customer{c1, c2, c3}
	return CustomerSvcImpl{customers: customers}
}
