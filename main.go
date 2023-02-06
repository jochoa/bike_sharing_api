package main

import (
	"bike_sharing_api/controller"
	"bike_sharing_api/database"
	"bike_sharing_api/middleware"
    "bike_sharing_api/model"
	"bike_sharing_api/development"
	"os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
    docs "bike_sharing_api/docs"

)

func main() {
    loadEnv()
    loadDatabase()
	seedBicycleSamples()
	serveApplication()
}

func loadDatabase() {
    database.Connect()
	database.Database.AutoMigrate(&model.User{}, &model.Bicycle{}, &model.RentTransactions{} )

}

func seedBicycleSamples() {

	//check environament 
	environment := os.Getenv("ENVIRONMENT")

	if environment == "development" {
		log.Print("Seeding sample data")
		development.Seed()
	}

}

func loadEnv() {
    //err := godotenv.Load(".env.local")
    err := godotenv.Load(".env.local")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func serveApplication() {
    router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

    publicRoutes := router.Group("/auth")
    publicRoutes.POST("/register", controller.Register)
    publicRoutes.POST("/login", controller.Login)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/transaction", controller.AddTransaction)
	protectedRoutes.GET("/transaction", controller.GetAllTransactions)
	protectedRoutes.PATCH("/return_bicycle", controller.UpdateTransaction) //update rent transactions to reflect a bicycle was returned
	protectedRoutes.POST("/add_bicycle", controller.AddBicycle)
	protectedRoutes.GET("/get_all_bicycles", controller.GetAllBicycles)
	protectedRoutes.GET("/get_all_available_bicycles", controller.GetAllAvailableBicycles)

	//TODO: create custom log.logger
	//TODO: swagger documentation needs tunning 

    router.Run(":8000")
    log.Print("Server running on port 8000")

}