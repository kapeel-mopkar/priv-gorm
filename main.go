package main

import (
	"fmt"

	"github.com/kapeel-mopkar/priv-gorm/models"
)

func main() {
	fmt.Println("List the privileges for each user")
	ListPrivilegesForEachUser()

	fmt.Println("List the users for each privilege")
	ListUsersForEachScopedPrivilege()
}

func ListPrivilegesForEachUser() {
	var userModel models.UserModel
	users, _ := userModel.FindAll()
	for _, user := range users {
		fmt.Println(user.ToString())
		fmt.Println("Privileges: ", len(user.Privileges))
		if len(user.Privileges) > 0 {
			for _, privilege := range user.Privileges {
				fmt.Println(privilege.ToString())
				fmt.Println("========================")
			}
		}
		fmt.Println("-----------------------------")
	}
}

func ListUsersForEachScopedPrivilege() {
	var privilegeModel models.PrivilegeModel
	privileges, _ := privilegeModel.FindAll()
	for _, privilege := range privileges {
		fmt.Println(privilege.ToString())
		fmt.Println("Users: ", len(privilege.Users))
		if len(privilege.Users) > 0 {
			for _, user := range privilege.Users {
				fmt.Println(user.ToString())
				fmt.Println("========================")
			}
		}
		fmt.Println("-----------------------------")
	}
}
