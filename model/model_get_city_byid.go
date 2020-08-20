package model

type GetCityByID struct {
	RajaOngkir DataCityByID `json:"rajaongkir"`
}

type DataCityByID struct {
	Query  Query        `json:"query"`
	Status StatusRespon `json:"status"`
	Result City         `json:"results"`
}