package models

import "time"

type Packages struct {
	PackagesID          uint64     `json:"packages_id" gorm:"primary_key;column:packages_id" `
	PackageName         string     `json:"packages_name" gorm:"column:packages_name" validate:"required"`
	PackagesPrice       float64    `json:"packages_price" gorm:"column:packages_price" validate:"required"`
	PackagesDescription string     `json:"packages_description" gorm:"column:packages_description" validate:"required"`
	PackagesStatus      uint64     `json:"packages_status" gorm:"column:packages_status"  validate:"required"`
	IsSpecial           uint8      `json:"is_special" gorm:"column:is_special"`
	PackagesSlug        string     `json:"packages_slug" gorm:"column:packages_slug" validate:"required"`
	AddedBy             uint64     `json:"added_by" gorm:"column:added_by"validate:"required"`
	PackagesTaxClassID  uint64     `json:"packages_tax_class_id" gorm:"column:packages_tax_class_id"`
	PackagesViewed      int64      `json:"packages_viewed" gorm:"column:packages_viewed"`
	PackagesUrl         string     `json:"packages_url" gorm:"column:packages_url"`
	PackagesLiked       uint64     `json:"packages_liked" gorm:"column:packages_liked"`
	PackagesBannerUrl   string     `json:"packages_banner" gorm:"column:packages_banner_url"`
	PackagesImageUrl    string     `json:"packages_image_url" gorm:"column:packages_image_url"`
	PackagesIconUrl     string     `json:"packages_icon_url" gorm:"column:packages_icon_url"`
	CreatedAt           time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt           time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt           *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type GetAllPackagesRequest struct {
	Limit int
	Page  int
}
