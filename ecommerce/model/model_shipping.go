package model

type Shipping struct {
	OriginDetail      City           `json:"origin_details"`
	DestinationDetail City           `json:"destination_details"`
	Result            []DataShipping `json:"results"`
}