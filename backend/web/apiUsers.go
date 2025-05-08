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

	var users database.User
	err := a.PSQL.Find(&users).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error getting users: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting users")
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (a *API) createUser(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("Invalid Token")
	}

	var user database.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid Body")
	}

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

	var user database.User
	err := a.PSQL.First(&user, c.Params("id")).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error getting users: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting users")
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid Body")
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

	err := a.PSQL.Delete(&database.User{}, c.Params("id")).Error
	if err != nil {
		log.Printf("Error deleting user: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error deleting user")
	}

	return c.Status(fiber.StatusOK).JSON("")
}
