package services

import (
	"context"
	"strings"

	"github.com/iamgenii/email"
	imgnErr "github.com/iamgenii/error"
	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
	repository "github.com/iamgenii/svc/admins/pkg/v1/repositories"
	"github.com/iamgenii/utils/crypto_utils"
	"github.com/iamgenii/validator"
	"github.com/dgrijalva/jwt-go"
)

// Claims describes Claims in token.
type Claims struct {
	ID        uint32 `json:"id"`
	AdminType string `json:"adminType"`
	jwt.StandardClaims
}

// AdminsService describes the service.
type AdminsService interface {
	CreateAdmin(ctx context.Context, createReq models.Admin) (interface{}, *imgnErr.IMGNError)
	GetAdminByID(context.Context, string) (*models.AdminResp, *imgnErr.IMGNError)
	GetAdmins(context.Context, models.GetAllAdminReq) (*models.GetAdminsResponse, *imgnErr.IMGNError)
	UpdateAdmin(context.Context, string, models.UpdateAdminRequest) (*string, *imgnErr.IMGNError)
	DeleteAdmin(context.Context, string) (interface{}, *imgnErr.IMGNError)
}

// AdminsServiceImpl having dependencies fro service
type AdminsServiceImpl struct {
	adminRepo          repository.AdminRepository
	sendBlue           *email.SendInBlue
	cryptoUtils        crypto_utils.CryptoUtils
	repoErrInterceptor imgnErr.RepoErrorInterceptor
	mobileValidator    validator.MobileNumberValidator
	passwordValidator  validator.PasswordValidator
}

// NewAdminServiceImpl inject dependencies admin repository
func NewAdminServiceImpl(adminRepo repository.AdminRepository,
	sendBlue *email.SendInBlue,
	cryptoUtils crypto_utils.CryptoUtils,
	interceptor imgnErr.RepoErrorInterceptor,
	mobileValidator validator.MobileNumberValidator,
	passwordValidator validator.PasswordValidator,
) AdminsService {
	return &AdminsServiceImpl{
		adminRepo:          adminRepo,
		sendBlue:           sendBlue,
		cryptoUtils:        cryptoUtils,
		repoErrInterceptor: interceptor,
		mobileValidator:    mobileValidator,
		passwordValidator:  passwordValidator,
	}
}

// CreateAdmin checks if email id the admin is exist in database or
// not by calling GetAdminByEmail function
// if exist return error Email address already in use
// otherwise call CreateAdmin function
func (b *AdminsServiceImpl) CreateAdmin(ctx context.Context,
	admin models.Admin) (interface{}, *imgnErr.IMGNError) {
	log.Logger(ctx).Debug("AdminsService.CreateAdmin Request: ", admin)

	if validationErr := b.passwordMobileValidator(ctx, admin); validationErr != nil {
		return nil, validationErr
	}

	_, err := b.adminRepo.CheckAdminExistOrNot(ctx, admin.Email, admin.Username, admin.Phone)
	if err == nil {
		log.Logger(ctx).Error("AdminsService.CreateAdmin: Error in create admin: ", admin)
		return admin, imgnErr.ErrorSQLRecordExist
	}

	encPass, encryptionErr := b.cryptoUtils.Encrypt(ctx, admin.Password)
	if encryptionErr != nil {
		log.Logger(ctx).Error("AdminsService.CreateAdmin: Error in encrypting password: ", err)
		return nil, encryptionErr
	}
	admin.Password = encPass
	resp, repoErr := b.adminRepo.CreateAdmin(ctx, admin)
	if repoErr != nil {
		log.Logger(ctx).Error("AdminsService.CreateAdmin: Error by repository: ", repoErr)
		return nil, b.repoErrInterceptor.ErrorMapper(ctx, repoErr)
	}

	return resp, nil
}

