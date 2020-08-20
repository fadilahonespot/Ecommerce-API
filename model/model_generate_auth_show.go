package model

type AuthEmailShow struct {
	Auth   string `json:"auth"`
	Active string `json:"active"`
}

type AuthHPShow struct {
	Otp    string `json:"otp_code"`
	Active string `json:"active"`
}
