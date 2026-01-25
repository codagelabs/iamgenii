package services

import (
	"context"
	"fmt"

	imgnError "github.com/iamgenii/error"

	log "github.com/iamgenii/logs"

	"github.com/iamgenii/models"
	repository "github.com/iamgenii/svc/services/pkg/v1/repositories"
)

type IamgeniiCategoriesToServices interface {
	GetServiceByCategoriesId(context context.Context, req models.GetServicesByCategoriesReq) (*[]models.Services, *imgnError.IMGNError)
	ServiceToCategoryMapping(context context.Context, mappingRequest models.ServicesToCategoriesReq) (interface{}, *imgnError.IMGNError)
	DeleteServiceToCategoryMapping(ctx context.Context, serviceId, categoryId uint64) (interface{}, *imgnError.IMGNError)
}

type iamgeniiCategoriesToServices struct {
	servicesRepo         repository.ServicesToCategoriesRepository
	repoErrorInterceptor imgnError.RepoErrorInterceptor
}

func NewIamgeniiCategoriesToServices(servicesRepo repository.ServicesToCategoriesRepository) IamgeniiCategoriesToServices {
	return &iamgeniiCategoriesToServices{
		servicesRepo: servicesRepo,
	}
}

func (iamgeniiServices *iamgeniiCategoriesToServices) GetServiceByCategoriesId(context context.Context, req models.GetServicesByCategoriesReq) (*[]models.Services, *imgnError.IMGNError) {
	log.Logger(context).Debug("IamgeniiCategoriesToServices.GetServiceByCategoriesId: Request:  ", req)
	insertedServiceResponse, err := iamgeniiServices.servicesRepo.FetchServiceByCategoriesID(context, req)
	if err != nil {
		log.Logger(context).Debug("IamgeniiCategoriesToServices.GetServiceByCategoriesId: Error raised by repository. Error: ", err)
		return nil, iamgeniiServices.repoErrorInterceptor.ErrorMapper(context, err)
	}
	return insertedServiceResponse, nil

}
func (iamgeniiServices *iamgeniiCategoriesToServices) ServiceToCategoryMapping(context context.Context, mappingRequest models.ServicesToCategoriesReq) (interface{}, *imgnError.IMGNError) {
	log.Logger(context).Info("IamgeniiCategoriesToServices.ServiceToCategoryMapping Request : serviceId : ", mappingRequest)

	serviceToCategoriesReq := make([]models.ServicesToCategories, 0)

	for _, categoryId := range mappingRequest.CategoriesId {
		svcToCat := models.ServicesToCategories{
			CategoriesId: categoryId,
			ServicesId:   mappingRequest.ServicesId,
		}
		serviceToCategoriesReq = append(serviceToCategoriesReq, svcToCat)

	}
	fmt.Print(serviceToCategoriesReq)
	response, err := iamgeniiServices.servicesRepo.InsertPackagesToServicesRecords(context, serviceToCategoriesReq)
	if err != nil {
		log.Logger(context).Debug("IamgeniiCategoriesToServices.GetServiceByCategoriesId: Error raised by repository. Error: ", err)
		return nil, iamgeniiServices.repoErrorInterceptor.ErrorMapper(context, err)
	}
	return response, nil
}

func (iamgeniiServices *iamgeniiCategoriesToServices) DeleteServiceToCategoryMapping(ctx context.Context, serviceId, categoryId uint64) (interface{}, *imgnError.IMGNError) {

	log.Logger(ctx).Info("IamgeniiCategoriesToServices.DeleteServiceToCategoryMapping: Request  for service id: ", serviceId, " Category id: ", categoryId)
	err := iamgeniiServices.servicesRepo.DeleteServiceCategoryRecord(ctx, serviceId, categoryId)
	if err != nil {
		log.Logger(ctx).Debug("IamgeniiCategoriesToServices.DeleteServiceToCategoryMapping: Error raised by repository. Error: ", err)
		return nil, iamgeniiServices.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	return "Record deleted successfully", nil
}
