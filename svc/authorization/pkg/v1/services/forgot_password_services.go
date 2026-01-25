package services

import (
	"context"
	"crypto/rand"
	"math/big"
	"strconv"

	"github.com/iamgenii/utils/crypto_utils"

	customError "github.com/iamgenii/error"
	"github.com/iamgenii/svc/authorization/pkg/v1/repositories"
	"github.com/iamgenii/validator"

	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
)

// ForgotPasswordService stores methods
type ForgotPasswordService interface {
	VerifyUserAndSendOTP(context.Context, models.SendOTPReq) (interface{}, error)
	ValidateOTP(context.Context, models.ValidateOtpReq) (interface{}, error)
	UpdatePassword(context.Context, models.UpdatePasswordReq) (interface{}, error)
}

// ForgotPasswordServiceImpl stores all password members
type forgotPasswordServiceImpl struct {
	passRepo           repositories.ForgotPasswordRepositories
	passValidator      validator.PasswordValidator
	passUtils          crypto_utils.HashUtils
	mobileNoValidatore validator.MobileNumberValidator
}

// NewForgotPasswordService inject dependencies
func NewForgotPasswordService(
	passRepo repositories.ForgotPasswordRepositories,
	passValidator validator.PasswordValidator,
	passUtils crypto_utils.HashUtils,
	mobileNoValidators validator.MobileNumberValidator,
) ForgotPasswordService {
	return forgotPasswordServiceImpl{
		passRepo:           passRepo,
		passValidator:      passValidator,
		passUtils:          passUtils,
		mobileNoValidatore: mobileNoValidators,
	}
}

// VerifyUserAndSendOTP *
func (b forgotPasswordServiceImpl) VerifyUserAndSendOTP(ctx context.Context,
	req models.SendOTPReq) (interface{}, error) {

	log.Logger(ctx).Info("VerifyMobileNumber req : ", req)

	//validate mobile Number
	if !b.mobileNoValidatore.IsValidMobileNumber(req.ContactNumber) {
		return nil, customError.ErrInvalidMobileNumber
	}

	otpDetails := models.MobileVarification{}

	if req.UserType == "admin" {
		adminUser, err := b.passRepo.VerifyAdminUser(ctx, req.ContactNumber)
		if err != nil {
			return nil, err
		}
		otpDetails.MobileNumber = adminUser.Phone
		otpDetails.UserID = adminUser.AdministratorsID
		otpDetails.UserType = req.UserType
	} else if req.UserType == "customer" {
		custUser, err := b.passRepo.VerifyCustomer(ctx, req.ContactNumber)
		if err != nil {
			return nil, err
		}
		otpDetails.MobileNumber = custUser.CustomersPhone
		otpDetails.UserID = uint64(custUser.CustomersID)
		otpDetails.UserType = req.UserType
	} else {
		return nil, customError.ErrInvalidUserType
	}

	//send OTP
	otp, err := b.sendOTP(ctx, otpDetails.MobileNumber)
	if err != nil {
		return nil, err
	}

	otpDetails.OTP = otp

	id, err := b.passRepo.SaveOTPDetails(ctx, otpDetails)
	if err != nil {
		return nil, err
	}

	journeyID := strconv.FormatUint(id, 10)

	resp := models.VerificationResp{}
	resp.OTP = otp
	resp.JournyID = journeyID

	return resp, nil
}

// SendOTP get customer register mobile number and send sms to register mobile number
func (b forgotPasswordServiceImpl) sendOTP(ctx context.Context,
	mobileNumber string) (string, error) {
	log.Logger(ctx).Info("SendOTP to number : ", mobileNumber)
	otp, err := b.generateOTP(ctx)
	if err != nil {
		return "", err
	}
	//TODO Send OTP Logic
	return otp, nil

}

// getRandNum returns a random number of size four
func (b forgotPasswordServiceImpl) generateOTP(ctx context.Context) (string, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(8999))
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}

// ValidateOTP *
func (b forgotPasswordServiceImpl) ValidateOTP(ctx context.Context,
	req models.ValidateOtpReq) (interface{}, error) {
	log.Logger(ctx).Info("Validate OTP Request : ", req)
	err := b.passRepo.ValidateOTP(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := models.ValidateOtpResp{
		JournyID: req.JournyID,
		Message:  "OTP Verification Success",
	}
	return resp, nil

}

func (b forgotPasswordServiceImpl) UpdatePassword(ctx context.Context,
	req models.UpdatePasswordReq) (interface{}, error) {
	log.Logger(ctx).Info("Update Password Req : ", req)
	isValid, err := b.passValidator.IsPassworPolicyCompliant(req.NewPassword)
	if err != nil {
		return nil, customError.ErrInvalidPassword
	}
	if !isValid {
		return nil, customError.ErrInvalidPassword
	}
	if err = b.passRepo.ValidateUpdateReq(ctx, req.ContactNumber, req.JournyID); err != nil {
		return nil, customError.ErrNotAuthorizedToUpdatePassword
	}
	hashPassword, err := b.passUtils.GenerateBcrtptHash(req.NewPassword)
	if err != nil {
		return nil, err
	}
	if req.UserType == "admin" {
		_, err := b.passRepo.UpdateAdminPassword(ctx, req.ContactNumber, hashPassword)
		if err != nil {
			return nil, err
		}
	} else if req.UserType == "customer" {
		_, err := b.passRepo.UpdateCustomerPassword(ctx, req.ContactNumber, hashPassword)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, customError.ErrInvalidUserType
	}
	return "password updated successfully", nil
}
