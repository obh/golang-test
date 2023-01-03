package main

import (
	"fmt"
	"simple/src/delivery"
	"simple/src/service/serviceimpl"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("Starting server!")

	e := echo.New()
	e.Use(middleware.Logger())

	customerSvc := serviceimpl.InitCustomerSvcImpl()
	productSvc := serviceimpl.InitProductSvcImpl()
	discountSvc := serviceimpl.InitDiscountSvcImpl(&productSvc, &customerSvc)

	delivery.ConfigureDiscountsHandler(e, &customerSvc, &productSvc, &discountSvc)

	e.Start(":8282")
}
