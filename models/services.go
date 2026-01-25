package models

import (
	"time"
)

type Services struct {
	ServicesID          uint64              `json:"services_id" gorm:"primary_key;column:services_id" `
	ServicesLiked       uint64              `json:"services_liked" gorm:"column:services_liked"`
	IsSpecial           uint8               `json:"is_special" gorm:"column:is_special"`
	ServicesSlug        string              `json:"services_slug" gorm:"column:services_slug" validate:"required"`
	AddedBy             uint64              `json:"added_by" gorm:"column:added_by"validate:"required"`
	ServiceDescriptions ServiceDescription `json:"service_descriptions" gorm:"ForeignKey:ServicesID"`
	//Categories          []Categories        `json:"categories" gorm:"many2many:services_to_categories;association_foreignkey:services_id;foreignkey:services_id"`
	ServicesImageUrl    string              `json:"services_image_url" gorm:"column:services_image_url"`
	ServicesIconUrl     string              `json:"services_icon_url" gorm:"column:services_icon_url"`
	CreatedAt           time.Time           `json:"created_at" gorm:"column:created_at"`
	UpdatedAt           time.Time           `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt           *time.Time          `json:"deleted_at" gorm:"column:deleted_at"`
}

type ServiceDescription struct {
	ServicesID          uint64     `json:"services_id" gorm:"column:services_id"`
	ServicesName        string     `json:"services_name" gorm:"column:services_name" validate:"required"`
	ServicesDescription string     `json:"services_description" gorm:"column:services_description" validate:"required"`
	ServicesUrl         string     `json:"services_url" gorm:"column:services_url"`
	ServicesViewed      int64      `json:"services_viewed" gorm:"column:services_viewed"`
	ServicesBannerUrl   string     `json:"services_banner" gorm:"column:services_banner_url"`
	DeletedAt           *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type ServicesImages struct {
	ImageID     uint64    `json:"id" gorm:"column:id"`
	ServicesID  uint64    `json:"services_id" gorm:"column:services_id"`
	ImageUrl    string    `json:"image" gorm:"column:image"`
	HtmlContent string    `json:"html_content" gorm:"column:html_content"`
	SortOrder   string    `json:"sort_order" gorm:"column:sort_order"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

//GetAllServiceRequest consist parameters to filter out all customers data
type GetAllServiceRequest struct {
	Limit int
	Page  int
}


func (ServiceDescription) TableName() string {
	return "services_description"
}
