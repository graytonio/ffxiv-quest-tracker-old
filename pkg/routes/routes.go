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
	logger       logrus.FieldLogger
}

type HandlerConfig struct {
	DiscordClientID     string
	DiscordClientSecret string
	BaseDomain          string
}

func NewHandler(db *gorm.DB, logger logrus.FieldLogger, config *HandlerConfig) *Handler {
	return &Handler{
		db:     db,
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

	quests, err := db.FetchUserQuests(h.db, user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "", templates.Index(templates.CategorizeQuests(quests), user))
}

func (h *Handler) UpdateQuest(complete bool) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		err = h.db.Joins("Genre").First(&quest, quest).Error
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		quest.Completed = complete

		if user != nil {
			err = db.UpdateUserQuests(h.db, complete, user, quest)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}

		c.HTML(http.StatusOK, "", templates.QuestRow(quest))
	}
}

func (h *Handler) UpdateSection(complete bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := h.getUserFromSession(sessions.Default(c))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var requestData struct {
			Key string `json:"key"`
		}
		c.BindJSON(&requestData)
		section := requestData.Key

		quests, err := db.QuestsInSection(h.db, section)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		logrus.WithField("section", section).WithField("quests", len(quests)).Info("Completing Quests")
		for i := range quests {
			quests[i].Completed = complete
		}

		if user != nil {
			err = db.UpdateUserQuests(h.db, complete, user, quests...)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}

		categorized := templates.CategorizeQuests(quests)
		c.HTML(http.StatusOK, "", templates.QuestSection(section, categorized[section]))
	}

}

func (h *Handler) UpdateCategory(complete bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := h.getUserFromSession(sessions.Default(c))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var requestData struct {
			Key string `json:"key"`
		}
		c.BindJSON(&requestData)
		category := requestData.Key

		quests, err := db.QuestsInCategory(h.db, category)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		logrus.WithField("section", category).WithField("quests", len(quests)).Info("Completing Quests")
		for i := range quests {
			quests[i].Completed = complete
		}

		if user != nil {
			err = db.UpdateUserQuests(h.db, complete, user, quests...)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}

		c.HTML(http.StatusOK, "", templates.QuestCategory(category, quests))
	}
}
