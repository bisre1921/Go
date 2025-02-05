package controllers

import (
	"Note-Managment-System/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var notes []models.Note
var nextId = 1

func GetNotes(c *gin.Context) {
	c.JSON(http.StatusOK, notes)
}

func GetNoteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	for _, note := range notes {
		if note.Id == id {
			c.JSON(http.StatusOK, note)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
}

func CreateNote(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note"})
		return
	}
	note.Id = nextId
	nextId++
	notes = append(notes, note)
	c.JSON(http.StatusCreated, note)
}

func UpdateNoteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, note := range notes {
		if note.Id == id {
			if err := c.ShouldBindJSON(&notes[i]); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
				return
			}

			notes[i].Id = id // Keep the ID unchanged
			c.JSON(http.StatusOK, notes[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
}

func DeleteNoteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, note := range notes {
		if note.Id == id {
			notes = append(notes[:i], notes[i+1:]...) // Remove the note
			c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
}
