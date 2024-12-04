package services

import (
	"errors"
	"jasen-dev/jd-note/config"
	"jasen-dev/jd-note/models"

	"gorm.io/gorm"
)

func GetAllNotes() ([]models.Note, error) {
	var notes []models.Note
	if err := config.DB.Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func GetNoteByID(id string) (*models.Note, error) {
	var note models.Note
	if err := config.DB.First(&note, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &note, nil
}

func CreateNote(note models.Note) (*models.Note, error) {
	if err := config.DB.Create(&note).Error; err != nil {
		return nil, err
	}
	return &note, nil
}

func UpdateNote(id string, note models.Note) (*models.Note, error) {
	var existNote models.Note
	if err := config.DB.First(&existNote, id).Error; err != nil {
		return nil, err
	}

	existNote.Title = note.Title
	existNote.Content = note.Content
	existNote.Archived = note.Archived
	existNote.Favorite = note.Favorite

	if err := config.DB.Save(&existNote).Error; err != nil {
		return nil, err
	}

	return &existNote, nil
}

func DeleteNote(id string) error {
	if err := config.DB.Delete(&models.Note{}, id).Error; err != nil {
		return err
	}
	return nil
}

func ArchiveNote(id string, status bool) error {
	return updateNoteStatus(id, "archived", status)
}

func FavoriteNote(id string, status bool) error {
	return updateNoteStatus(id, "favorite", status)
}

func updateNoteStatus(id, status string, value bool) error {
	var note models.Note
	if err := config.DB.First(&note, id).Error; err != nil {
		return err
	}

	switch status {
	case "archived":
		note.Archived = value
		note.Favorite = !value
	case "favorite":
		note.Favorite = value
		note.Archived = !value
	}

	if err := config.DB.Save(&note).Error; err != nil {
		return err
	}
	return nil
}
