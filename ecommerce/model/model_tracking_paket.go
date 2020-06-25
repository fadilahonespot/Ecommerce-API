package model

type TrackingPackage struct {
	Result  bool           `json:"result"`
	Data    DetailTracking `json:"data"`
	Message string         `json:"message"`
}

type DetailTracking struct {
	Courier  string         `json:"courier"`
	Waybill  string         `json:"waybill"`
	Shipped  interface{}    `json:"shipped"`
	Received interface{}    `json:"received"`
	Tracking []DataTracking `json:"tracking"`
	Status   string         `json:"status"`
}

type ReceivedPackage struct {
	Name      string `json:"name"`
	Recipient string `json:"recipient"`
	Address   string `json:"addr"`
	Date      string `json:"date"`
}

type ShippedCourier struct {
	Name    string `json:"name"`
	Address string `json:"addr"`
	Date    string `json:"date"`
}

type DataTracking struct {
	Date string `json:"date"`
	Desc string `json:"desc"`
	Status string `json:"status"`
}

type InputTracking struct {
	AWB     string `json:"awb"`
	Courier string `json:"courier"`
}
