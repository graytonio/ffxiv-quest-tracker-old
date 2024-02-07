package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graytonio/ffxivquesttracker/pkg/db"
)

func (h *Handler) UpdateGenreDefiniton(c *gin.Context) {
	var inputGenre db.Genre
	err := c.ShouldBindJSON(&inputGenre)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = h.db.Assign(&inputGenre).Where(&db.Genre{ID: inputGenre.ID}).FirstOrCreate(&inputGenre).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusAccepted)
}

func (h *Handler) UpdateQuestDefinition(c *gin.Context) {
	var inputQuest db.Quest
	err := c.ShouldBindJSON(&inputQuest)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = h.db.Assign(&inputQuest).Where(&db.Quest{ID: inputQuest.ID}).FirstOrCreate(&inputQuest).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusAccepted)
}