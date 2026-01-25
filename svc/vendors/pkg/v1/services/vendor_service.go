package services

import (
	"context"

	"github.com/iamgenii/email"
	imgnError "github.com/iamgenii/error"
	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
	repository "github.com/iamgenii/svc/vendors/pkg/v1/repositories"
	"github.com/iamgenii/utils/crypto_utils"
)

// VendorService describes the service.
type VendorService interface {
	CreateVendor(ctx context.Context, createReq models.Vendor) (interface{}, *imgnError.IMGNError)
	GetVendorByID(context.Context, string) (*models.VendorResp, *imgnError.IMGNError)
	GetVendors(context.Context, models.GetAllVendorReq) ([]*models.VendorResp, *imgnError.IMGNError)
	UpdateVendor(context.Context, string, models.VendorUpdateReq) (interface{}, *imgnError.IMGNError)
	DeleteVendor(context.Context, string) (interface{}, *imgnError.IMGNError)
}

type vendorService struct {
	vendorRepo           repository.VendorRepository
	sendBlue             *email.SendInBlue
	hashedUtils          crypto_utils.HashUtils
	repoErrorInterceptor imgnError.RepoErrorInterceptor
}

func NewVendorService(vendorRepo repository.VendorRepository,
	sendBlue *email.SendInBlue,
	hashedUtils crypto_utils.HashUtils,
	repoErrorInterceptor imgnError.RepoErrorInterceptor,
) VendorService {

	return &vendorService{
		vendorRepo:           vendorRepo,
		sendBlue:             sendBlue,
		hashedUtils:          hashedUtils,
		repoErrorInterceptor: repoErrorInterceptor,
	}
}

func (b *vendorService) CreateVendor(ctx context.Context, vendor models.Vendor) (interface{}, *imgnError.IMGNError) {
	log.Logger(ctx).Debug("VendorService.CreateVendor: Create Vendor Request: ", vendor)
	_, err := b.vendorRepo.GetVendorByEmail(ctx, vendor.Email)
	if err == nil {
		log.Logger(ctx).Error("VendorService.CreateVendor: Error vendor request emails record already exist. Error: ", err)
		return vendor, imgnError.ErrorEmailAlreadyExist
	}
	_, err = b.vendorRepo.GetVendorByPhone(ctx, vendor.MobileNumber)
	if err == nil {
		log.Logger(ctx).Error("VendorService.CreateVendor: Error vendor request phones record already exist. Error: ", err)
		return vendor, imgnError.ErrorMobileNumberAlreadyExist
	}
	vendor.Password, err = b.hashedUtils.GenerateBcrtptHash(vendor.Password)
	if err != nil {
		log.Logger(ctx).Error("VendorService.CreateVendor: Error in generate password hash. Error: ", err)
		return nil, b.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	resp, err := b.vendorRepo.CreateVendor(ctx, vendor)
	if err != nil {
		log.Logger(ctx).Error("VendorService.CreateVendor: Error raised by repository. Error: ", err)
		return nil, b.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	return resp, nil
}

func (b *vendorService) GetVendorByID(ctx context.Context, id string) (*models.VendorResp, *imgnError.IMGNError) {
	log.Logger(ctx).Debug("VendorService.GetVendorByID: Get vendor by id request id: ", id)
	resp, err := b.vendorRepo.GetVendorByID(ctx, id)
	if err != nil {
		log.Logger(ctx).Error("VendorService.GetVendorByID: Error raised by repository. Error: ", err)
		return nil, b.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	return resp, nil
}

func (b *vendorService) GetVendors(ctx context.Context, req models.GetAllVendorReq) ([]*models.VendorResp, *imgnError.IMGNError) {
	log.Logger(ctx).Debug("VendorService.GetVendors: Get all vendors request. Request: ", req)
	resp, err := b.vendorRepo.GetVendors(ctx, req)
	if err != nil {
		log.Logger(ctx).Error("VendorService.GetVendors: Error raised by repository. Error: ", err)
		return nil, b.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	return resp, nil
}

func (b *vendorService) UpdateVendor(ctx context.Context, id string, vendor models.VendorUpdateReq) (interface{}, *imgnError.IMGNError) {
	log.Logger(ctx).Debug("VendorService.UpdateVendor: update vendor request: ", vendor, " Vendor Id: ", id)
	resp, err := b.vendorRepo.UpdateVendor(ctx, id, vendor)
	if err != nil {
		log.Logger(ctx).Error("VendorService.UpdateVendor: Error raised by repository. Error: ", err)
		return nil, b.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	return resp, nil
}

func (b *vendorService) DeleteVendor(ctx context.Context, id string) (interface{}, *imgnError.IMGNError) {
	log.Logger(ctx).Info("VendorService.DeleteVendor: delete vendor request Id: ", id)
	resp, err := b.vendorRepo.DeleteVendor(ctx, id)
	if err != nil {
		log.Logger(ctx).Error("VendorService.DeleteVendor: Error raised by repository. Error: ", err)
		return nil, b.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	return resp, nil
}
