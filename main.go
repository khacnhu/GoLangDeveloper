package main

import (
	"fmt"
	"go-tutorial/controllers"
	internal "go-tutorial/internals/databases"
	"go-tutorial/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	db := internal.InitDB()

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

	if db == nil {
		fmt.Println("CONNECT DB FAILED HUHU")
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

	router.Run(fmt.Sprintf(":%s", port))
}
