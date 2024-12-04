package routes

import (
	"jasen-dev/jd-note/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/notes", controllers.GetNotes)
	r.GET("/notes/:id", controllers.GetNoteByID)
	r.POST("/notes", controllers.CreateNote)
	r.PUT("/notes/:id", controllers.UpdateNote)
	r.DELETE("/notes/:id", controllers.DeleteNote)
	r.PATCH("/notes/:id/archive", controllers.ArchiveNote)
	r.PATCH("/notes/:id/favorite", controllers.FavoriteNote)
}
