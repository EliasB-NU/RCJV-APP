package web

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"rcjv-app/backend/database"
	"rcjv-app/backend/util"
	"strconv"
)

func (a *API) getFields(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{})
	}

	var fields []database.Field
	err := a.PSQL.Find(&fields).Error
	if err != nil {
		log.Printf("Error retrieving fields: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting fields")
	}

	return c.Status(fiber.StatusOK).JSON(fields)
}

func (a *API) createField(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{})
	}

	// Parse into database struct and create new field
	var field database.Field
	if err := c.BodyParser(&field); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}

	err := a.PSQL.Model(&database.Field{}).Create(&field).Error
	if err != nil {
		log.Printf("Error creating field: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}

	return c.Status(fiber.StatusOK).JSON("Field created")
}

func (a *API) updateField(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{})
	}

	// Get user, parse into user, update user
	var field database.Field
	err := a.PSQL.First(&field, c.Params("id")).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{})
		}
		log.Printf("Error fetching field: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting field")
	}

	if err := c.BodyParser(&field); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid field")
	}

	err = a.PSQL.Model(&field).Updates(field).Error
	if err != nil {
		log.Printf("Error updating field: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error updating field")
	}

	return c.Status(fiber.StatusOK).JSON("")
}

func (a *API) deleteField(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{})
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON("invalid id")
	}

	err = a.PSQL.Delete(&database.Field{}, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON("field not found")
		}
		log.Printf("Error deleting field: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error deleting field")
	}

	return c.Status(fiber.StatusOK).JSON("Deleted field")
}
