package controllers

import (
	"go-tutorial/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	noteService services.NoteServices
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine, noteService services.NoteServices) {
	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())
	notes.POST("/createNotes", n.CreateNotes())
	notes.GET("/status", n.GetNotesByStatus())
	notes.PUT("/", n.UpdateNotes())
	notes.DELETE("/:id", n.DeleteNotes())
	n.noteService = noteService
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		notes, err := n.noteService.GetNotesServices()

		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
		}

		ctx.JSON(200, gin.H{
			"LIST NOTES": notes,
		})

	}
}

func (n *NotesController) GetNotesByStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		status := ctx.Query("status")

		if status == "" {
			status = "false"
		}

		actualStatus, errStatus := strconv.ParseBool(status)

		if errStatus != nil {
			ctx.JSON(400, gin.H{
				"message": errStatus.Error(),
			})
		} else {
			notes, err := n.noteService.GetNotesByStatusServices(actualStatus)

			if err != nil {
				ctx.JSON(400, gin.H{
					"message": err.Error(),
				})
				return
			}

			ctx.JSON(200, gin.H{
				"LIST NOTES": notes,
			})
		}

	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {

	type NoteBody struct {
		Id     int    `json:"id"`
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status"`
	}

	return func(ctx *gin.Context) {

		var noteBody NoteBody

		if err := ctx.BindJSON(&noteBody); err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})

			return
		}
		note, err := n.noteService.CreateNotesService(noteBody.Id, noteBody.Title, noteBody.Status)

		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err,
			})
			return
		}

		ctx.JSON(200, gin.H{
			"note": note,
		})

	}
}

func (n *NotesController) UpdateNotes() gin.HandlerFunc {

	type NoteBody struct {
		Id     int    `json:"id" binding:"required"`
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status"`
	}

	return func(ctx *gin.Context) {

		var noteBody NoteBody

		if err := ctx.BindJSON(&noteBody); err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})

			return
		}
		note, err := n.noteService.UpdateNotesService(noteBody.Title, noteBody.Status, noteBody.Id)

		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err,
			})
			return
		}

		ctx.JSON(200, gin.H{
			"note": note,
		})

	}
}

func (n *NotesController) DeleteNotes() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Param("id")
		noteDeleteId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			ctx.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = n.noteService.DeleteNotesService(int(noteDeleteId))

		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"note": "Delet successfully",
		})

	}
}
