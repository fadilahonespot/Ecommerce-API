package model

type GetCity struct {
	RajaOngkir DataGetCity `json:"rajaongkir"`
}

type DataGetCity struct {
	Query  []string     `json:"query"`
	Status StatusRespon `json:"status"`
	Result []City       `json:"results"`
}
