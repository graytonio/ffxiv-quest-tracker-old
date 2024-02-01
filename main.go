package main

import (
	"os"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/graytonio/ffxivquesttracker/pkg/db"
	"github.com/graytonio/ffxivquesttracker/pkg/routes"
	"github.com/graytonio/ffxivquesttracker/pkg/templates"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"gorm.io/driver/mysql"
)

func main() {
	dbConn, err := db.Init(mysql.Open(os.Getenv("DB_DSN")))
	if err != nil {
		panic(err)
	}

	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	h := routes.NewHandler(
		dbConn, 
		log,
		&routes.HandlerConfig{
			DiscordClientID: os.Getenv("DISCORD_CLIENT_ID"),
			DiscordClientSecret: os.Getenv("DISCORD_CLIENT_SECRET"),
			BaseDomain: os.Getenv("DOMAIN"),
		},
	)

	store := gormsessions.NewStore(dbConn, true, []byte("secret"))

	r := gin.New()
	r.HTMLRender = &templates.TemplRender{}
	r.Use(ginlogrus.Logger(log), gin.Recovery())
	r.Use(sessions.Sessions("usersessions", store))
	r.Static("/assets", "assets/dist")
	
	r.GET("/load", h.LoadQuests)

	r.GET("/", h.IndexPage)

	r.GET("/auth/discord", h.DiscordAuth)
	r.GET("/auth/discord/callback", h.DiscordAuthCallback)
	r.GET("/logout", h.Logout)

	r.POST("/:id", h.CompleteQuest)
	r.DELETE("/:id", h.UncompleteQuest)

	r.Run()
}