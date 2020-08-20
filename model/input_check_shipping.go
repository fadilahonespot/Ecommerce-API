package model

type CheckShippingInput struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
	CityID    int `json:"city_id"`
}
