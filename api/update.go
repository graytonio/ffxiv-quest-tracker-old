package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/graytonio/ffxivquesttracker/pkg/db"
	"github.com/graytonio/ffxivquesttracker/pkg/templates"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	dbSession, err := gorm.Open(mysql.Open(os.Getenv("DB_DSN")))
	if err != nil {
		logrus.WithError(err).Error()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := GetUserFromSession(r, dbSession, store)
	if err != nil {
		logrus.WithError(err).Error()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	questID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		logrus.WithError(err).Error()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	quest := db.Quest{ID: questID}
	err = dbSession.First(&quest, quest).Error
	if err != nil {
		logrus.WithError(err).Error()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	quest.Completed = r.Method == http.MethodPost

	if user == nil {
		templates.QuestRow(quest).Render(r.Context(), w)
		return
	}

	if r.Method == http.MethodPost {
		err = dbSession.Model(user).Association("CompletedQuests").Append(&quest)
	} else {
		err = dbSession.Model(user).Association("CompletedQuests").Delete(&quest)
	}
	if err != nil {
		logrus.WithError(err).Error()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templates.QuestRow(quest).Render(r.Context(), w)
}