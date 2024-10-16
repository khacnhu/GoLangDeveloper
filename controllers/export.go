package controllers

import (
	"go-tutorial/services"

	"github.com/gin-gonic/gin"
)

type ExportController struct {
	exportService services.ExportService
}

func (e *ExportController) InitExportController(exportService services.ExportService) *ExportController {
	e.exportService = exportService
	return e
}

func (e *ExportController) InitExportRoutes(router *gin.Engine) {
	exportRouter := router.Group("/")

	exportRouter.GET("/export-excel", e.ExportExcel()) // you should add middleware for role ADMIN **
	exportRouter.GET("/pdf", e.ExportPdf())
}

func (e *ExportController) ExportExcel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		exportData := e.exportService.ExportExcelService(ctx)

		if exportData == "failed" {
			ctx.JSON(400, gin.H{
				"error": "export excel failed sorry",
			})
			return

		}

		ctx.JSON(200, gin.H{
			"message": "export file excel for admin successfully",
		})

	}
}

func (e *ExportController) ExportPdf() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		exportData := e.exportService.ExportPdfService(ctx)

		if exportData == "FAILED" {
			ctx.JSON(400, gin.H{
				"error": "export pdf failed sorry",
			})
			return

		}

		ctx.JSON(200, gin.H{
			"message": exportData,
		})

	}
}
