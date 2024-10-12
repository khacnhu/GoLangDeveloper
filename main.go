package main

import (
	"fmt"
	"go-tutorial/controllers"
	internal "go-tutorial/internals/databases"
	"go-tutorial/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := internal.InitDB()

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

	router.Run(":3000")
}
