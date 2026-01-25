package models

import "time"

type PackageServiceMapping struct {
	PackagesId   uint64     `json:"package_id" gorm:"column:packages_id"`
	ServicesId uint64     `json:"service_id" gorm:"column:services_id"`
	DeletedAt    *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type PackageServiceMappingReq struct {
	PackageId   uint64   `json:"package_id" gorm:o "column:package_id" validator:"required"`
	ServicesIds []uint64 `json:"services_ids" gorm:"column:services_ids" validator:"required"`
}

type GetPackageServiceMappingReq struct {
	Limit      int
	Page       int
	PackageId uint64
}
