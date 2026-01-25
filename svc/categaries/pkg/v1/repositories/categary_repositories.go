package repositories

import (
	"context"
	"time"

	customError "github.com/iamgenii/error"

	log "github.com/iamgenii/logs"

	"github.com/iamgenii/models"
	"github.com/jinzhu/gorm"
)

// CategoriesRepository implements all methods in CategoriesRepository
type CategoriesRepository interface {
	CreateCategories(context.Context, models.Categories) (interface{}, error)
	GetCategoriesByID(context.Context, string) (*models.Categories, error)
	GetSubCategoriesByID(context.Context, string) ([]*models.Categories, error)
	DeleteCategoriesByID(context.Context, string) (interface{}, error)
}

// categoriesRepository
type categoriesRepository struct {
	dbConn *gorm.DB
}

// NewCategoriesRepository inject dependencies of DataStore
func NewCategoriesRepository(dbConn *gorm.DB) CategoriesRepository {
	return &categoriesRepository{dbConn: dbConn}
}

// CreateCategories create categories entry in database
func (categoriesRepo categoriesRepository) CreateCategories(ctx context.Context, categories models.Categories) (interface{}, error) {
	dbConn := categoriesRepo.dbConn
	if err := dbConn.Create(&categories).
		Scan(&categories).
		Error; err != nil {
		log.Logger(ctx).Error("CategoriesRepository.CreateCategories: Error: ", err)
		return nil, err
	}
	return &categories, nil
}

// GetCategoriesByID retries  records by provided categories Id
func (categoriesRepo categoriesRepository) GetCategoriesByID(ctx context.Context,
	id string) (*models.Categories, error) {

	dbConn := categoriesRepo.dbConn
	categories := models.Categories{}

	err := dbConn.Table("categories").
		Where("categories_id=?", id).
		First(&categories).
		Error
	if gorm.IsRecordNotFoundError(err) {
		log.Logger(ctx).Error("CategoriesRepository.GetCategoriesByID: Error: ", err)
		return nil, customError.ErrRecordNotFound
	}
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return &categories, nil
}

// GetSubCategoriesByID retries  records by provided categories Id
func (categoriesRepo categoriesRepository) GetSubCategoriesByID(ctx context.Context, id string) ([]*models.Categories, error) {

	dbConn := categoriesRepo.dbConn
	var categories []*models.Categories
	err := dbConn.Table("categories").
		Where("parent_categories_id=?", id).
		Find(&categories).
		Error
	if gorm.IsRecordNotFoundError(err) || len(categories) <= 0 {
		log.Logger(ctx).Error("CategoriesRepository.GetSubCategoriesByID: Error: ", err)
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return categories, nil
}

// DeleteCategoriesByID delete records by provided categories Id
func (categoriesRepo categoriesRepository) DeleteCategoriesByID(ctx context.Context, id string) (interface{}, error) {

	dbConn := categoriesRepo.dbConn
	categories := models.Categories{}
	if dbConn.Table("categories").
		Where("categories_id=?", id).
		First(&categories).
		RecordNotFound() {
		log.Logger(ctx).Error("CategoriesRepository.CreateCategories: Error: Record not found error ")
		return nil, gorm.ErrRecordNotFound
	}
	err := dbConn.Table("categories").
		Where("categories_id=?", id).
		Update("deleted_at", time.Now()).
		Error
	if err != nil {
		log.Logger(ctx).Error("CategoriesRepository.CreateCategories: Error: ", err)
		return nil, err
	}
	return "Record deleted successfully", nil
}
