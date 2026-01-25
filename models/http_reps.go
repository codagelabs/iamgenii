package models

//HTTPResp is used to send http api responce
type HTTPResp struct {
	Message string
	Status  int
	Data    interface{}
}

//HTTPErrResp is used to send http api responce
type HTTPErrResp struct {
	Message string
	Status  int
}
