package cmd

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	repoErrInterceptor "github.com/iamgenii/error"
	authUtils2 "github.com/iamgenii/utils/auth_util"
	"github.com/iamgenii/utils/crypto_utils"
	"github.com/iamgenii/utils/http/request"
	"github.com/iamgenii/utils/http_utils"

	"github.com/iamgenii/utils"

	"github.com/iamgenii/validator"

	"github.com/iamgenii/configs"

	"github.com/iamgenii/email"

	"github.com/iamgenii/svc/middleware"

	"github.com/iamgenii/database"
	adminEndpoints "github.com/iamgenii/svc/admins/pkg/v1/endpoints"
	adminHandlers "github.com/iamgenii/svc/admins/pkg/v1/handlers"
	adminRepositories "github.com/iamgenii/svc/admins/pkg/v1/repositories"
	adminServices "github.com/iamgenii/svc/admins/pkg/v1/services"

	authEndpoints "github.com/iamgenii/svc/authorization/pkg/v1/endpoints"
	authHandlers "github.com/iamgenii/svc/authorization/pkg/v1/handlers"
	authRepositories "github.com/iamgenii/svc/authorization/pkg/v1/repositories"
	authServices "github.com/iamgenii/svc/authorization/pkg/v1/services"

	catEndpoints "github.com/iamgenii/svc/categaries/pkg/v1/endpoints"
	catHandlers "github.com/iamgenii/svc/categaries/pkg/v1/handlers"
	catRepositories "github.com/iamgenii/svc/categaries/pkg/v1/repositories"
	catServices "github.com/iamgenii/svc/categaries/pkg/v1/services"

	cityEndpoints "github.com/iamgenii/svc/cities/pkg/v1/endpoints"
	cityHandlers "github.com/iamgenii/svc/cities/pkg/v1/handlers"
	cityRepositories "github.com/iamgenii/svc/cities/pkg/v1/repositories"
	cityServices "github.com/iamgenii/svc/cities/pkg/v1/services"

	customerEndpoints "github.com/iamgenii/svc/customers/pkg/v1/endpoints"
	customerHandlers "github.com/iamgenii/svc/customers/pkg/v1/handlers"
	customerRepositories "github.com/iamgenii/svc/customers/pkg/v1/repositories"
	customerServices "github.com/iamgenii/svc/customers/pkg/v1/services"

	servicesEndpoints "github.com/iamgenii/svc/services/pkg/v1/endpoints"
	servicesHandlers "github.com/iamgenii/svc/services/pkg/v1/handlers"
	servicesRepositories "github.com/iamgenii/svc/services/pkg/v1/repositories"
	servicesServices "github.com/iamgenii/svc/services/pkg/v1/services"

	packagesEndpoints "github.com/iamgenii/svc/packages/pkg/v1/endpoints"
	packagesHandlers "github.com/iamgenii/svc/packages/pkg/v1/handlers"
	packagesRepositories "github.com/iamgenii/svc/packages/pkg/v1/repositories"
	packagesServices "github.com/iamgenii/svc/packages/pkg/v1/services"

	iamgesEndpoints "github.com/iamgenii/svc/images/pkg/v1/endpoints"
	imagesHandlers "github.com/iamgenii/svc/images/pkg/v1/handlers"

	venEndpoints "github.com/iamgenii/svc/vendors/pkg/v1/endpoints"
	venHandlers "github.com/iamgenii/svc/vendors/pkg/v1/handlers"
	venRepositories "github.com/iamgenii/svc/vendors/pkg/v1/repositories"
	venServices "github.com/iamgenii/svc/vendors/pkg/v1/services"

	probesEndpoints "github.com/iamgenii/svc/probes/pkg/v1/endpoints"
	probesHandlers "github.com/iamgenii/svc/probes/pkg/v1/handlers"
	probesRepositories "github.com/iamgenii/svc/probes/pkg/v1/repositories"
	probesServices "github.com/iamgenii/svc/probes/pkg/v1/services"

	_ "github.com/iamgenii/docs" // swagger docs
	httpSwagger "github.com/swaggo/http-swagger"

	"net/http"

	"github.com/gorilla/mux"
)

