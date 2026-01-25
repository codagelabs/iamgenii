package repositories

import (
	"context"

	log "github.com/iamgenii/logs"

	"github.com/iamgenii/models"
	"github.com/jinzhu/gorm"
)

// CitiesRepository implements all methods in CitiesRepository
type CitiesRepository interface {
	CreateCities(context.Context, models.Cities) (interface{}, error)
	GetCities(context.Context) (*[]models.Cities, error)
}

type citiesRepository struct {
	dbConn *gorm.DB
}

// NewCitiesRepository inject dependencies of DataStore
func NewCitiesRepository(dbConn *gorm.DB) CitiesRepository {
	return &citiesRepository{dbConn: dbConn}
}

// CreateCities create Cities entry in database
func (repo citiesRepository) CreateCities(ctx context.Context, cities models.Cities) (interface{}, error) {
	dbConn := repo.dbConn
	if err := dbConn.Create(&cities).Scan(&cities).Error; err != nil {
		log.Logger(ctx).Error("CitiesRepository.CreateCities : Error : ", err)
		return nil, err
	}
	return &cities, nil
}

// GetCitiesByID retries  records by provided cities Id
func (repo citiesRepository) GetCities(ctx context.Context) (*[]models.Cities, error) {
	dbConn := repo.dbConn
	var cities []models.Cities
	err := dbConn.Table("cities").Find(&cities).Error
	if err != nil {
		log.Logger(ctx).Error("CitiesRepository.GetCities : Error : ", err)
		return nil, err
	}
	return &cities, nil
}
