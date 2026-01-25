package error

import "errors"


//ErrSQLConn throws an error when sql sql connection not happened
var ErrSQLConn = errors.New("sql connection error")

//ErrSQLRecordExist throws an error when record already exist
var ErrSQLRecordExist = errors.New("record already exist")

//ErrRecordNotFound returns an error when no record founds
var ErrRecordNotFound = errors.New("record not found")

//ErrURLParamNotFound return an error when specified parameter not found in url path
var ErrURLParamNotFound = errors.New("url parameter not found")

//ErrURLQueryParamNotFound return an error when specified parameter not found in url path
var ErrURLQueryParamNotFound = errors.New("url query parameter not found")

//ErrPathVariableNotFound return an error when path variable not found
var ErrPathVariableNotFound = errors.New("url path variable not found")

//ErrQueryParamVariableNotInt return when path variable not an type of int
var ErrQueryParamVariableNotInt = errors.New("url query parameter value not a number")

//ErrPathVariableNotInt return when path variable not an type of int
var ErrPathVariableNotInt = errors.New("url path variable value  not a number")

//ErrInvalidLoginDetails returns an error when login details not match with existing details
var ErrInvalidLoginDetails = errors.New("invalid login details")

//ErrUpdateMobileNumber throw an error when no mobile number is updated by customer
var ErrUpdateMobileNumber = errors.New("update mobile number")

//ErrOTPExpires returns an error when otp expires
var ErrOTPExpires = errors.New("your otp is expired")

//ErrInvalidPassword when password is not password policy compliant
var ErrInvalidPassword = errors.New("invalid password")

//ErrInvalidMobileNumber when number is invalid
var ErrInvalidMobileNumber = errors.New("invalid mobile number")

//ErrNotAuthorizedToUpdatePassword when number is invalid
var ErrNotAuthorizedToUpdatePassword = errors.New("not authorized to update password ")

var ErrEmailAlreadyExist = errors.New("email already exist")

var ErrMobileNumberAlreadyExist = errors.New("mobile number already exist")

var ErrUsernameAlreadyExist = errors.New("username already exist")

var ErrInvalidUserType = errors.New("invalid user name")

//var ErrDuplicateEntry= errors.New("DUPLICATE ENTRY ERROR")