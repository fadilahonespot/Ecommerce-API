package model

type Cost struct {
	Value int    `json:"value"`
	Etd   string `json:"etd"`
	Note  string `json:"note"`
}

type Costs struct {
	Service     string `json:"service"`
	Description string `json:"description"`
	Cost        []Cost `json:"cost"`
}

type DataShipping struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Costs []Costs `json:"costs"`
}

type QueryDetail struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Weight      int    `json:"weight"`
	Courier     string `json:"courier"`
}

type DetailShipping struct {
	Query             QueryDetail    `json:"query"`
	Status            StatusRespon   `json:"status"`
	OriginDetail      City           `json:"origin_details"`
	DestinationDetail City           `json:"destination_details"`
	Result            []DataShipping `json:"results"`
}

type DetailShippingMap struct {
	RajaOngkir DetailShipping `json:"rajaongkir"`
}
