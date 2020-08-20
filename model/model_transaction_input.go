package model

type TransactionInput struct {
	PaymentID        uint   `json:"payment_id"`
	AddressID        uint   `json:"address_id"`
	CourierName      string `json:"courier_name"`
	Note             string `json:"note"`
	CourierServiceID uint   `json:"courier_service_id"`
}

type BasketTrx struct {
	BasketID uint `json:"basket_id"`
}