var (
	//HTTPPortFlag cli flag name for http port
	HTTPPortFlag = "http-port"

	//DBHostFlag cli flag name for database host
	DBHostFlag = "db-host"

	//DBUserFlag cli flag name for  database username
	DBUserFlag = "db-user"

	//DBPassFlag cli flag name for  database password
	DBPassFlag = "db-pass"

	//DBPortFlag cli flag name for  database port
	DBPortFlag = "db-port"

	//DBNameFlag cli flag name for  database name
	DBNameFlag = "db-name"

	//EmailAPIKeyFlag cli flag name for  database name
	EmailAPIKeyFlag = "email-api-key"

	//HTTPPortEnvVar  stores name of http port enviroment variable
	HTTPPortEnvVar = "HTTP_PORT"

	//DBHostEnvVar stores name of  database host enviroment variable
	DBHostEnvVar = "MYSQL_DB_HOST"

	//DBUserEnvVar stores name of  databse user enviroment variable
	DBUserEnvVar = "MYSQL_DB_USER"

	//DBPassEnvVar stores name of  databse password enviroment variable
	DBPassEnvVar = "MYSQL_DB_PASS"

	//DBAddrEnvVar stores name of  databse port enviroment variable
	DBAddrEnvVar = "MYSQL_DB_PORT"

	//DBNameEnvVar stores name of  databse name enviroment variable
	DBNameEnvVar = "MYSQL_DB_NAME"

	//EmailAPIKeyEnvVar Store evniroment variable name for api key
	EmailAPIKeyEnvVar = "EMAIL_API_KEY"
)

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("iamgenii", flag.ExitOnError)

// HTTPAddr http Port
var HTTPAddr = fs.String(HTTPPortFlag, "8080", "HTTP listen address defaults to 8080")

// DBHost db  hostname
var DBHost = fs.String(DBHostFlag, "", "Hostname for DB")

// DBUser db username
var DBUser = fs.String(DBUserFlag, "", "Username for DB")

// DBPass db password
var DBPass = fs.String(DBPassFlag, "", "Password for DB")

// DBAddr db port address
var DBAddr = fs.String(DBPortFlag, "", "Port Number for DB defaults to 3306")

// DBName db port address
var DBName = fs.String(DBNameFlag, "", "DB name")

// EmailAPIKey api key for email
var EmailAPIKey = fs.String(EmailAPIKeyFlag, "", "email api key name")

func init() {
	if err := fs.Parse(os.Args[1:]); err != nil {
		panic("Error in parsing command line parameters. ")
	}
}

// GetEnvironmentVariables reads an os environments variables
func GetEnvironmentVariables() {
	//fs.Parse(os.Args[1:])
	//get dbHost from environment variables
	var dbHost = os.Getenv(DBHostEnvVar)
	if len(dbHost) > 0 && (DBHost == nil || len(*DBHost) == 0) {
		DBHost = &dbHost
	}

	//get dbUser from environments variables
	var dbUser = os.Getenv(DBUserEnvVar)
	if len(dbUser) > 0 && (DBUser == nil || len(*DBUser) == 0) {
		DBUser = &dbUser
	}

	//get dbPass from environments variables
	var dbPass = os.Getenv(DBPassEnvVar)
	if len(dbPass) > 0 && (DBPass == nil || len(*DBPass) == 0) {
		DBPass = &dbPass
	}

	//get dbPort from environments variables
	var dbPort = os.Getenv(DBAddrEnvVar)
	if len(dbPort) > 0 && (DBAddr == nil || len(*DBAddr) == 0) {
		DBAddr = &dbPort
	}

	//get httpAddr from environments variables
	var httpAddr = os.Getenv(HTTPPortEnvVar)
	if len(httpAddr) > 0 && (HTTPAddr == nil || len(*HTTPAddr) == 0) {
		HTTPAddr = &httpAddr
	}

	//get dbName from environments variables
	var dbName = os.Getenv(DBNameEnvVar)
	if len(dbName) > 0 && (DBName == nil || len(*DBName) == 0) {
		DBName = &dbName
	}

	//get EmailAPIKey from environments variables
	var emailAPIKey = os.Getenv(EmailAPIKeyEnvVar)
	if len(emailAPIKey) > 0 && (EmailAPIKey == nil || len(*EmailAPIKey) == 0) {
		EmailAPIKey = &dbName
	}

}

