package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Checkin struct {
	DOB DOB `json:"dob"`
}

type DOB struct {
	time.Time
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

func main() {
	j := `{"dob":"1993/10/04 00:00:00"}`

	var c Checkin
	err := json.Unmarshal([]byte(j), &c)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c)
}
