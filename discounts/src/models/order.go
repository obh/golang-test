package models

/*
{
  "id": "1",
  "customer-id": "1",
  "items": [
    {
      "product-id": "B102",
      "quantity": "10",
      "unit-price": "4.99",
      "total": "49.90"
    }
  ],
  "total": "49.90"
}
*/
type Order struct {
	Id         string  `json:"id"`
	CustomerId string  `json:"customer-id"`
	Items      []Item  `json:"items"`
	OrderTotal float64 `json:"total"`
}

type Item struct {
	ProductId string  `json:"product-id"`
	Quantity  float64 `json:"quantity"`
	UnitPrice float64 `json:"unit-price"`
	ItemTotal float64 `json:"total"`
}
