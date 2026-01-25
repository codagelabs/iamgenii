package models

import "time"

type Vendor struct {
	VendorId           uint64     `json:"vendors_id" gorm:"column:vendors_id;primary_key"`
	Name               string     `json:"name" gorm:"column:name" validate:"required"`
	BusinessName       string     `json:"business_name" gorm:"column:business_name" validate:"required"`
	BusinessAddress    string     `json:"business_address" gorm:"column:business_address" validate:"required"`
	Email              string     `json:"email" gorm:"column:email" validate:"required"`
	City               string     `json:"city" gorm:"column:city" validate:"required"`
	PanCardId          string     `json:"pan_card_id" gorm:"column:pan_card_id" validate:"required"`
	GstNo              string     `json:"gst_no" gorm:"column:gst_no" validate:"required"`
	MobileNumber       string     `json:"mobile_no" gorm:"column:mobile_no" validate:"required"`
	Password           string     `json:"password" gorm:"column:password" validate:"required"`
	VendorPicture      string     `json:"vendors_picture" gorm:"column:vendors_picture" validate:"required"`
	Status             string     `json:"status" gorm:"column:status"`
	Ratting            int        `json:"rating" gorm:"column:rating"`
	Otp                string     `json:"otp" gorm:"column:otp"`
	ActivationStatus   string     `json:"activation_status" gorm:"column:activation_status"`
	ActivationDate     string     `json:"activation_date" gorm:"column:activation_date"`
	DeactivationReason string     `json:"deactivation_reason" gorm:"column:deactivation_reason"`
	FcmToken           string     `json:"fcm_token" gorm:"column:fcm_token"`
	CreatedAt          time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt          time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt          *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type VendorResp struct {
	VendorId           uint64     `json:"vendors_id" gorm:"column:vendors_id;primary_key"`
	Name               string     `json:"name" gorm:"column:name" validate:"required"`
	BusinessName       string     `json:"business_name" gorm:"column:business_name" validate:"required"`
	BusinessAddress    string     `json:"business_address" gorm:"column:business_address" validate:"required"`
	Email              string     `json:"email" gorm:"column:email" validate:"required"`
	City               string     `json:"city" gorm:"column:city" validate:"required"`
	PanCardId          string     `json:"pan_card_id" gorm:"column:pan_card_id" validate:"required"`
	GstNo              string     `json:"gst_no" gorm:"column:gst_no" validate:"required"`
	MobileNumber       string     `json:"mobile_no" gorm:"column:mobile_no" validate:"required"`
	VendorPicture      string     `json:"vendors_picture" gorm:"column:vendors_picture" validate:"required"`
	Status             string     `json:"status" gorm:"column:status"`
	Ratting            int        `json:"rating" gorm:"column:rating"`
	Otp                string     `json:"otp" gorm:"column:otp"`
	ActivationStatus   string     `json:"activation_status" gorm:"column:activation_status"`
	ActivationDate     string     `json:"activation_date" gorm:"column:activation_date"`
	DeactivationReason string     `json:"deactivation_reason" gorm:"column:deactivation_reason"`
	FcmToken           string     `json:"fcm_token" gorm:"column:fcm_token"`
	CreatedAt          time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt          time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt          *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type GetAllVendorReq struct {
	Limit int `gorm:"email" json:"email"`
	Page  int `gorm:"name" json:"name"`
}

type VendorUpdateReq struct {
	VendorId           uint64     `json:"vendors_id" gorm:"column:vendors_id;primary_key"`
	Name               string     `json:"name" gorm:"column:name" validate:"required"`
	BusinessName       string     `json:"business_name" gorm:"column:business_name" validate:"required"`
	BusinessAddress    string     `json:"business_address" gorm:"column:business_address" validate:"required"`
	Email              string     `json:"email" gorm:"column:email" validate:"required"`
	City               string     `json:"city" gorm:"column:city" validate:"required"`
	PanCardId          string     `json:"pan_card_id" gorm:"column:pan_card_id" validate:"required"`
	GstNo              string     `json:"gst_no" gorm:"column:gst_no" validate:"required"`
	MobileNumber       string     `json:"mobile_no" gorm:"column:mobile_no" validate:"required"`
	VendorPicture      string     `json:"vendors_picture" gorm:"column:vendors_picture" validate:"required"`
	Status             string     `json:"status" gorm:"column:status"`
	Ratting            int        `json:"rating" gorm:"column:rating"`
	Otp                string     `json:"otp" gorm:"column:otp"`
	ActivationStatus   string     `json:"activation_status" gorm:"column:activation_status"`
	ActivationDate     string     `json:"activation_date" gorm:"column:activation_date"`
	DeactivationReason string     `json:"deactivation_reason" gorm:"column:deactivation_reason"`
	FcmToken           string     `json:"fcm_token" gorm:"column:fcm_token"`
}