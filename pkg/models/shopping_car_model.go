package models

type ShoppingCarModel struct {
	Quantity  uint `json:"quantity" validate:"required"`
	ProductId uint `json:"productId"`
}

type CheckOut struct {
	UUID string             `json:"uuid" validate:"required"`
	Car  []ShoppingCarModel `json:"car" validate:"required"`
}
