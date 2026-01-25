package services

import (
	"context"
	"strconv"

	"github.com/iamgenii/configs"
	imgnErr "github.com/iamgenii/error"
	log "github.com/iamgenii/logs"
	adminRepository "github.com/iamgenii/svc/admins/pkg/v1/repositories"
	vendorRepository "github.com/iamgenii/svc/vendors/pkg/v1/repositories"
	authUtils "github.com/iamgenii/utils/auth_util"
	"github.com/iamgenii/utils/crypto_utils"

	"github.com/iamgenii/models"
	loginRepository "github.com/iamgenii/svc/authorization/pkg/v1/repositories"
)

type LoginService interface {
	LoginByEmailOrUsername(ctx context.Context, createReq models.LoginReq) (*models.CustLoginResp, *imgnErr.IMGNError)
}

type loginService struct {
	authRepo    loginRepository.LoginRepository
	adminRepo   adminRepository.AdminRepository
	vendorRepo  vendorRepository.VendorRepository
	jwtConfig   configs.JwtConfig
	authUtil    authUtils.AuthUtils
	repoErrInCp imgnErr.RepoErrorInterceptor
	hashUtils   crypto_utils.HashUtils
}

// NewLoginService inject dependencies categories login Repository
func NewLoginService(
	authRepo loginRepository.LoginRepository,
	adminRepo adminRepository.AdminRepository,
	vendorRepo vendorRepository.VendorRepository,
	repoErrorInterceptor imgnErr.RepoErrorInterceptor,
	hashUtils crypto_utils.HashUtils,
	jwtConfig configs.JwtConfig,
	authUtil authUtils.AuthUtils,
) LoginService {
	return &loginService{
		authRepo:    authRepo,
		adminRepo:   adminRepo,
		vendorRepo:  vendorRepo,
		jwtConfig:   jwtConfig,
		authUtil:    authUtil,
		repoErrInCp: repoErrorInterceptor,
		hashUtils:   hashUtils,
	}
}

// LoginByEmailOrUsername method
func (svc *loginService) LoginByEmailOrUsername(ctx context.Context, req models.LoginReq) (*models.CustLoginResp, *imgnErr.IMGNError) {
	var resp models.CustLoginResp
	log.Logger(ctx).Info("LoginService.LoginByEmailOrUsername: LoginByEmailOrUsername Request ", req)

	var password string
	idToken := authUtils.IdToken{}
	if req.UserType == "admin" {
		storeAdmin, err := svc.adminRepo.GetAdminByUsername(ctx, req.Username)
		if err != nil {
			log.Logger(ctx).Error("LoginService.LoginByEmailOrUsername: Error in get admin user. Error:  ", err)
			return nil, svc.repoErrInCp.ErrorMapper(ctx, err)
		}
		idToken = svc.adminIdTokenMapper(*storeAdmin)
		password = storeAdmin.Password

	} else if req.UserType == "customer" {
		storedCustomer, err := svc.authRepo.CustomerLogin(ctx, req)
		if err != nil {
			log.Logger(ctx).Error(err)
			return nil, svc.repoErrInCp.ErrorMapper(ctx, err)

		}
		idToken = svc.customerIdTokenMapper(*storedCustomer)
		password = storedCustomer.Password
	} else if req.UserType == "vendor" {

		vendor, err := svc.authRepo.VendorLogin(ctx, req)
		if err != nil {
			log.Logger(ctx).Error(err)
			return nil, svc.repoErrInCp.ErrorMapper(ctx, err)

		}
		idToken = svc.vendorIdTokenMapper(*vendor)
		password = vendor.Password
	} else {

		return nil, imgnErr.ErrorInvalidLoginDetails
	}
	//err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password))
	//if err != nil {
	//	log.Logger(ctx).Error(err)
	//	return nil, imgnErr.ErrorInvalidLoginDetails
	//}
	err := svc.hashUtils.MatchBcryptedHash(password, req.Password)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, imgnErr.ErrorInvalidLoginDetails
	}
	encryptIdToken, imgnError := svc.authUtil.EncryptIdToken(ctx, idToken)
	if imgnError != nil {
		log.Logger(ctx).Error(err)
		return nil, imgnError
	}

	resp.Token = encryptIdToken
	resp.UserType = req.UserType
	return &resp, nil
}
func (svc *loginService) adminIdTokenMapper(admin models.Admin) authUtils.IdToken {
	adminId := strconv.FormatInt(int64(admin.AdministratorsID), 10)
	return authUtils.IdToken{
		UserType:     []authUtils.UserType{authUtils.ADMIN_USER},
		AppType:      []authUtils.AppType{authUtils.ADMIN_APP},
		Email:        admin.Email,
		FirstName:    admin.FirstName,
		LastName:     admin.LastName,
		LoginMode:    authUtils.LOGIN_MODE_WEB,
		MobileNumber: admin.Phone,
		UserId:       adminId,
	}

}
func (svc *loginService) customerIdTokenMapper(customer models.Customer) authUtils.IdToken {
	customersID := strconv.Itoa(int(customer.CustomersID))
	return authUtils.IdToken{
		UserType:     []authUtils.UserType{authUtils.ADMIN_USER},
		AppType:      []authUtils.AppType{authUtils.ADMIN_APP},
		Email:        customer.Email,
		FirstName:    customer.CustomersFirstName,
		LastName:     customer.CustomersLastName,
		LoginMode:    authUtils.LOGIN_MODE_WEB,
		MobileNumber: customer.CustomersPhone,
		UserId:       customersID,
	}

}

func (svc *loginService) vendorIdTokenMapper(vendor models.Vendor) authUtils.IdToken {
	customersID := strconv.FormatInt(int64(vendor.VendorId), 10)
	return authUtils.IdToken{
		UserType:     []authUtils.UserType{authUtils.ADMIN_USER},
		AppType:      []authUtils.AppType{authUtils.ADMIN_APP},
		Email:        vendor.Email,
		FirstName:    vendor.Name,
		LastName:     vendor.Name,
		LoginMode:    authUtils.LOGIN_MODE_WEB,
		MobileNumber: vendor.MobileNumber,
		UserId:       customersID,
	}

}
