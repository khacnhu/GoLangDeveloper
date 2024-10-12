package services

import (
	"fmt"
	internal "go-tutorial/internals/models"

	"gorm.io/gorm"
)

type NoteServices struct {
	db *gorm.DB
}

type Note struct {
	Id   int
	Name string
}

func (n *NoteServices) InitService(database *gorm.DB) {
	n.db = database
	// n.db.AutoMigrate(&internal.Notes{})
}

func (n *NoteServices) GetNotesServices() []Note {
	data := []Note{
		{Id: 1, Name: "Note 1"},
		{Id: 2, Name: "Note 2"},
	}

	return data

}

func (n *NoteServices) CreateNotesService(title string, status bool) (*internal.Notes, error) {
	note := &internal.Notes{
		Title:  title,
		Status: status,
	}

	if err := n.db.Create(note).Error; err != nil {
		fmt.Println("create error ", err)
		return nil, err
	}

	// result := n.db.Create(&note)

	// if result.Error != nil {
	// 	fmt.Println("Record created failed: ", result.Error)
	// }

	fmt.Println("Record created successfully:", note)

	return note, nil
}
