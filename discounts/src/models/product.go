package models

type Product struct {
	Id          string  `json:"string"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
}
