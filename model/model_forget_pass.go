package model

type InputEmailForget struct {
	Email string `json:"email"`
}

type ShowAuthForgetPass struct {
	Auth   string `json:"auth_code"`
	Active string `json:"active"`
}
