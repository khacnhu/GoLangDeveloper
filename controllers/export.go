package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type ExportController struct {
}

func (a *ExportController) InitExportController() *ExportController {
	return a
}

func (a *ExportController) InitExportRoutes(router *gin.Engine) {
	authRouter := router.Group("/")

	authRouter.GET("/export-excel", ExportExcel())
	authRouter.GET("/csv")
	authRouter.GET("/pdf")
}

func ExportExcel() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a new Excel file
		f := excelize.NewFile()

		// Create a sheet and write data
		sheetName := "Sheet"
		f.SetSheetName(f.GetSheetName(1), sheetName)
		f.SetCellValue(sheetName, "A1", "Name")
		f.SetCellValue(sheetName, "B1", "Age")
		f.SetCellValue(sheetName, "A2", "Alice")
		f.SetCellValue(sheetName, "B2", 30)
		f.SetCellValue(sheetName, "A3", "Bob")
		f.SetCellValue(sheetName, "B3", 25)

		// Create a buffer to store the Excel file
		buf, err := f.WriteToBuffer()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error creating Excel file")
			return
		}

		// Serve the file as a downloadable response
		c.Header("Content-Disposition", "attachment; filename=users_file.xlsx")
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Length", string(len(buf.Bytes())))

		c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf.Bytes())
	}
}
