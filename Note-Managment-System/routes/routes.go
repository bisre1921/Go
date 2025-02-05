package routes

import (
	"Note-Managment-System/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRoutes(router *gin.Engine) {
	router.GET("/notes", controllers.GetNotes)
	router.GET("/notes/:id", controllers.GetNoteByID)
	router.POST("/notes", controllers.CreateNote)
	router.PUT("/notes/:id", controllers.UpdateNoteByID)
	router.DELETE("/notes/:id", controllers.DeleteNoteByID)
}
