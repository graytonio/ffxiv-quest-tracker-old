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

func GetCategoryData() ([]db.Genre, error) {
	resp, err := http.DefaultClient.Get("https://www.garlandtools.org/db/doc/core/en/3/data.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var responseStruct struct {
		Genres map[string]db.Genre `json:"questGenreIndex"`
	}
	err = json.NewDecoder(resp.Body).Decode(&responseStruct)
	if err != nil {
		return nil, err
	}

	var results []db.Genre
	for _, g := range responseStruct.Genres {
		results = append(results, g)
	}

	return results, nil
}

func (h *Handler) LoadData(c *gin.Context) {
	genres, err := GetCategoryData()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	logrus.WithField("genres", len(genres)).Info("loading generes")
	for i, g := range genres {
		err = h.db.Assign(&g).FirstOrCreate(&g, db.Genre{ID: g.ID}).Error
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		logrus.WithFields(logrus.Fields{
			"count": i,
			"name":  g.Name,
			"id":    g.ID,
		}).Info("loaded genre")
	}
	logrus.WithField("genres", len(genres)).Info("finished loading generes successfully")

	quests, err := GetFullQuestList()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	logrus.WithField("quests", len(quests)).Info("Loading quests into db")
	for i, q := range quests {
		if q.GenreID == 0 {
			continue
		}

		err = h.db.Assign(&q).FirstOrCreate(&q, db.Quest{ID: q.ID}).Error
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		logrus.WithFields(logrus.Fields{
			"count": i,
			"name":  q.Name,
			"id":    q.ID,
		}).Info("Loaded quest into db")
	}
	logrus.WithField("quests", len(quests)).Info("Loaded quests successfully")
	c.Status(200)
}

// func (h *Handler) LoadQuests(c *gin.Context) {
// 	quests, err := GetFullQuestList()
// 	if err != nil {
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	logrus.WithField("quests", len(quests)).Info("Loading quests into db")
// 	for i, q := range quests {
// 		q.Category = db.CategoryMap[q.GroupID]
// 		err = h.db.Assign(&q).FirstOrCreate(&q, db.Quest{ID: q.ID}).Error
// 		if err != nil {
// 			c.AbortWithError(http.StatusInternalServerError, err)
// 			return
// 		}
// 		logrus.WithFields(logrus.Fields{
// 			"count": i,
// 			"name": q.Name,
// 			"id": q.ID,
// 		}).Info("Loaded quest into db")
// 	}
// 	logrus.WithField("quests", len(quests)).Info("Loaded quests successfully")
// 	c.Status(200)
// }
