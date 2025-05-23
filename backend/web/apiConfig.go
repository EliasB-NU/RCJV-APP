package web

import (
	"errors"
	"fmt"
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
	err := a.PSQL.First(&config).Error
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

	// Store config in the redis database
	a.RDB.Set(a.CTX, "rcj:appEnabled", config.AppEnabled, 0)
	a.RDB.Set(a.CTX, "rcj:appName", config.EventName, 0)
	a.RDB.Set(a.CTX, "rcj:soccerURL", config.SoccerURL, 0)
	a.RDB.Set(a.CTX, "rcj:soccerAbbreviation", config.SoccerAbbreviation, 0)
	a.RDB.Set(a.CTX, "rcj:rescueURL", config.RescueURL, 0)

	// Update soccer req url
	a.RDB.Set(a.CTX, "rcj:soccerRURL", fmt.Sprintf("https://%s/rest/v1/%s", config.SoccerURL, config.SoccerAbbreviation), 0)

	return c.Status(fiber.StatusOK).JSON("config updated successfully")
}

// Check if the app should be enabled
func (a *API) getEnabled(c *fiber.Ctx) error {
	enabled, err := a.RDB.Get(a.CTX, "rcj:appEnabled").Result()
	if err != nil {
		log.Printf("Error getting enabled config: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting enabled config")
	}

	if enabled == "true" {
		return c.Status(fiber.StatusOK).JSON("App is enabled")
	} else {
		return c.Status(fiber.StatusLocked).JSON("App is disabled")
	}
}

// Get the name of the current event
func (a *API) getName(c *fiber.Ctx) error {
	name, err := a.RDB.Get(a.CTX, "rcj:appName").Result()
	if err != nil {
		log.Printf("Error getting enabled config: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting enabled config")
	}

	return c.Status(fiber.StatusOK).JSON(name)
}

func (a *API) getRescueURL(c *fiber.Ctx) error {
	url, err := a.RDB.Get(a.CTX, "rcj:rescueURL").Result()
	if err != nil {
		log.Printf("Error getting rescueURL %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting rescueURL")
	}

	return c.Status(fiber.StatusOK).JSON(url)
}
