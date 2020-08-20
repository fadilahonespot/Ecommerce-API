package model

type Auth struct {
	Token string `json:"token"`
}

type AuthDetail struct {
	ID       int    `json:"id"`
	Password string `json:"pass"`
}
