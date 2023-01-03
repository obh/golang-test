package models

const DiscountTypeProduct = "PRODUCT"
const DiscountTypeAmount = "AMOUNT"

type Discount struct {
	Id                 string  `json:"id"`
	Description        string  `json:"description"`
	DiscountType       string  `json:"discountType"`
	DiscountAmount     float64 `json:"discountAmount"`
	DiscountPercentage float64 `json:"discountPercentage"`
	TotalAmount        float64 `json:"totalAmount"`
	TotalPayable       float64 `json:"totalPayable"`
	DiscountProduct    string  `json:"discountProductId"`
	DiscountQuantity   float64 `json:"discountProductQty"`
}
