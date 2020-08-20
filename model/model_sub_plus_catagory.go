package model

type SubPlusCatagory struct {
	CatagoryID      uint   `json:"catagory_id"`
	SubCatagoryID   uint   `json:"subcatagory_id"`
	CatagoryName    string `json:"catagory_name"`
	SubCatagoryName string `json:"subcatagory_name"`
}
