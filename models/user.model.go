package models

import (
	"github.com/kapeel-mopkar/priv-gorm/config"
	"github.com/kapeel-mopkar/priv-gorm/entities"
)

type UserModel struct {
}

func (userModel UserModel) FindAll() ([]entities.User, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var users []entities.User
		db.Preload("Privileges").Find(&users)
		return users, nil
	}
}
