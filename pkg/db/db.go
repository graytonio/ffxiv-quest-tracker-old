package db

import "gorm.io/gorm"

func Init(connector gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(connector)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&User{}, &Quest{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func FetchQuestListings(db *gorm.DB, user *User) (quests []Quest, err error) {
	tx := db.Table("quests").Order("category_section ASC, category_category ASC")
	if user == nil {
		err = tx.Scan(&quests).Error
		return quests, err
	}

	err = tx.Select("quests.*, CASE WHEN user_quests.quest_id IS NOT NULL THEN true ELSE false END AS completed").
			Joins("LEFT JOIN user_quests ON quests.id = user_quests.quest_id AND user_quests.user_id = ?", user.ID).
			Scan(&quests).Error
	return quests, err
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
