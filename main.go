package main

import (
	"fmt"
	"go-tutorial/controllers"
	internal "go-tutorial/databases"
	"go-tutorial/services"
	"go-tutorial/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
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
