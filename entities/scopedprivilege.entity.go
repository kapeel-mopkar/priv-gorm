package entities

import "fmt"

type ScopedPrivilege struct {
	PrivilegeName string `gorm:"primary_key"`
	Users         []User `gorm:"many2many:user_scoped_privileges"`
}

func (scopedPrivilege *ScopedPrivilege) TableName() string {
	return "scoped_privileges"
}

func (scopedPrivilege ScopedPrivilege) ToString() string {
	return fmt.Sprintf("name: %s", scopedPrivilege.PrivilegeName)
}
