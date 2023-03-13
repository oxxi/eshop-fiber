package models

type ProductModel struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    uint64  `json:"quantity"`
	Price       float64 `json:"price"`
}
