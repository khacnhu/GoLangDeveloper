package main

import (
	"fmt"
	"go-tutorial/controllers"
	internal "go-tutorial/internals/databases"
	"go-tutorial/services"

	"github.com/gin-gonic/gin"
)

type BOOK struct {
	Id     int    `json:"id" gorm:"PrimaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

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
	authController := &controllers.AuthController{}
	authController.InitAuthController()
	authController.InitAuthRoutes(router)

	router.Run(":3000")
}
