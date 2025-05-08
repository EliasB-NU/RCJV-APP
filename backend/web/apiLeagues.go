package web

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"rcjv-app/backend/database"
	"rcjv-app/backend/util"
)

// There is only on entry in this table, it gets pulled from the database
// And send to the client, the json shema is already defined in the database.Leagues struct
func (a *API) getLeagues(c *fiber.Ctx) error {
	var league database.Leagues
	err := a.PSQL.Find(&league).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Error creating database, leagues table not found")
			return c.Status(fiber.StatusInternalServerError).JSON("Leagues not found, unexpected behaviour")
		}
		log.Printf("Error getting leagues: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error fetching leagues")
	}

	return c.Status(fiber.StatusOK).JSON(league)
}

// Updates the currently activated leagues, the json is already defined in the
// database.Leagues struct so I just need to parse it and gorm models and saves it
func (a *API) updateLeagues(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("Invalid login")
	}

	var league database.Leagues
	err := a.PSQL.Find(&league).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Error creating database, leagues table not found")
			return c.Status(fiber.StatusInternalServerError).JSON("Leagues not found, unexpected behaviour")
		}
		log.Printf("Error getting leagues: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting leagues")
	}

	if err := c.BodyParser(&league); err != nil {
		log.Printf("Error updating leagues: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON("Error parsing data")
	}

	err = a.PSQL.Model(&league).Save(league).Error
	if err != nil {
		log.Printf("Error updating leagues: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error updating leagues")
	}

	return c.Status(fiber.StatusOK).JSON("Leagues updated")
}
