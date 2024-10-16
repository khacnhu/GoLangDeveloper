package services

import (
	"fmt"
	internal "go-tutorial/internals/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
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

	var users []*internal.User
	err := e.db.Find(&users).Error
	if err != nil {
		log.Fatal("error when get data of users")
	}

	usersCount := len(users) // Get the length of the users slice

	for i := 1; i < usersCount; i++ {
		var a = fmt.Sprintf("A%d", i)
		var b = fmt.Sprintf("B%d", i)
		var nameA = users[i].Email
		var nameB = users[i].Role
		f.SetCellValue(sheetName, a, nameA)
		f.SetCellValue(sheetName, b, nameB)

	}

	// Create a buffer to store the Excel file
	buf, err := f.WriteToBuffer()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating Excel file")
		return "failed"
	}

	// Serve the file as a downloadable response
	c.Header("Content-Disposition", "attachment; filename=users_file.xlsx")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", fmt.Sprintf("%d", len(buf.Bytes())))

	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf.Bytes())
	return "success"
}
