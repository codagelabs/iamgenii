package models

import (
	"time"
)

//Customer all customer details
type Customer struct {
	CustomersID               uint       `gorm:"primary_key column:customers_id" json:"customers_id"`
	CustomersGender           string     `gorm:"column:customers_gender" json:"customers_gender" validate:"required"`
	CustomersFirstName        string     `gorm:"column:customers_firstname" json:"customers_firstname" validate:"required"`
	CustomersLastName         string     `gorm:"column:customers_lastname" json:"customers_lastname"  validate:"required"`
	CustomersDob              string     `gorm:"column:customers_dob" json:"customers_dob"  validate:"required,date"`
	Email                     string     `gorm:"column:email" json:"email"  validate:"required,email"`
	UserName                  string     `gorm:"column:user_name" json:"user_name" validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	CustomersDefaultAddressID uint32     `gorm:"column:customers_default_address_id" json:"customers_default_address_id"`
	CustomersTelephone        string     `gorm:"column:customers_telephone" json:"customers_telephone"`
	CustomersPhone            string     `gorm:"column:customers_phone" json:"customers_phone"`
	CustomersFax              string     `gorm:"column:customers_fax" json:"customers_fax"`
	Password                  string     `gorm:"column:password" json:"password"`
	CustomersNewsletter       string     `gorm:"column:customers_newsletter" json:"customers_newsletter"`
	IsActive                  int8       `gorm:"column:is_active" json:"-"`
	FbID                      string     `gorm:"column:fb_id" json:"fb_id"`
	GoogleID                  string     `gorm:"column:google_id" json:"google_id"`
	IsSeen                    bool       `gorm:"column:is_seen" json:"is_seen"`
	RememberToken             string     `gorm:"column:remember_token" json:"remember_token"`
	CustomersPicture          string     `gorm:"column:customers_picture" json:"customers_picture"`
	CreatedAt                 time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                 time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt                 *time.Time `sql:"index" gorm:"column:deleted_at" json:"deleted_at"`
}

//GetAllCustReq consist parameters to filter out all customers data
type GetAllCustReq struct {
	Limit int `gorm:"email" json:"email"`
	Page  int `gorm:"name" json:"name"`
}

//CustLoginReq hold login req details
type CustLoginReq struct {
	Email    string `json:"email" validate:"required" email:"true"`
	Password string `json:"password" validate:"required"`
}

//CustLoginResp send data after login is sucessfull
type CustLoginResp struct {
	Token    string `json:"token"`
	UserType string `json:"usertype"`
}

//UpdateCustReq consist the feild that are updateble
type UpdateCustReq struct {
	CustomersID               uint       `gorm:"column:customers_id" json:"customers_id"`
	CustomersGender           string     `gorm:"column:customers_gender" json:"customers_gender" validate:"required"`
	CustomersFirstName        string     `gorm:"column:customers_firstname" json:"customers_firstname" validate:"required"`
	CustomersLastName         string     `gorm:"column:customers_lastname" json:"customers_lastname"  validate:"required"`
	CustomersDob              string     `gorm:"column:customers_dob" json:"customers_dob"  validate:"required,date"`
	Email                     string     `gorm:"column:email" json:"email"  validate:"required,email"`
	UserName                  string     `gorm:"column:user_name" json:"user_name" validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	CustomersDefaultAddressID uint32     `gorm:"column:customers_default_address_id" json:"customers_default_address_id"`
	CustomersTelephone        string     `gorm:"column:customers_telephone" json:"customers_telephone"`
	CustomersPhone            string     `gorm:"column:customers_phone" json:"customers_phone"`
	CustomersFax              string     `gorm:"column:customers_fax" json:"customers_fax"`
	Password                  string     `gorm:"column:password" json:"-"`
	CustomersNewsletter       string     `gorm:"column:customers_newsletter" json:"customers_newsletter"`
	IsActive                  int8       `gorm:"column:is_active" json:"-"`
	FbID                      string     `gorm:"column:fb_id" json:"fb_id"`
	GoogleID                  string     `gorm:"column:google_id" json:"google_id"`
	IsSeen                    bool       `gorm:"column:is_seen" json:"is_seen"`
	RememberTocken            string     `gorm:"column:remember_token" json:"remember_token"`
	Role                      string     `gorm:"column:role" json:"role"`
	OTP                       string     `gorm:"column:otp" json:"otp"`
	AdminApproval             bool       `gorm:"column:admin_approval" json:"admin_approval"`
	CustomersPicture          string     `gorm:"column:customers_picture" json:"customers_picture"`
	RejectReason              string     `gorm:"column:reject_reason" json:"reject_reason"`
	CreatedAt                 time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                 time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt                 *time.Time `sql:"index" gorm:"column:deleted_at" json:"deleted_at"`
}

//CustomerResp all customer details
type CustomerResp struct {
	CustomersID               uint       `gorm:"column:customers_id" json:"customers_id"`
	CustomersGender           string     `gorm:"column:customers_gender" json:"customers_gender" validate:"required"`
	CustomersFirstName        string     `gorm:"column:customers_firstname" json:"customers_firstname" validate:"required"`
	CustomersLastName         string     `gorm:"column:customers_lastname" json:"customers_lastname"  validate:"required"`
	CustomersDob              string     `gorm:"column:customers_dob" json:"customers_dob"  validate:"required,date"`
	Email                     string     `gorm:"column:email" json:"email"  validate:"required,email"`
	UserName                  string     `gorm:"column:user_name" json:"-" validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	CustomersDefaultAddressID uint32     `gorm:"column:customers_default_address_id" json:"customers_default_address_id"`
	CustomersTelephone        string     `gorm:"column:customers_telephone" json:"customers_telephone"`
	CustomersPhone            string     `gorm:"column:customers_phone" json:"customers_phone"`
	CustomersFax              string     `gorm:"column:customers_fax" json:"customers_fax"`
	Password                  string     `gorm:"column:password" json:"-"`
	CustomersNewsletter       string     `gorm:"column:customers_newsletter" json:"customers_newsletter"`
	IsActive                  int8       `gorm:"column:is_active" json:"-"`
	FbID                      string     `gorm:"column:fb_id" json:"fb_id"`
	GoogleID                  string     `gorm:"column:google_id" json:"google_id"`
	IsSeen                    bool       `gorm:"column:is_seen" json:"is_seen"`
	RememberTocken            string     `gorm:"column:remember_token" json:"remember_token"`
	Role                      string     `gorm:"column:role" json:"role"`
	OTP                       string     `gorm:"column:otp" json:"otp"`
	AddminApproval            bool       `gorm:"column:addmin_approval" json:"addmin_approval"`
	CustomersPicture          string     `gorm:"column:customers_picture" json:"customers_picture"`
	RejectReason              string     `gorm:"column:reject_reason" json:"reject_reason"`
	CreatedAt                 time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                 time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt                 *time.Time `sql:"index" gorm:"column:deleted_at" json:"deleted_at"`
}
