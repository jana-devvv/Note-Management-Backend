package controllers

import (
	"jasen-dev/jd-note/models"
	"jasen-dev/jd-note/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {
	notes, err := services.GetAllNotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving notes"})
		return
	}
	c.JSON(http.StatusOK, notes)
}

func GetNoteByID(c *gin.Context) {
	id := c.Param("id")
	note, err := services.GetNoteByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Note not found"})
		return
	}
	c.JSON(http.StatusOK, note)
}

func CreateNote(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	createdNote, err := services.CreateNote(note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating note"})
		return
	}
	c.JSON(http.StatusCreated, createdNote)
}

func UpdateNote(c *gin.Context) {
	id := c.Param("id")
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	updatedNote, err := services.UpdateNote(id, note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating note"})
		return
	}
	c.JSON(http.StatusOK, updatedNote)
}

func DeleteNote(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteNote(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting note"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}

func ArchiveNote(c *gin.Context) {
	id := c.Param("id")
	statusQuery := c.DefaultQuery("status", "false")

	status, err := strconv.ParseBool(statusQuery)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid status value, expected 'true' or 'false'"})
		return
	}

	err = services.ArchiveNote(id, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error archiving note"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Note archived successfully"})
}

func FavoriteNote(c *gin.Context) {
	id := c.Param("id")
	statusQuery := c.DefaultQuery("status", "false")

	status, err := strconv.ParseBool(statusQuery)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid status value, expected 'true' or 'false'"})
		return
	}

	err = services.FavoriteNote(id, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error favoriting note"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Note favorited successfully"})
}
