package models

import (
	"github.com/kapeel-mopkar/priv-gorm/config"
	"github.com/kapeel-mopkar/priv-gorm/entities"
)

type PrivilegeModel struct {
}

func (privilegeModel PrivilegeModel) FindAll() ([]entities.Privilege, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var privileges []entities.Privilege
		db.Preload("Users").Find(&privileges)
		return privileges, nil
	}
}
