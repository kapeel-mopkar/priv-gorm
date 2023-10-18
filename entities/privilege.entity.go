package entities

type Privilege struct {
	Id   int `gorm:"primary_key, AUTO_INCREMENT"`
	Name string
	Users []User `gorm:"many2many:users_privileges"`
}

func (privilege *Privilege) TableName() string {
	return "privileges"
}

func (privilege Privilege) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":   privilege.Id,
		"name": privilege.Name,
		"users": privilege.Users,
	}
}

