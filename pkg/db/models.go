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

type Quest struct {
	ID                 int      `json:"i" gorm:"primarykey"`
	Name               string   `json:"n"`
	GroupID            int      `json:"g"`
	Category           Category `gorm:"embedded;embeddedPrefix:category_"`
	Location           string   `json:"l"`
	SortID             int      `json:"s"`
	UnlocksFunction    int      `json:"f"`
	PercentageComplete int
	Completed          bool
}