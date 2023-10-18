package entities

type User struct {
	Id         int `gorm:"primary_key, AUTO_INCREMENT"`
	Email      string
	Name       string
	Privileges []Privilege `gorm:"many2many:users_privileges"`
}

func (user User) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":         user.Id,
		"email":      user.Email,
		"name":       user.Name,
		"privileges": user.Privileges,
	}
}

func (user *User) TableName() string {
	return "users"
}
