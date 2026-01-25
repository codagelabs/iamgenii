package services

import (
	"context"

	imgnError "github.com/iamgenii/error"

	log "github.com/iamgenii/logs"

	"github.com/iamgenii/models"
	repository "github.com/iamgenii/svc/services/pkg/v1/repositories"
)

type IamgeniiServices interface {
	CreateServices(context context.Context, createServiceReq models.Services) (createServiceResp interface{}, err *imgnError.IMGNError)
	GetServiceById(context context.Context, id string) (resp *models.Services, err *imgnError.IMGNError)
	GetAllService(context context.Context, req models.GetAllServiceRequest) (resp *[]models.Services, err *imgnError.IMGNError)
	UpdateService(context context.Context, updateServiceReq models.Services, serviceId string) (resp interface{}, err *imgnError.IMGNError)
	DeleteServiceById(context context.Context, Id string) (resp interface{}, err *imgnError.IMGNError)
}
type iamgeniiServices struct {
	servicesRepo         repository.ServicesRepository
	repoErrorInterceptor imgnError.RepoErrorInterceptor
}

func NewIamgeniiServices(servicesRepo repository.ServicesRepository, repoErrorInterceptor imgnError.RepoErrorInterceptor) IamgeniiServices {
	return &iamgeniiServices{
		servicesRepo:         servicesRepo,
		repoErrorInterceptor: repoErrorInterceptor,
	}
}

// CreateServices create service and return
func (iamgeniiServices *iamgeniiServices) CreateServices(context context.Context, createServiceReq models.Services) (interface{}, *imgnError.IMGNError) {
	log.Logger(context).Debug("IamgeniiServices.CreateServices Request : ", createServiceReq)
	insertedServiceResponse, err := iamgeniiServices.servicesRepo.InsertServiceData(context, createServiceReq)
	if err != nil {
		log.Logger(context).Error("IamgeniiServices.CreateServices: Error raised by repository. Error: ", createServiceReq)
		return nil, iamgeniiServices.repoErrorInterceptor.ErrorMapper(context, err)
	}
	return insertedServiceResponse, nil
}

func (iamgeniiServices *iamgeniiServices) GetServiceById(context context.Context, Id string) (*models.Services, *imgnError.IMGNError) {
	log.Logger(context).Debug("IamgeniiServices.GetServiceById: Request service id: ", Id)
	servicesResponse, err := iamgeniiServices.servicesRepo.FetchServiceById(context, Id)
	if err != nil {
		log.Logger(context).Error("IamgeniiServices.CreateServices: Error raised by repository. Error: ", err)
		return nil, iamgeniiServices.repoErrorInterceptor.ErrorMapper(context, err)
	}
	return servicesResponse, nil
}

func (iamgeniiServices *iamgeniiServices) GetAllService(context context.Context, req models.GetAllServiceRequest) (*[]models.Services, *imgnError.IMGNError) {
	log.Logger(context).Debug("IamgeniiServices.GetAllPackages: Request serviceId: ", req)
	resp, err := iamgeniiServices.servicesRepo.FetchAllServicesData(context, req)
	if err != nil {
		log.Logger(context).Error("IamgeniiServices.CreateServices: Error raised by repository. Error: ", err)
		return nil, iamgeniiServices.repoErrorInterceptor.ErrorMapper(context, err)
	}
	return resp, nil
}
func (iamgeniiServices *iamgeniiServices) UpdateService(context context.Context, updateServiceReq models.Services, serviceId string) (interface{}, *imgnError.IMGNError) {
	log.Logger(context).Debug("IamgeniiServices.UpdateService: Update request data: ", updateServiceReq, " Update request id: ", serviceId)
	resp, err := iamgeniiServices.servicesRepo.UpdateServiceData(context, updateServiceReq, serviceId)
	if err != nil {
		log.Logger(context).Error("IamgeniiServices.CreateServices: Error raised by repository. Error: ", err)
		return nil, iamgeniiServices.repoErrorInterceptor.ErrorMapper(context, err)
	}
	return resp, nil
}

func (iamgeniiServices *iamgeniiServices) DeleteServiceById(context context.Context, Id string) (interface{}, *imgnError.IMGNError) {
	log.Logger(context).Debug("IamgeniiServices.GetServiceById: Request service id: ", Id)
	err := iamgeniiServices.servicesRepo.SoftDeleteServiceData(context, Id)
	if err != nil {
		log.Logger(context).Error("IamgeniiServices.DeleteServiceById: Error raised by repository. Error: ", err)
		return nil, iamgeniiServices.repoErrorInterceptor.ErrorMapper(context, err)
	}
	return "Record Deleted Successfully", nil
}
