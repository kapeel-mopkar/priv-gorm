package entities

import "fmt"

type Privilege struct {
	Id    int `gorm:"primary_key, AUTO_INCREMENT"`
	Name  string
	Users []User `gorm:"many2many:users_privileges"`
}

func (privilege *Privilege) TableName() string {
	return "privileges"
}

func (privilege Privilege) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s", privilege.Id, privilege.Name)
}
