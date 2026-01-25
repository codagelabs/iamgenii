package models

import "time"

//MobileVarification **
type MobileVarification struct {
	VarificationID uint64     `gorm:"varification_id"`
	MobileNumber   string     `gorm:"mobile_no"`
	OTP            string     `gorm:"otp"`
	UserType       string     `gorm:"UserType"`
	UserID         uint64     `gorm:"user_id"`
	CreatedAt      time.Time  `gorm:"created_at" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"updated_at" json:"updated_at"`
	DeletedAt      *time.Time `gorm:"deleted_at" json:"deleted_at"`
}
