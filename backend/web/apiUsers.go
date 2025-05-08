package web

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"rcjv-app/backend/database"
	"rcjv-app/backend/util"
)

func (a *API) getUsers(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("Invalid Token")
	}

	// Reformating the data, so I don't send bullshit to the frontend
	type user struct {
		ID       uint64 `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	var data []user

	var users []database.User
	err := a.PSQL.Find(&users).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error getting users: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting users")
	}

	for _, v := range users {
		var u user
		u.ID = v.ID
		u.Username = v.Username
		u.Email = v.Email
		data = append(data, u)
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func (a *API) createUser(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("Invalid Token")
	}

	// I have json tags on the structs for the database, so I format all my json the same and
	// can directly parse into the database structs and gorm can handle the modeling
	var user database.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid Body")
	}

	// Has Password
	user.Password = util.HashString(user.Password)

	err := a.PSQL.Create(&user).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Error creating user")
	}

	return c.Status(fiber.StatusOK).JSON("")
}

func (a *API) updateUser(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("Invalid Token")
	}

	// I first get the user, so I can update it later
	var user database.User
	err := a.PSQL.First(&user, c.Params("id")).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error getting users: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting users")
	}

	// Caching password
	var password string = user.Password

	// I have json tags on the structs for the database, so I format all my json the same and
	// can directly parse into the database structs and gorm can handle the modeling
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid Body")
	}

	// If you don't change the password in the frontend, it will be set to an empty string in the request
	// This piece of code ensures consistent passwords
	if user.Password == "" {
		user.Password = password
	} else {
		// Hash password if changed
		user.Password = util.HashString(user.Password)
	}

	err = a.PSQL.Model(&user).Updates(user).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Error updating user")
	}

	return c.Status(fiber.StatusOK).JSON("")
}

func (a *API) deleteUser(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("Invalid Token")
	}

	// The request has the id of the user, so after an auth check, the user directly gets soft deleted
	err := a.PSQL.Delete(&database.User{}, c.Params("id")).Error
	if err != nil {
		log.Printf("Error deleting user: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error deleting user")
	}

	return c.Status(fiber.StatusOK).JSON("")
}
