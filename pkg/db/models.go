package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email           string
	DiscordID       string
	AvatarHash      string
	CharacterID     string
	APIKey          string
	Admin           bool
	CompletedQuests []Quest `gorm:"many2many:user_quests"`
}

type Genre struct {
	ID       int   `json:"id" gorm:"primarykey"`
	Name     string `json:"name" gorm:"unqueIndex"`
	Category string `json:"category" gorm:"index"`
	Section  string `json:"section" gorm:"index"`
}

type Quest struct {
	ID                 int   `json:"id" gorm:"primarykey"`
	Name               string `json:"name"`
	GenreID            uint   `json:"genre_id"`
	Genre              Genre
	Location           string `json:"location"`
	SortKey            int    `json:"sort_key"`
	PercentageComplete int

	// Shadow column filled out per request
	Completed bool
}
