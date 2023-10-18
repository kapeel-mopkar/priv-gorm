package models

import (
	"github.com/kapeel-mopkar/priv-gorm/config"
	"github.com/kapeel-mopkar/priv-gorm/entities"
)

type ScopedPrivilegeModel struct {
}

func (scopedPrivilegeModel ScopedPrivilegeModel) FindAll() ([]entities.ScopedPrivilege, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var scopedPrivileges []entities.ScopedPrivilege
		db.Preload("Users").Find(&scopedPrivileges)
		return scopedPrivileges, nil
	}
}
