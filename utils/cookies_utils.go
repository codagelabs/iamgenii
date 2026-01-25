package utils

import "net/http"

type Cookies interface {
	SetTokenCookies(w http.ResponseWriter, token string, usertype string)
	DeleteTokenCookies(w http.ResponseWriter)
	GetTokenCookies(w http.ResponseWriter)
}

type cookies struct {
}

func NewCookies() Cookies {
	return &cookies{}
}

func (cookie cookies) SetTokenCookies(w http.ResponseWriter, token string, userType string) {
	http.SetCookie(w, &http.Cookie{
		Name:   "id_token",
		MaxAge: 2,
		Value:  token,
		Path:   "",
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "user_type",
		MaxAge: 2,
		Value:  userType,
		Path:   "",
	})
}
func (cookie cookies) DeleteTokenCookies(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "user_type",
		MaxAge: 2,
		Value:  "",
		Path:   "",
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "auth_token",
		MaxAge: 2,
		Value:  "",
		Path:   "",
	})
}

func (cookie cookies) GetTokenCookies(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "user_type",
		MaxAge: 2,
		Value:  "",
		Path:   "",
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "auth_token",
		MaxAge: 2,
		Value:  "",
		Path:   "",
	})
}