// GetAdminByID calls an GetAdminByID method of repository to retrieve records
func (b *AdminsServiceImpl) GetAdminByID(ctx context.Context,
	id string) (*models.AdminResp, *imgnErr.IMGNError) {
	log.Logger(ctx).Debug("AdminsService.GetAdminByID: Request id: ", id)

	resp, repoErr := b.adminRepo.GetAdminByID(ctx, id)
	if repoErr != nil {
		log.Logger(ctx).Error("AdminsService.GetAdminByID: Error by repository: ", repoErr)
		return nil, b.repoErrInterceptor.ErrorMapper(ctx, repoErr)
	}
	return resp, nil
}

// GetAdmins call GetAdmins function of repository to get all admins
func (b *AdminsServiceImpl) GetAdmins(ctx context.Context,
	req models.GetAllAdminReq) (*models.GetAdminsResponse, *imgnErr.IMGNError) {
	log.Logger(ctx).Debug("AdminsService.GetAdmins: In get admin request : ", req)

	resp, repoErr := b.adminRepo.GetAdmins(ctx, req)
	if repoErr != nil {
		log.Logger(ctx).Error("AdminsService.GetAdmins: Error by repository: ", repoErr)
		return nil, b.repoErrInterceptor.ErrorMapper(ctx, repoErr)
	}
	if len(resp.Admins) == 0 {
		log.Logger(ctx).Error("AdminsService.GetAdmins: Error by repository: Record not not found")
		return nil, imgnErr.ErrorRecordNotFound
	}
	return resp, nil
}

// UpdateAdmin **
func (b *AdminsServiceImpl) UpdateAdmin(ctx context.Context, id string,
	admin models.UpdateAdminRequest) (*string, *imgnErr.IMGNError) {
	log.Logger(ctx).Debugf("AdminsService.UpdateAdmin: In update admin request data  %v and updated by %s: ", admin, id)

	if strings.TrimSpace(admin.Phone) != "" {
		isValid := b.mobileValidator.IsValidMobileNumber(admin.Phone)
		if !isValid {
			log.Logger(ctx).Error("AdminsService.CreateAdmin: Error in mobile validation: ", admin)
			return nil, imgnErr.BadRequestErrorFunc("Invalid mobile number error")
		}
	}
	if err := b.adminRepo.UpdateAdmin(ctx, id, admin); err != nil {
		log.Logger(ctx).Error("AdminsService.UpdateAdmin: Error by repository: ", err)
		return nil, b.repoErrInterceptor.ErrorMapper(ctx, err)
	}
	resp := "Record Updated Successfully"
	return &resp, nil
}

// DeleteAdmin  send request to delete admin from db
func (b *AdminsServiceImpl) DeleteAdmin(ctx context.Context, id string) (interface{}, *imgnErr.IMGNError) {
	log.Logger(ctx).Debugf("AdminsService.DeleteAdmin: Delete admin request delete by %s: ", id)

	err := b.adminRepo.DeleteAdmin(ctx, id)
	if err != nil {
		log.Logger(ctx).Error("AdminsService.DeleteAdmin: Error by repository: ", err)
		return nil, b.repoErrInterceptor.ErrorMapper(ctx, err)
	}
	return "deletion successfully", nil
}

func (b *AdminsServiceImpl) passwordMobileValidator(ctx context.Context, admin models.Admin) *imgnErr.IMGNError {
	compliant, passCheckErr := b.passwordValidator.IsPassworPolicyCompliant(admin.Password)
	if !compliant || passCheckErr != nil {
		log.Logger(ctx).Error("AdminsService.CreateAdmin: Error in password validation: ", admin)
		return imgnErr.BadRequestErrorFunc(passCheckErr.Error())
	}

	isValid := b.mobileValidator.IsValidMobileNumber(admin.Phone)
	if !isValid {
		log.Logger(ctx).Error("AdminsService.CreateAdmin: Error in mobile validation: ", admin)
		return imgnErr.BadRequestErrorFunc("Invalid mobile number error")
	}
	return nil
}
