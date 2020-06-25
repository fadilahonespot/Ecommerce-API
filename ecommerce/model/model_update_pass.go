package model

type UpdatePassword struct {
	OldPasswor string `json:"old_pass"`
	NewPassword string `json:"new_pass"`
}