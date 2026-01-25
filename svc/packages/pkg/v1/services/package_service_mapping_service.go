package services

import (
	"context"

	imgnErr "github.com/iamgenii/error"

	log "github.com/iamgenii/logs"

	"github.com/iamgenii/models"
	repository "github.com/iamgenii/svc/packages/pkg/v1/repositories"
)

type PackagesServicesMappingService interface {
	CreatePackageServiceMapping(context context.Context, req models.PackageServiceMappingReq) (interface{}, *imgnErr.IMGNError)
	GetPackageServices(context context.Context, req models.GetPackageServiceMappingReq) ([]models.Services, *imgnErr.IMGNError)
	DeleteServiceFromPackage(context context.Context, packageId, serviceId uint64) (string, *imgnErr.IMGNError)
}
type packagesServicesMappingService struct {
	packageServiceMappingRepo repository.PackageServiceMappingRepository
	repoErrorInterceptor      imgnErr.RepoErrorInterceptor
}

func NewPackagesServicesMappingService(packageServiceMappingRepo repository.PackageServiceMappingRepository,
	repoErrorInterceptor imgnErr.RepoErrorInterceptor) PackagesServicesMappingService {
	return &packagesServicesMappingService{
		packageServiceMappingRepo: packageServiceMappingRepo,
		repoErrorInterceptor:      repoErrorInterceptor,
	}
}

func (pkgSvcMapping packagesServicesMappingService) CreatePackageServiceMapping(context context.Context, req models.PackageServiceMappingReq) (interface{}, *imgnErr.IMGNError) {
	log.Logger(context).Debug("PackagesServicesMappingService.CreatePackageServiceMapping: Request: ", req)
	pkgSvcMap := make([]models.PackageServiceMapping, 0)
	for _, record := range req.ServicesIds {
		pkgSvcMap = append(pkgSvcMap, models.PackageServiceMapping{
			PackagesId: req.PackageId,
			ServicesId: record,
		})
	}
	resp, err := pkgSvcMapping.packageServiceMappingRepo.InsertPackagesToServicesRecords(context, pkgSvcMap)
	if err != nil {
		log.Logger(context).Error("PackagesServicesMappingService.CreatePackageServiceMapping: Error raised by repository. Error: ", err)
		return nil, pkgSvcMapping.repoErrorInterceptor.ErrorMapper(context, err)
	}
	return resp, nil
}

func (pkgSvcMapping packagesServicesMappingService) GetPackageServices(context context.Context, req models.GetPackageServiceMappingReq) ([]models.Services, *imgnErr.IMGNError) {
	log.Logger(context).Info("PackagesServicesMappingService.GetPackageServices Request :  ", req)
	resp, err := pkgSvcMapping.packageServiceMappingRepo.GetAllPackageServices(context, req)
	if err != nil {
		log.Logger(context).Error("PackagesServicesMappingService.GetPackageServices: Error raised by repository. Error: ", err)
		return nil, pkgSvcMapping.repoErrorInterceptor.ErrorMapper(context, err)
	}
	return resp, nil

}

func (pkgSvcMapping packagesServicesMappingService) DeleteServiceFromPackage(context context.Context, packageId, serviceId uint64) (string, *imgnErr.IMGNError) {
	log.Logger(context).Info("PackagesServicesMappingService.DeleteServiceFromPackage: Request package id: ", packageId, " serviceId : ", serviceId)
	err := pkgSvcMapping.packageServiceMappingRepo.DeletePackageServiceRecord(context, serviceId, packageId)
	if err != nil {
		log.Logger(context).Error("PackagesServicesMappingService.DeleteServiceFromPackage: Error raised by repository. Error: ", err)
		return "", pkgSvcMapping.repoErrorInterceptor.ErrorMapper(context, err)
	}
	return "Record Deleted Successfully", nil

}
