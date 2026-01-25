package models

import "time"

//Categories defination
type Categories struct {
	CategoriesID          uint64                  `gorm:"categories_id" json:"categories_id"`
	CategoriesName        string                  `gorm:"categories_name" json:"categories_name"`
	CategoriesIcon        string                  `gorm:"categories_icon" json:"categories_icon"`
	CategoriesImage       string                  `gorm:"categories_image" json:"categories_image"`
	ParentCategoriesID    uint                    `gorm:"parent_categories_id" json:"parent_categories_id"`
	CategoriesSlug        string                  `gorm:"categories_slug" json:"categories_slug"`
	CategoriesDescription []CategoriesDescription `gorm:"discription" json:"discription"`
	CreatedAt             time.Time               `gorm:"created_at" json:"created_at"`
	UpdatedAt             time.Time               `gorm:"updated_at" json:"updated_at"`
	DeletedAt             *time.Time              `gorm:"deleted_at" json:"deleted_at"`
}

//CategoriesDescription **
type CategoriesDescription struct {
	CategoriesDescriptionID uint64     `gorm:"categories_description_id" json:"categories_description_id"`
	CategoriesID            uint64     `gorm:"categories_id" json:"categories_id"`
	CategoriesDescription   string     `gorm:"categories_description" json:"categories_description"`
	LanguageID              uint64     `gorm:"language_id" json:"language_id"`
	CreatedAt               time.Time  `gorm:"created_at" json:"created_at"`
	UpdatedAt               time.Time  `gorm:"updated_at" json:"updated_at"`
	DeletedAt               *time.Time `gorm:"deleted_at" json:"deleted_at"`
}

//SubCategoriesReq hots req parameters fro categaries
type SubCategoriesReq struct {
	ParentCategoriesID uint `gorm:"parent_categary_id" json:"parent_categary_id"`
}
