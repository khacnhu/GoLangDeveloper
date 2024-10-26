package main

import (
	"fmt"
	"go-tutorial/configs"
	"go-tutorial/controllers"
	internal "go-tutorial/databases"
	"go-tutorial/services"
	"go-tutorial/utils"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment configuration
	configs.LoadConfig()
	gin.SetMode(configs.AppConfig.Server.Mode)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://trankhacnhu.com"}, // Allow specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},                     // Allow specific HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},          // Allow specific headers
		ExposeHeaders:    []string{"Content-Length"},                                   // Expose specific headers
		AllowCredentials: true,                                                         // Allow credentials (cookies, etc.)
		MaxAge:           12 * time.Hour,                                               // Preflight request cache duration
	}))
	router.Use(utils.Logger())

	db := internal.InitDB()
	// rdb := middlewares.InitRedis()
	// router.Use(middlewares.CacheMiddleware(rdb, 10*time.Minute))

	if db == nil {
		fmt.Println("CONNECT DB FAILED HUHU")
		panic("db server is error so I will stop everything")
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Reading environment variables
	port := os.Getenv("PORT")
	fmt.Println("check port = ", port)
	if port == "" {
		port = "8080" // default value
	}

	// NOTE COMPONENT
	notesServices := &services.NoteServices{}
	notesServices.InitService(db)
	notesController := &controllers.NotesController{}
	notesController.InitNotesController(*notesServices)
	notesController.InitRoutes(router)

	// AUTH COMPONENT
	authServices := &services.AuthService{}
	authServices.InitService(db)
	authController := &controllers.AuthController{}
	authController.InitAuthController(*authServices)
	authController.InitAuthRoutes(router)

	// EXPORT COMPONENT
	exportServices := &services.ExportService{}
	exportServices.InitService(db)
	exportController := &controllers.ExportController{}
	exportController.InitExportController(*exportServices)
	exportController.InitExportRoutes(router)

	router.Run(fmt.Sprintf(":%s", port))
}
