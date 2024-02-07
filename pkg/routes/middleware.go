package routes

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/graytonio/ffxivquesttracker/pkg/db"
)

func (h *Handler) APIKeyAuth(c *gin.Context) {
	apiKey := c.GetHeader("Authorization")
	if apiKey == "" || !strings.HasPrefix(apiKey, "Bearer ") {
		c.Next()
		return
	}

	var user *db.User
	if err := h.db.Where(&db.User{APIKey: strings.TrimPrefix(apiKey, "Bearer ")}).First(user).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		c.AbortWithStatus(403)
		return
	}

	if err := h.setUserSession(sessions.Default(c), user); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Next()
}

func (h *Handler) Admin(c *gin.Context) {
	user, err := h.getUserFromSession(sessions.Default(c))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if !user.Admin {
		c.AbortWithStatus(403)
		return
	}

	c.Next()
}

func (h *Handler) ErrorHandler(c *gin.Context) {
	c.Next()

	// TODO Handle error states correctly

	for _, err := range c.Errors {
		h.logger.WithError(err).Error("error occured while handling route")
	}
}