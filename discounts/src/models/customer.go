package models

type Customer struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Since   string  `json:"since"`
	Revenue float64 `json:"revenue"`
}
