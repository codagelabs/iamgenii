package models

//LoginReq hold login req details
type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	UserType string `json:"user_type" validate:"required"`
}
