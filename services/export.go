package services

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type ExportService struct {
	db *gorm.DB
}

func (e *ExportService) InitService(database *gorm.DB) {
	e.db = database
}

func (e *ExportService) ExportExcelService(c *gin.Context) string {
	// Create a new Excel file
	f := excelize.NewFile()

	// Create a sheet and write data
	sheetName := "Sheet1"
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
		return "failed"
	}

	// Serve the file as a downloadable response
	c.Header("Content-Disposition", "attachment; filename=users_file.xlsx")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", string(len(buf.Bytes())))

	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf.Bytes())
	return "success"
}
