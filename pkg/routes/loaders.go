package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graytonio/ffxivquesttracker/pkg/db"
	"github.com/sirupsen/logrus"
)

func GetFullQuestList() ([]db.Quest, error) {
	resp, err := http.DefaultClient.Get("https://www.garlandtools.org/db/doc/browse/en/2/quest.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var responseStruct struct {
		Browse []db.Quest `json:"browse"`
	}
	err = json.NewDecoder(resp.Body).Decode(&responseStruct)
	if err != nil {
		return nil, err
	}

	return responseStruct.Browse, nil
}

func (h *Handler) LoadQuests(c *gin.Context) {
	quests, err := GetFullQuestList()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	logrus.WithField("quests", len(quests)).Info("Loading quests into db")
	for i, q := range quests {
		q.Category = db.CategoryMap[q.GroupID]
		err = h.db.Assign(&q).FirstOrCreate(&q, db.Quest{ID: q.ID}).Error
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		logrus.WithFields(logrus.Fields{
			"count": i,
			"name": q.Name,
			"id": q.ID,
		}).Info("Loaded quest into db")
	}
	logrus.WithField("quests", len(quests)).Info("Loaded quests successfully")
	c.Status(200)
}