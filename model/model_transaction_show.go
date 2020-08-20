package model

type TransactionShow struct {
	ProductID          uint   `json:"product_id"`
	ProductName        string `json:"product_name"`
	ProductPrince      int    `json:"product_prince"`
	StoreName          string `json:"store_name"`
	Note               string `json:"note"`
	Quantity           int    `json:"quantity"`
	UserAddressID      uint   `json:"address_id"`
	Origin             string `json:"origin"`
	Destination        string `json:"destination"`
	CourierName        string `json:"courier_name"`
	CourierServiceID   int    `json:"courier_service_id"`
	CourierServiceName string `json:"courier_service_name"`
	CourierCost        int    `josn:"courier_cost"`
	PaymentID          uint   `json:"payment_id"`
	PaymentType        string `json:"payment_type"`
	PaymentAccount     string `json:"payment_account"`
	PaymentName        string `json:"payment_name"`
	PaymentNumber      string `json:"payment_number"`
	UnixCode           int    `json:"unix_code"`
	TotalPrince        int    `json:"total_prince"`
}
