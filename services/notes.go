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

// Get All List Notes
func (n *NoteServices) GetNotesServices() ([]*internal.Notes, error) {

	var notes []*internal.Notes

	if err := n.db.Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil

}

func (n *NoteServices) GetNotesByStatusServices(status bool) ([]*internal.Notes, error) {

	var notes []*internal.Notes

	if err := n.db.Where("status = ?", status).Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil

}

// Create Notes
func (n *NoteServices) CreateNotesService(id int, title string, status bool) (*internal.Notes, error) {
	note := &internal.Notes{
		Id:     id,
		Title:  title,
		Status: status,
	}

	if err := n.db.Create(note).Error; err != nil {
		fmt.Println("create error ", err)
		return nil, err
	}

	fmt.Println("Record created successfully:", note)

	return note, nil
}
