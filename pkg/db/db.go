package db

import (
	"gorm.io/gorm"
)

func Init(connector gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(connector)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Quest{}, &Genre{})
}

func FetchUserQuests(db *gorm.DB, user *User) (quests []Quest, err error) {	
	tx := db.Model(&Quest{}).
		Joins("Genre").
		Select("quests.*, CASE WHEN user_quests.quest_id IS NOT NULL THEN true ELSE false END AS completed").
		Order("Genre.name ASC, sort_id ASC, id ASC")

	if user != nil {
		tx = tx.Joins("LEFT JOIN user_quests ON quests.id = user_quests.quest_id AND user_quests.user_id = ?", user.ID)
	}

	if tx.Find(&quests).Error != nil {
		return nil, err
	}

	return quests, nil
}

func QuestsInGenre(db *gorm.DB, genreID int) (quests []Quest, err error) {
	if db.Model(&Quest{}).Find(&quests, &Quest{GenreID: genreID}) != nil {
		return nil, err
	}

	return quests, nil
}

func QuestsInCategory(db *gorm.DB, category string) (quests []Quest, err error) {
	tx := db.Model(&Quest{}).Joins("Genre").Where("Genre.category = ?", category).Order("Genre.name ASC, sort_id ASC")
	if tx.Find(&quests).Error != nil {
		return nil, err
	}

	return quests, nil
}

func QuestsInSection(db *gorm.DB, section string) (quests []Quest, err error) {
	tx := db.Model(&Quest{}).Joins("Genre").Where("Genre.section = ?", section).Order("Genre.name ASC, sort_id ASC")

	if tx.Find(&quests).Error != nil {
		return nil, err
	}

	return quests, nil
}

func UpdateUserQuests(db *gorm.DB, complete bool, user *User, quests ...Quest) error {
	if complete {
		return CompleteUserQuest(db, user, quests...)
	}

	return UncompleteUserQuest(db, user, quests...)
}

func CompleteUserQuest(db *gorm.DB, user *User, quests ...Quest) error {
	return db.Model(user).Association("CompletedQuests").Append(quests)
}

func UncompleteUserQuest(db *gorm.DB, user *User, quests ...Quest) error {
	return db.Model(user).Association("CompletedQuests").Delete(quests)
}
