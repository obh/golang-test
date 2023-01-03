package delivery

import (
	"fmt"
	"net/http"
	"simple/src/models"
	"simple/src/service"

	"github.com/labstack/echo/v4"
)

const (
	offers = "/discounts"
)

type DiscountsHandler struct {
	customerSvc service.CustomerService
	productSvc  service.ProductService
	discountSvc service.DiscountService
}

func ConfigureDiscountsHandler(e *echo.Echo, cs service.CustomerService, ps service.ProductService, ds service.DiscountService) {
	o := &DiscountsHandler{
		customerSvc: cs,
		productSvc:  ps,
		discountSvc: ds,
	}
	e.Add("POST", offers, o.getDiscounts)
}

func (d *DiscountsHandler) getDiscounts(c echo.Context) error {
	ctx := c.Request().Context()
	order := &models.Order{}
	err := c.Bind(&order)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	discounts, err := d.discountSvc.Find(ctx, *order)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(200, discounts)
}
