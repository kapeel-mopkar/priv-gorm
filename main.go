package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kapeel-mopkar/priv-gorm/models"
)

func main() {
	app := fiber.New()

	// Define your API routes
	app.Get("/users", ListPrivilegesForEachUser)
	app.Get("/privileges", ListUsersForEachScopedPrivilege)

	// Listen on a port
	app.Listen(":9093")
}

func ListPrivilegesForEachUser(c *fiber.Ctx) error {
	var userModel models.UserModel
	users, _ := userModel.FindAll()

	// Prepare JSON response
	response := make([]map[string]interface{}, 0)
	for _, user := range users {
		userData := user.ToJSON()
		response = append(response, userData)
	}

	return c.JSON(response)
}

func ListUsersForEachScopedPrivilege(c *fiber.Ctx) error {
	var privilegeModel models.PrivilegeModel
	privileges, _ := privilegeModel.FindAll()

	// Prepare JSON response
	response := make([]map[string]interface{}, 0)
	for _, privilege := range privileges {
		privilegeData := privilege.ToJSON()
		response = append(response, privilegeData)
	}

	return c.JSON(response)
}
