package web

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"rcjv-app/backend/database"
	"rcjv-app/backend/util"
)

// Returns the whole config to the admin panel
func (a *API) getConfig(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("unauthorized")
	}

	var config database.Config
	err := a.PSQL.Find(&config).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Database config not found, unexpected behaviour!!!")
			return c.Status(fiber.StatusInternalServerError).JSON("Database config not found, Unexpected behaviour!!!")
		}
		log.Printf("Error getting enabled config: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error fetching config")
	}

	return c.Status(fiber.StatusOK).JSON(config)
}

// Update the config
func (a *API) updateConfig(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("unauthorized")
	}

	var config database.Config
	err := a.PSQL.First(&config).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Database config not found, unexpected behaviour!!!")
			return c.Status(fiber.StatusInternalServerError).JSON("Database config not found, Unexpected behaviour!!!")
		}
		log.Printf("Error getting enabled config: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error parsing config")
	}

	if err := c.BodyParser(&config); err != nil {
		log.Printf("Error updating config: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON("Error parsing config")
	}

	err = a.PSQL.Model(&config).Save(config).Error
	if err != nil {
		log.Printf("Error updating config: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error updating config")
	}

	return c.Status(fiber.StatusOK).JSON("config updated successfully")
}

// Check if the app should be enabled
func (a *API) getEnabled(c *fiber.Ctx) error {
	var config database.Config
	err := a.PSQL.Find(&config).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Database config not found, unexpected behaviour!!!")
			return c.Status(fiber.StatusInternalServerError).JSON("Database config not found, Unexpected behaviour!!!")
		}
		log.Printf("Error getting enabled config: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error fetching config")
	}

	if config.AppEnabled {
		return c.Status(fiber.StatusOK).JSON("App is enabled")
	} else {
		return c.Status(fiber.StatusLocked).JSON("App is disabled")
	}
}

// Get the name of the current event
func (a *API) getName(c *fiber.Ctx) error {
	var config database.Config
	err := a.PSQL.Find(&config).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Database config not found, unexpected behaviour!!!")
			return c.Status(fiber.StatusInternalServerError).JSON("Database config not found, Unexpected behaviour!!!")
		}
		log.Printf("Error getting enabled config: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error fetching config")
	}

	return c.Status(fiber.StatusOK).JSON(config.EventName)
}
