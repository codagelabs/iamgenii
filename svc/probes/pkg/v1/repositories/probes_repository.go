package repositories

import (
	"context"

	"github.com/jinzhu/gorm"
)

// ProbesRepository implements all methods in ProbesRepository
type ProbesRepository interface {
	HealthCheck(context.Context) error
}

type probesRepository struct {
	dbConn *gorm.DB
}

// NewProbesRepository inject dependencies of DataStore
func NewProbesRepository(dbConn *gorm.DB) ProbesRepository {
	return &probesRepository{dbConn: dbConn}
}

// HealthCheck checks database connection
func (repo *probesRepository) HealthCheck(ctx context.Context) error {
	return repo.dbConn.DB().Ping()
}
