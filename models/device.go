package models

import "time"

//Device strores device details
type Device struct {
	DeviceID         uint
	DeviceType       string
	RegistrationDate time.Time
	UpdatetionDate   time.Time
}
