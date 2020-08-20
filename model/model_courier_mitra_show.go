package model

type CourierMitraShow struct {
	MitraID     uint   `json:"mitra_id"`
	CourierID   uint   `json:"courier_id"`
	CourierName string `json:"courier_name"`
}
