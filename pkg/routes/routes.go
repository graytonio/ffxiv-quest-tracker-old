package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/graytonio/ffxivquesttracker/pkg/db"
	"github.com/graytonio/ffxivquesttracker/pkg/templates"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type Handler struct {
	db           *gorm.DB
	discordOauth *oauth2.Config
	logger logrus.FieldLogger
}

type HandlerConfig struct {
	DiscordClientID string
	DiscordClientSecret string
	BaseDomain string
}

func NewHandler(db *gorm.DB, logger logrus.FieldLogger, config *HandlerConfig) *Handler {
	return &Handler{
		db: db,
		logger: logger,
		discordOauth: &oauth2.Config{
			ClientID:     config.DiscordClientID,
			ClientSecret: config.DiscordClientSecret,
			RedirectURL:  fmt.Sprintf("%s/auth/discord/callback", config.BaseDomain),
			Scopes:       []string{"identify", "email"},
			Endpoint: oauth2.Endpoint{
				AuthURL:   "https://discord.com/api/oauth2/authorize",
				TokenURL:  "https://discord.com/api/oauth2/token",
				AuthStyle: oauth2.AuthStyleInParams,
			},
		},
	}
}

func (h *Handler) IndexPage(c *gin.Context) {
	user, err := h.getUserFromSession(sessions.Default(c))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	quests, err := db.FetchQuestListings(h.db, user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	c.HTML(http.StatusOK, "", templates.Index(templates.CategorizeQuests(quests), user))
}

func (h *Handler) CompleteQuest(c *gin.Context) {
	user, err := h.getUserFromSession(sessions.Default(c))
	if err != nil {
	  c.AbortWithError(http.StatusInternalServerError, err)
	  return
	}

	questID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	quest := db.Quest{ID: questID}
	err = h.db.First(&quest, quest).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	quest.Completed = true

	if user != nil {
		err = h.db.Model(user).Association("CompletedQuests").Append(&quest)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	
	c.HTML(http.StatusOK, "", templates.QuestRow(quest))
}

func (h *Handler) UncompleteQuest(c *gin.Context) {
	user, err := h.getUserFromSession(sessions.Default(c))
	if err != nil {
	  c.AbortWithError(http.StatusInternalServerError, err)
	  return
	}

	questID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	quest := db.Quest{ID: questID}
	err = h.db.First(&quest, quest).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	quest.Completed = false

	if user != nil {
		err = h.db.Model(user).Association("CompletedQuests").Delete(&quest)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	
	c.HTML(http.StatusOK, "", templates.QuestRow(quest))
}