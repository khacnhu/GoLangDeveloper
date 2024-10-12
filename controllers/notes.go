package controllers

import (
	"go-tutorial/services"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	noteService services.NoteServices
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine, noteService services.NoteServices) {
	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())
	notes.POST("/createNotes", n.CreateNotes())
	n.noteService = noteService
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"notes": n.noteService.GetNotesServices(),
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {

	type NoteBody struct {
		Title  string `json:"title"`
		Status bool   `json:"status"`
	}

	return func(ctx *gin.Context) {

		var noteBody NoteBody

		if err := ctx.BindJSON(noteBody); err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})

			return
		}
		n.noteService.CreateNotesService(noteBody.Title, noteBody.Status)

	}
}
