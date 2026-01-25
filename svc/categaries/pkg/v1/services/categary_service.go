package services

import (
	"context"

	imgnErr "github.com/iamgenii/error"

	log "github.com/iamgenii/logs"

	"github.com/iamgenii/models"
	repository "github.com/iamgenii/svc/categaries/pkg/v1/repositories"
)

// CategoriesService describes the service interface for categories.
type CategoriesService interface {
	CreateCategories(ctx context.Context, createReq models.Categories) (interface{}, *imgnErr.IMGNError)
	GetCategoriesByID(context.Context, string) (*models.Categories, *imgnErr.IMGNError)
	GetSubCategoriesByID(context.Context, string) ([]*models.Categories, *imgnErr.IMGNError)
	DeleteCategories(context.Context, string) (interface{}, *imgnErr.IMGNError)
}

// categoriesService having dependencies for categories repository
type categoriesService struct {
	categoriesRepositories repository.CategoriesRepository
	repoErrorInterceptor   imgnErr.RepoErrorInterceptor
}

// NewCategoriesServiceImpl inject dependencies categories repository
func NewCategoriesServiceImpl(categoriesRepo repository.CategoriesRepository, repoErrorInterceptor imgnErr.RepoErrorInterceptor) CategoriesService {

	return &categoriesService{
		categoriesRepositories: categoriesRepo,
		repoErrorInterceptor:   repoErrorInterceptor,
	}
}

// CreateCategories LLL
func (categoriesSvc *categoriesService) CreateCategories(ctx context.Context,
	req models.Categories) (interface{}, *imgnErr.IMGNError) {

	log.Logger(ctx).Debug("CategoriesService.CreateCategories: Request : ", req)
	resp, err := categoriesSvc.categoriesRepositories.CreateCategories(ctx, req)
	if err != nil {
		log.Logger(ctx).Error("CategoriesService.CreateCategories: Error raised by repo : Error  : ", err)
		return nil, categoriesSvc.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	return resp, nil
}

// GetCategoriesByID calls an GetCategoriesByID method of repository to retrieve records
func (categoriesSvc *categoriesService) GetCategoriesByID(ctx context.Context,
	id string) (*models.Categories, *imgnErr.IMGNError) {
	log.Logger(ctx).Debug(" CategoriesService.GetCategoriesByID: Request : ", id)
	resp, err := categoriesSvc.categoriesRepositories.GetCategoriesByID(ctx, id)
	if err != nil {
		log.Logger(ctx).Error("CategoriesService.GetCategoriesByID: Error raised by repo : Error  : ", err)
		return nil, categoriesSvc.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	return resp, nil
}

// GetSubCategoriesByID call GetCategories function of repository to get all users
func (categoriesSvc *categoriesService) GetSubCategoriesByID(ctx context.Context,
	id string) ([]*models.Categories, *imgnErr.IMGNError) {
	log.Logger(ctx).Debug("CategoriesService.GetSubCategoriesByID Request : ", id)
	resp, err := categoriesSvc.categoriesRepositories.GetSubCategoriesByID(ctx, id)
	if err != nil {
		log.Logger(ctx).Error("CategoriesService.GetSubCategoriesByID: Error raised by repo : Error  : ", err)
		return nil, categoriesSvc.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	return resp, nil

}

// DeleteCategories  send request to delete customer from db
func (categoriesSvc *categoriesService) DeleteCategories(ctx context.Context,
	id string) (interface{}, *imgnErr.IMGNError) {
	log.Logger(ctx).Info("CategoriesService.DeleteCustomer req  Id : ", id)
	resp, err := categoriesSvc.categoriesRepositories.DeleteCategoriesByID(ctx, id)
	if err != nil {
		log.Logger(ctx).Error("CategoriesService.DeleteCategories: Error raised by repo : Error  : ", err)
		return nil, categoriesSvc.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	return resp, nil
}
