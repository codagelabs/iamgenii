package models

import (
	"strings"
	"time"
)

type DOB struct {
	time.Time `gorm:"column:customers_dob"`
}

func (dob *DOB) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)
	newTime, err := time.Parse("2006/01/02 15:04:05", strInput)
	if err != nil {
		return err
	}

	dob.Time = newTime
	return nil
}
