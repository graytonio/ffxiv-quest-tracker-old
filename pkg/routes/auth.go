package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/graytonio/ffxivquesttracker/pkg/db"
)

func (h *Handler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user_id")
	session.Save()
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (h *Handler) setUserSession(s sessions.Session, user *db.User) error {
	s.Set("user_id", user.ID)
	return s.Save()
}

func (h *Handler) getUserFromSession(s sessions.Session) (*db.User, error) {
	userID := s.Get("user_id")
	if userID == nil {
		return nil, nil
	}

	h.logger.WithField("user", userID).Info("looking up user")
	var user db.User
	err := h.db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	h.logger.Info("found user")

	return &user, nil
}

func (h *Handler) DiscordAuth(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, h.discordOauth.AuthCodeURL("state"))
}

func (h *Handler) DiscordAuthCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := h.discordOauth.Exchange(c, code)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	resp, err := h.discordOauth.Client(c, token).Get("https://discord.com/api/users/@me")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	type discordUserData struct {
		Id    string `json:"id"`
		Email string `json:"email"`
		Avatar string `json:"avatar"`
	}

	// Decode user information
	var discordUser discordUserData
	if err := json.NewDecoder(resp.Body).Decode(&discordUser); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	user := db.User{DiscordID: discordUser.Id, Email: discordUser.Email, AvatarHash: discordUser.Avatar}
	err = h.db.Assign(user).FirstOrCreate(&user, db.User{DiscordID: user.DiscordID}).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = h.setUserSession(sessions.Default(c), &user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}
