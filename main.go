package main

import (
	"log"

	"ecommerce/config"
	ongkirHandler "ecommerce/ongkir/handler"
	ongkirRepo "ecommerce/ongkir/repo"
	ongkirUsecase "ecommerce/ongkir/usecase"
	productHandler "ecommerce/product/handler"
	productRepo "ecommerce/product/repo"
	productUsecase "ecommerce/product/usecase"
	transactionHandler "ecommerce/transaction/handler"
	transactionRepo "ecommerce/transaction/repo"
	transactionUsecase "ecommerce/transaction/usecase"
	userHandler "ecommerce/user/handler"
	userRepo "ecommerce/user/repo"
	userUsecase "ecommerce/user/usecase"
	paymentRepo "ecommerce/payment/repo"
	paymentUsecase "ecommerce/payment/usecase"
	paymentHandler "ecommerce/payment/handler"
	"ecommerce/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @name Authorization
// @in header

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

func init() {
	viper.SetConfigFile("config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	DB := config.DbConnect()
	defer DB.Close()

	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This testing request API Ecommerce Indonesia"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8679"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()

	ongkirRepo := ongkirRepo.CreateOngkirRepo(DB)
	userRepo := userRepo.CreateUserRepo(DB)
	productRepo := productRepo.CreateProducRepo(DB)
	paymentRepo := paymentRepo.CreatePaymetRepo(DB)
	transactionRepo := transactionRepo.CreateTransactionRepo(DB)
	
	ongkirUsecase := ongkirUsecase.CreateOngkirUsecase(ongkirRepo)
	userUsecase := userUsecase.CreateUserUsecase(userRepo, ongkirRepo, transactionRepo)
	productUsecase := productUsecase.CreateProductUsecase(productRepo, userRepo, ongkirRepo, transactionRepo)
	paymentUsecase := paymentUsecase.CreatePaymentUsecase(paymentRepo)
	transactionUsecase := transactionUsecase.CreateTransactionUsecase(transactionRepo, productRepo, userRepo, ongkirRepo)

	ongkirHandler.CreateOngkirHandler(router, ongkirUsecase, userUsecase)
	userHandler.CreateUserHandler(router, userUsecase)
	productHandler.CreateProductHandler(router, productUsecase, userUsecase)
	paymentHandler.CreatePaymentHandler(router, paymentUsecase, userUsecase)
	transactionHandler.CreateTransactionHandler(router, transactionUsecase, userUsecase, productUsecase, ongkirUsecase)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := router.Run(":" + viper.GetString("port.port"))
	if err != nil {
		log.Fatal(err)
		return
	}
}