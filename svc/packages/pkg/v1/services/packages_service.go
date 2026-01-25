package services

import (
	"context"

	imgnError "github.com/iamgenii/error"
	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
	repository "github.com/iamgenii/svc/packages/pkg/v1/repositories"
)

type PackagesServices interface {
	CreatePackages(context context.Context, req models.Packages) (resp *models.Packages, err *imgnError.IMGNError)
	UpdatePackages(context context.Context, updateReq models.Packages, packageId string) (respPackage *models.Packages, err *imgnError.IMGNError)
	GetPackagesById(context context.Context, packageId string) (respPackage *models.Packages, err *imgnError.IMGNError)
	GetAllPackages(context context.Context, req models.GetAllPackagesRequest) (resp *[]models.Packages, err *imgnError.IMGNError)
	DeletePackageByID(ctx context.Context, id string) (resp string, err *imgnError.IMGNError)
}
type packagesServices struct {
	packagesRepo         repository.PackagesRepository
	repoErrorInterceptor imgnError.RepoErrorInterceptor
}

func NewPackagesServices(servicesRepo repository.PackagesRepository, repoErrorInterceptor imgnError.RepoErrorInterceptor) PackagesServices {
	return &packagesServices{
		packagesRepo:         servicesRepo,
		repoErrorInterceptor: repoErrorInterceptor,
	}
}

func (packagesServices *packagesServices) CreatePackages(context context.Context, req models.Packages) (*models.Packages, *imgnError.IMGNError) {
	log.Logger(context).Debug("PackagesServices.CreatePackages: Create Packages Request: ", req)
	resp, err := packagesServices.packagesRepo.InsertPackageData(context, req)
	if err != nil {
		log.Logger(context).Error("PackagesServices.CreatePackages: Error raised by repository. Error: ", err)
		return nil, packagesServices.repoErrorInterceptor.ErrorMapper(context, err)
	}
	return resp, nil
}

func (packagesServices *packagesServices) UpdatePackages(context context.Context, updateReq models.Packages, packageId string) (*models.Packages, *imgnError.IMGNError) {
	log.Logger(context).Debug("PackagesServices.UpdatePackages: Update packages request : ", updateReq)
	respPackage, err := packagesServices.packagesRepo.UpdatePackageData(context, updateReq, packageId)
	if err != nil {
		log.Logger(context).Error("PackagesServices.UpdatePackages: Error raised by repository. Error: ", err)
		return nil, packagesServices.repoErrorInterceptor.ErrorMapper(context, err)
	}
	return respPackage, nil
}

func (packagesServices *packagesServices) GetPackagesById(context context.Context, packageId string) (*models.Packages, *imgnError.IMGNError) {
	log.Logger(context).Info("PackagesServices.GetPackagesById: Get packages by package id: ", packageId)
	respPackage, err := packagesServices.packagesRepo.GetPackageRecordsById(context, packageId)
	if err != nil {
		log.Logger(context).Error("PackagesServices.GetPackagesById: Error raised by repository. Error: ", err)
		return nil, packagesServices.repoErrorInterceptor.ErrorMapper(context, err)
	}
	return respPackage, nil
}

func (packagesServices *packagesServices) GetAllPackages(context context.Context, req models.GetAllPackagesRequest) (*[]models.Packages, *imgnError.IMGNError) {
	log.Logger(context).Info("PackagesServices.GetAllPackages: Get all packages by packages request: ", req)
	respPackage, err := packagesServices.packagesRepo.GetAllPackageRecords(context, req)
	if err != nil {
		log.Logger(context).Error("PackagesServices.GetAllPackages: Error raised by repository. Error: ", err)
		return nil, packagesServices.repoErrorInterceptor.ErrorMapper(context, err)
	}
	return respPackage, nil
}

func (packagesServices *packagesServices) DeletePackageByID(context context.Context, id string) (string, *imgnError.IMGNError) {
	log.Logger(context).Info("PackagesServices.DeletePackageByID: delete package by id: ", id)
	err := packagesServices.packagesRepo.DeletePackageRecordsById(context, id)
	if err != nil {
		log.Logger(context).Error("PackagesServices.DeletePackageByID: Error raised by repository. Error: ", err)
		return "", packagesServices.repoErrorInterceptor.ErrorMapper(context, err)
	}
	return "Record deleted successfully ", nil
}
