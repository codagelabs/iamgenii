package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// type Config interface(
// )

//NewConfiguration Reads all Config details
func NewConfiguration(fileName string) *config {
	//TODO Not able to read Config should kill application
	config := config{}
	// Open our json configuration file
	fileName = fmt.Sprint(fileName, ".json")
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &config)
	return &config

}
