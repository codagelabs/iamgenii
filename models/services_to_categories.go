package models

import "time"

type ServicesToCategories struct {
	ServicesId   uint64     `json:"services_id" gorm:"column:services_id"`
	CategoriesId uint64     `json:"categories_id" gorm:"column:categories_id"`
	DeletedAt    *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type ServicesToCategoriesReq struct {
	ServicesId   uint64   `json:"services_id" gorm:o "column:services_id" validator:"required"`
	CategoriesId []uint64 `json:"categories_ids" gorm:"column:categories_id" validator:"required"`
}

type GetServicesByCategoriesReq struct {
	Limit      int
	Page       int
	CategoryId uint64
}
