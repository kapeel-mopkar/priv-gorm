package entities

import "fmt"

type User struct {
	Id         int `gorm:"primary_key, AUTO_INCREMENT"`
	Email      string
	Name       string
	Privileges []Privilege `gorm:"many2many:users_privileges"`
}

func (user *User) TableName() string {
	return "users"
}

func (user User) ToString() string {
	return fmt.Sprintf("id: %d\nemail: %s\nname: %s", user.Id, user.Email, user.Name)
}
