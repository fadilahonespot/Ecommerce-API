package model

type BasketList struct {
	BasketID     uint            `json:"basket_id"`
	MitraID      uint            `json:"mitra_id"`
	MitraName    string          `json:"mitra_name"`
	ProductsList []BasketSubList `json:"products_list"`
	Total        int             `json:"total"`
}

type BasketSubList struct {
	SubBasketID uint   `json:"sub_basket_id"`
	ProductID   uint   `json:"product_id"`
	ProductName string `json:"product_name"`
	Prince      int    `json:"prince"`
	Quantinty   int    `json:"quantity"`
	SubTotal    int    `json:"subtotal"`
	Stock       string `json:"stock_product"`
	Image       string `json:"image"`
}
