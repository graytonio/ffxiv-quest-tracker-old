package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email           string
	DiscordID       string
	AvatarHash      string
	CharacterID     string
	CompletedQuests []Quest `gorm:"many2many:user_quests"`
}

type Genre struct {
	ID   int `json:"id" gorm:"primarykey"`
	Name string `json:"name" gorm:"unqueIndex"`
	Category string `json:"category" gorm:"index"`
	Section string `json:"section" gorm:"index"`
}

type Quest struct {
	ID                 int    `json:"i" gorm:"primarykey"`
	Name               string `json:"n"`
	GenreID            int    `json:"g"`
	Genre              Genre
	Location           string `json:"l"`
	SortID             int    `json:"s"`
	UnlocksFunction    int    `json:"f"`
	PercentageComplete int
	Completed          bool
}
