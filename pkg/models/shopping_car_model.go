package models

type ShoppingCarModel struct {
	Quantity  uint `json:"quantity"`
	ProductId uint `json:"productId"`
}

type CheckOut struct {
	UUID string             `json:"uuid"`
	Car  []ShoppingCarModel `json:"car"`
}
