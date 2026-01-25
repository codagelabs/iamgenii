package services

import (
	"context"

	imgError "github.com/iamgenii/error"

	log "github.com/iamgenii/logs"

	"github.com/iamgenii/models"
	repository "github.com/iamgenii/svc/cities/pkg/v1/repositories"
)

type CitiesService interface {
	CreateCities(context.Context, models.Cities) (interface{}, *imgError.IMGNError)
	GetCities(context.Context) (*[]models.Cities, *imgError.IMGNError)
}

// citiesService having dependencies for cities repository
type citiesService struct {
	citiesRepo         repository.CitiesRepository
	repoErrInterceptor imgError.RepoErrorInterceptor
}

func NewCitiesService(citiesRepo repository.CitiesRepository, repoErrInterceptor imgError.RepoErrorInterceptor) CitiesService {
	return &citiesService{
		citiesRepo:         citiesRepo,
		repoErrInterceptor: repoErrInterceptor,
	}
}

func (b *citiesService) CreateCities(ctx context.Context, req models.Cities) (interface{}, *imgError.IMGNError) {
	log.Logger(ctx).Debug("CitiesService.CreateCities: Request data: ", req)
	resp, err := b.citiesRepo.CreateCities(ctx, req)
	if err != nil {
		log.Logger(ctx).Error("CitiesService.CreateCities: Error raised by repo. Error: ", err)
		return nil, b.repoErrInterceptor.ErrorMapper(ctx, err)
	}
	return resp, nil

}

func (b *citiesService) GetCities(ctx context.Context) (*[]models.Cities, *imgError.IMGNError) {
	log.Logger(ctx).Debug("CitiesService.GetCities: Request for get all cities: ")
	resp, err := b.citiesRepo.GetCities(ctx)
	if err != nil {
		log.Logger(ctx).Error("CitiesService.GetCities: Error raised by repo. Error: ", err)
		return nil, b.repoErrInterceptor.ErrorMapper(ctx, err)
	}
	return resp, nil
}