// ValidateFlags checks the flags and update
func ValidateFlags() error {
	GetEnvironmentVariables()

	flagMessage := " is a required flag"
	if DBUser == nil || len(*DBUser) == 0 {
		return errors.New(DBUserFlag + flagMessage)
	}
	if DBPass == nil || len(*DBPass) == 0 {
		return errors.New(DBPassFlag + flagMessage)
	}
	if DBHost == nil || len(*DBHost) == 0 {
		return errors.New(DBHostFlag + flagMessage)
	}
	if DBName == nil || len(*DBName) == 0 {
		return errors.New(DBNameFlag + flagMessage)
	}
	if EmailAPIKey == nil || len(*EmailAPIKey) == 0 {
		return errors.New(EmailAPIKeyFlag + flagMessage)
	}

	return nil
}

// NewConnStr creates connection string for database
func NewConnStr() string {

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		*DBUser,
		*DBPass,
		*DBHost,
		*DBAddr,
		*DBName,
	)
	fmt.Println("database connection string : ", connStr)
	return connStr
}

// Run **
func Run() {

	// NewConfiguration loads configuration details
	configurations := configs.NewConfiguration("./configs/dev_config")

	if err := ValidateFlags(); err != nil {
		fmt.Println("Error While Validating flags: ", err)
		return
	}

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 60 * time.Second,
		}).DialContext,
	}
	httpClient := &http.Client{Transport: transport}
	request.NewHttpRequestBuilder(httpClient)

	connectionString := NewConnStr()
	connectionPool, err := database.NewDataStore(connectionString, configurations.DBConfig)
	if err != nil {
		return
	}
	router := mux.NewRouter()
	//api v1 routers
	routerAPIv1 := router.PathPrefix("/api/v1").Subrouter()
	//Email Dependency
	sendInBlue := email.NewSendInBlue(*EmailAPIKey)

	passwordPolicyConfiguration := configs.NewPasswordPolicyConfiguration(configurations.PasswordConfiguration)
	jwtConfig := configs.NewJwtConfig(configurations.JwtSecret)
	authConfig := configs.NewAuthConfig(configurations.AuthSecrets)

	passwordValidator := validator.NewPasswordValidator(passwordPolicyConfiguration)
	mobileNumberValidator := validator.NewMobileNumberValidator()
	requestValidator := validator.NewRequestValidator()

	hashedUtils := crypto_utils.NewHashUtils()
	httpReader := http_utils.NewHTTPReader(jwtConfig)
	httpWriter := http_utils.NewHTTPWriter()
	cookies := utils.NewCookies()

	cryptoUtils := crypto_utils.NewCryptoUtils(authConfig)
	jwtUtils := authUtils2.NewJwtUtils(jwtConfig, cryptoUtils)
	authUtils := authUtils2.NewAuthUtils(jwtUtils, cryptoUtils)

	repoErrorInterceptor := repoErrInterceptor.NewRepoErrorInterceptor()

	//User Services dependencies
	adminRepository := adminRepositories.NewAdminRepositoryImpl(connectionPool)
	adminService := adminServices.NewAdminServiceImpl(adminRepository, sendInBlue, cryptoUtils, repoErrorInterceptor, mobileNumberValidator, passwordValidator)
	adminHandler := adminHandlers.NewAdminHandler(adminService, httpReader, httpWriter, requestValidator)
	adminEndpoints.NewAdminRoute(routerAPIv1, adminHandler)

	//customer Services dependencies
	custRepository := customerRepositories.NewCustomerRepository(connectionPool)
	custService := customerServices.NewCustomerService(custRepository, sendInBlue, hashedUtils, repoErrorInterceptor)
	custHandler := customerHandlers.NewCustomerHandlers(custService, httpReader, httpWriter)
	customerEndpoints.NewCustomersRoute(routerAPIv1, custHandler)

	//Categories Services dependencies
	catRepository := catRepositories.NewCategoriesRepository(connectionPool)
	catService := catServices.NewCategoriesServiceImpl(catRepository, repoErrorInterceptor)
	catHandler := catHandlers.NewCategoriesHandlerImpl(catService, httpReader, httpWriter)
	catEndpoints.NewCategoriesRoutes(routerAPIv1, catHandler)

	//Forgot Password Services dependencies
	passRepository := authRepositories.NewForgotPasswordRepositories(connectionPool)
	passService := authServices.NewForgotPasswordService(passRepository, passwordValidator, hashedUtils, mobileNumberValidator)
	passHandler := authHandlers.NewForgotPasswordHandlers(passService, httpReader, httpWriter)
	authEndpoints.NewForgotPasswordRoutes(routerAPIv1, passHandler)

	//City master dependencies
	cityRepository := cityRepositories.NewCitiesRepository(connectionPool)
	cityService := cityServices.NewCitiesService(cityRepository, repoErrorInterceptor)
	cityHandler := cityHandlers.NewCitiesHandler(cityService, httpReader, httpWriter)
	cityEndpoints.NewCitiesRoutes(routerAPIv1, cityHandler)

	//services dependencies
	servicesRepos := servicesRepositories.NewServicesRepository(connectionPool)
	iamgeniiService := servicesServices.NewIamgeniiServices(servicesRepos, repoErrorInterceptor)
	servicesHandler := servicesHandlers.NewServicesHandlers(iamgeniiService, httpReader, httpWriter, requestValidator)
	servicesEndpoints.NewServicesRoutes(routerAPIv1, servicesHandler)

	//services categories dependencies
	servicesCategoriesRepos := servicesRepositories.NewServicesToCategoriesRepository(connectionPool)
	iamgeniiCategoriesService := servicesServices.NewIamgeniiCategoriesToServices(servicesCategoriesRepos)
	servicesCategoriesHandler := servicesHandlers.NewServicesCategoriesHandlers(iamgeniiCategoriesService, httpReader, httpWriter, requestValidator)
	servicesEndpoints.NewCategoriesToServicesRoutes(routerAPIv1, servicesCategoriesHandler)

	//packages dependencies
	packageRepositories := packagesRepositories.NewPackagesRepository(connectionPool)
	packageServices := packagesServices.NewPackagesServices(packageRepositories, repoErrorInterceptor)
	packageHandlers := packagesHandlers.NewPackagesHandlers(packageServices, httpReader, httpWriter, requestValidator)
	packagesEndpoints.NewPackageRoutes(routerAPIv1, packageHandlers)

	//packages service mapping dependencies
	packagesSvcMapRepositories := packagesRepositories.NewPackagesToServicesRepository(connectionPool)
	packagesSvcMapServices := packagesServices.NewPackagesServicesMappingService(packagesSvcMapRepositories, repoErrorInterceptor)
	packagesSvcMapHandlers := packagesHandlers.NewPackageServiceMappingHandler(packagesSvcMapServices, httpReader, httpWriter, requestValidator)
	packagesEndpoints.NewPackageServiceMappingRoutes(routerAPIv1, packagesSvcMapHandlers)

	imagesHandler := imagesHandlers.NewImageHandlers(httpReader, httpWriter, requestValidator)
	iamgesEndpoints.NewImageUploaderRoutes(routerAPIv1, imagesHandler)

	//customer Services dependencies
	vendorRepository := venRepositories.NewVendorRepository(connectionPool)
	vendorService := venServices.NewVendorService(vendorRepository, sendInBlue, hashedUtils, repoErrorInterceptor)
	vendorHandler := venHandlers.NewVendorHandlers(vendorService, httpReader, httpWriter, requestValidator)
	venEndpoints.NewVendorsRoute(routerAPIv1, vendorHandler)

	//Authorization Services dependencies
	authRepository := authRepositories.NewLoginRepository(connectionPool)
	authService := authServices.NewLoginService(authRepository, adminRepository, vendorRepository, repoErrorInterceptor, hashedUtils, jwtConfig, authUtils)
	authHandler := authHandlers.NewLoginHandler(authService, httpReader, httpWriter, cookies)
	prptectedUrlService := middleware.NewProtectedUrlService()
	authEndpoints.NewAuthorizationRoutes(routerAPIv1, authHandler)

	//probes dependencies
	probesRepos := probesRepositories.NewProbesRepository(connectionPool)
	probesSvc := probesServices.NewProbesService(probesRepos)
	probesHandler := probesHandlers.NewProbesHandlers(probesSvc, httpWriter)
	probesEndpoints.NewProbesRoutes(router, probesHandler)

	// Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	authMiddleware := middleware.NewAuthMiddleware(authUtils, jwtUtils, cryptoUtils, httpWriter, prptectedUrlService)

	routerAPIv1.Use(middleware.LoggingMiddleware)
	routerAPIv1.Use(middleware.CROSMiddleware)
	router.Use(authMiddleware.Middleware)

	fmt.Println("Starting HTTP server on :" + *HTTPAddr)
	fmt.Println(http.ListenAndServe(":"+*HTTPAddr, router))
}
