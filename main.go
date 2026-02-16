package main

import (
	cmd "github.com/iamgenii/cmd/service"
)

// @title           IamGenii API
// @version         1.0
// @description     IamGenii service API documentation
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  support@iamgenii.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cmd.Run()
}
