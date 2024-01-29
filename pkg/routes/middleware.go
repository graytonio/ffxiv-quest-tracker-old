package routes

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) ErrorHandler(c *gin.Context) {
	c.Next()

	// TODO Handle error states correctly

	for _, err := range c.Errors {
		h.logger.WithError(err).Error("error occured while handling route")
	}
}