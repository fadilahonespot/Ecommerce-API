package model

type GetCityByProvince struct {
	RajaOngkir DataCityByProvince `json:"rajaongkir"`
}

type DataCityByProvince struct {
	Query  Query        `json:"query"`
	Status StatusRespon `json:"status"`
	Result []City         `json:"results"`
}