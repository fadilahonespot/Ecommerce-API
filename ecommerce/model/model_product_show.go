package model

type ProductShow struct {
	DateCreate      string   `json:"date_create"`
	ProductID       uint     `json:"product_id"`
	Name            string   `json:"name"`
	Prince          int      `json:"prince"`
	Stock           string      `json:"stock"`
	Weight          int      `json:"weight"`
	Condition       string   `json:"condition"`
	Brand           string   `json:"brand"`
	Description     string   `json:"description"`
	MinPurchase     int      `json:"minimal_purchase"`
	Sold            int      `json:"sold"`
	MitraID         uint     `json:"mitra_id"`
	MitraCityID     int      `json:"mitra_city_id"`
	CatagoryID      uint     `json:"catagory_id"`
	SubcatagoryID   uint     `json:"subcatagory_id"`
	MitraName       string   `json:"mitra_name"`
	MitraCity       string   `json:"mitra_city"`
	CatagoryName    string   `json:"catagory_name"`
	SubCatagoryName string   `json:"subcatagory_name"`
	Rating          string   `json:"rating"`
	Images          []string `json:"images"`
}
