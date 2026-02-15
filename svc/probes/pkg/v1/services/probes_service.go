package services

import (
	"context"

	repository "github.com/iamgenii/svc/probes/pkg/v1/repositories"
)

// ProbesService interface for probes service
type ProbesService interface {
	HealthCheck(context.Context) error
}

// probesService implements ProbesService
type probesService struct {
	probesRepo repository.ProbesRepository
}

// NewProbesService create new probes service
func NewProbesService(probesRepo repository.ProbesRepository) ProbesService {
	return &probesService{
		probesRepo: probesRepo,
	}
}

// HealthCheck checks dependency health
func (s *probesService) HealthCheck(ctx context.Context) error {
	return s.probesRepo.HealthCheck(ctx)
}
