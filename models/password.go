package models

//SendOTPReq holds varify mobile params
type SendOTPReq struct {
	ContactNumber string `json:"contact_number"`
	UserType      string `json:"user_type"`
}

//VerificationResp returns
type VerificationResp struct {
	//TODO :Remove OTP
	OTP      string `json:"otp"`
	JournyID string `json:"journy_id"`
}

//ValidateOtpReq holds varify mobile params
type ValidateOtpReq struct {
	OTP           string `gorm:"otp" json:"otp"`
	ContactNumber string `json:"contact_number"`
	JournyID      string `json:"journy_id"`
	UserType      string `json:"user_type"`
}

//ValidateOtpResp holds varify mobile params
type ValidateOtpResp struct {
	JournyID string `json:"journy_id"`
	Message  string `json:"message"`
}

//ResetPasswordReq hold reset password parameters
type ResetPasswordReq struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	Username    string `json:"username"`
	UserType    string `json:"user_type"`
}

//UpdatePasswordReq hold update password parameters
type UpdatePasswordReq struct {
	NewPassword   string `json:"new_password"`
	UserType      string `json:"user_type"`
	JournyID      string `json:"journy_id"`
	ContactNumber string `json:"contact_number"`
}
