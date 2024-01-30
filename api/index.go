package handler

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/graytonio/ffxivquesttracker/pkg/db"
	"github.com/graytonio/ffxivquesttracker/pkg/templates"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func GetUserFromSession(r *http.Request, dbSession *gorm.DB, store *sessions.CookieStore) (*db.User, error) {
	session, _ := store.Get(r, "usersession")

	userID, ok := session.Values["user_id"]
	if !ok {
		return nil, nil
	}

	var user db.User
	if err := dbSession.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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

	quests, err := db.FetchQuestListings(dbSession, user)
	if err != nil {
		logrus.WithError(err).Error()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templates.Index(templates.CategorizeQuests(quests), user).Render(r.Context(), w)
}