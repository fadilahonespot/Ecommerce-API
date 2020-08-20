package model

type Provinces struct {
	RajaOngkir ProvinceData `json:"rajaongkir"`
}

type ProvinceData struct {
	Query  []string     `json:"query"`
	Status StatusRespon `json:"status"`
	Result []Province   `json:"results"`
}

type Province struct {
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
}
