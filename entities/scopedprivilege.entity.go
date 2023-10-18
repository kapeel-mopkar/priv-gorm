package entities

type ScopedPrivilege struct {
	PrivilegeName string `gorm:"primary_key"`
	Users         []User `gorm:"many2many:user_scoped_privileges"`
}

func (scopedPrivilege ScopedPrivilege) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"name":  scopedPrivilege.PrivilegeName,
		"users": scopedPrivilege.Users,
	}
}

func (scopedPrivilege *ScopedPrivilege) TableName() string {
	return "scoped_privileges"
}
