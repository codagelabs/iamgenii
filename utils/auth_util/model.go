package auth_util

import "github.com/dgrijalva/jwt-go"

type LoginMode string

const (
	LOGIN_MODE_WEB    LoginMode = "web"
	LOGIN_MODE_MOBILE LoginMode = ""
)

type UserType string

const (
	VENDOR_USER   UserType = "VENDOR"
	ADMIN_USER    UserType = "ADMIN"
	CUSTOMER_USER UserType = "CUSTOMER"
)

type AppType string

const (
	VENDOR_APP   AppType = "VENDOR_APP"
	ADMIN_APP    AppType = "ADMIN_PANEL"
	CUSTOMER_APP AppType = "CUSTOMER_APP"
)

type IdToken struct {
	FirstName    string
	LastName     string
	MobileNumber string
	UserType     []UserType
	UserId       string
	Email        string
	LoginMode    LoginMode
	AppType      []AppType
}

func (idToken IdToken) IsAdminLoggedIn() bool {
	if idToken.UserType[0] == ADMIN_USER && idToken.AppType[0] == ADMIN_APP {
		return true
	}
	return false
}

func (idToken IdToken) IsCustomerLoggedIn() bool {
	if idToken.UserType[0] == CUSTOMER_USER && idToken.AppType[0] == CUSTOMER_APP {
		return true
	}
	return false
}

func (idToken IdToken) IsVendorLoggedIn() bool {
	if idToken.UserType[0] == VENDOR_USER && idToken.AppType[0] == VENDOR_APP {
		return true
	}
	return false
}

// JWTClaims describes Claims in token.
type JWTClaims struct {
	idToken IdToken
	jwt.StandardClaims
}

//NewJWTClaims initialize tokens
func NewJWTClaims(idTocken IdToken) JWTClaims {
	return JWTClaims{
		idToken: idTocken,
	}
}
