package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Admin Model
type Admin struct {
	AdministratorsID uint64     `gorm:"administrators_id" json:"administrators_id" `
	FirstName        string     `gorm:"first_name" json:"first_name" validate:"required,alpha"`
	LastName         string     `gorm:"last_name" json:"last_name" validate:"required,alpha"`
	Email            string     `gorm:"email" json:"email"validate:"required,email"`
	Username         string     `gorm:"username" json:"username" validate:"required,alphanum,max=30,min=8"`
	Password         string     `gorm:"password" json:"password" validate:"required,min=8,max=20"`
	IsActive         bool       `gorm:"is_active DEFAULT:0" json:"is_active"`
	AdminTypeID      uint64     `gorm:"admin_type_id" json:"admin_type_id" validate:"required"`
	Address          string     `gorm:"address" json:"address"`
	City             string     `gorm:"city" json:"city"`
	State            string     `gorm:"state" json:"state"`
	Zip              string     `gorm:"zip" json:"zip"`
	Country          string     `gorm:"country" json:"country"`
	Phone            string     `gorm:"phone" json:"phone" validate:"required"`
	Image            string     `gorm:"image" json:"image"`
	RememberToken    string     `gorm:"remember_token" json:"remember_token"`
	CreatedAt        time.Time  `gorm:"created_at" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"updated_at" json:"updated_at"`
	DeletedAt        *time.Time `gorm:"deleted_at" json:"deleted_at"`
}

//AdminResp Model
type AdminResp struct {
	AdministratorsID uint64     `gorm:"administrators_id" json:"administrators_id"`
	FirstName        string     `gorm:"first_name" json:"first_name"`
	LastName         string     `gorm:"last_name" json:"last_name"`
	Email            string     `gorm:"email" json:"email"`
	IsActive         bool       `gorm:"is_active DEFAULT:0" json:"is_active"`
	AdminTypeID      uint64     `gorm:"admin_type_id" json:"admin_type_id"`
	Address          string     `gorm:"address" json:"address"`
	City             string     `gorm:"city" json:"city"`
	State            string     `gorm:"state" json:"state"`
	Zip              string     `gorm:"zip" json:"zip"`
	Country          string     `gorm:"country" json:"country"`
	Phone            string     `gorm:"phone" json:"phone"`
	Image            string     `gorm:"image" json:"image"`
	CreatedAt        time.Time  `gorm:"created_at" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"updated_at" json:"updated_at"`
	DeletedAt        *time.Time `gorm:"deleted_at" json:"deleted_at"`
}

//AdminType is the spefy the type of admin is
type AdminType struct {
	gorm.Model
	AdminTypeName string `gorm:"admin_type_name" json:"admin_type_name"`
}

//GetAllAdminReq return all user struct for creare filters for query
type GetAllAdminReq struct {
	Limit     int    `gorm:"email" json:"email"`
	Page      int    `gorm:"name" json:"name"`
	AdminType string `gorm:"uuid" json:"uuid"`
}

//AdminLoginRequest Model
type AdminLoginRequest struct {
	Email    string `json:"email" validate:"required" email:"true"`
	Password string `json:"password" validate:"required"`
}

//AdminLoginResponse Model
type AdminLoginResponse struct {
	Token     string `json:"token"`
	AdminType string `json:"user_type"`
}

//GetAdminsResponse Model
type GetAdminsResponse struct {
	Admins []AdminResp `json:"users"`
	Count  int         `json:"count"`
}

//UpdateAdminRequest Model
type UpdateAdminRequest struct {
	FirstName     string    `gorm:"first_name" json:"first_name"`
	LastName      string    `gorm:"last_name" json:"last_name"`
	Email         string    `gorm:"email" json:"email"`
	UUID          string    `gorm:"uuid" json:"uuid"`
	IsActive      bool      `gorm:"" json:""`
	AdminType     AdminType `gorm:"admin_type association_foreignkey:admin_type" json:"adminType"`
	Address       string    `gorm:"address" json:"address"`
	City          string    `gorm:"city" json:"city"`
	State         string    `gorm:"state" json:"state"`
	Zip           string    `gorm:"zip" json:"zip"`
	Country       string    `gorm:"country" json:"country"`
	Phone         string    `gorm:"phone" json:"phone"`
	Image         string    `gorm:"image" json:"image"`
	RememberToken string    `gorm:"remember_token" json:"remember_token"`
	AciveStatus   string    `gorm:"active_status" json:"active_status"`
}

//ForgotPasswordRequest **
type ForgotPasswordRequest struct {
	Email string `gorm:"email" json:"email"`
}

//ResetPasswordRequest **
type ResetPasswordRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

// UpdatePasswordRequest **
type UpdatePasswordRequest struct {
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

//GetAdminCourseSectionsReq **
type GetAdminCourseSectionsReq struct {
	ID        string
	CourseID  uint32
	SectionID uint32
}
