package web

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"rcjv-app/backend/database"
)

// Returns all teams in the below defined schema.
// It also adds the name of the institution
func (a *API) getTeams(c *fiber.Ctx) error {
	var (
		data []struct {
			ID          uint64 `json:"id"`
			Name        string `json:"name"`
			League      string `json:"league"`
			Institution string `json:"institution"`
		}

		teams        []database.Team
		institutions []database.Institution

		err error
	)

	err = a.PSQL.Find(&teams).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error getting teams: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting teams")
	}

	err = a.PSQL.Find(&institutions).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error getting institutions: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting institutions")
	}

	for _, team := range teams {
		var t = struct {
			ID          uint64 `json:"id"`
			Name        string `json:"name"`
			League      string `json:"league"`
			Institution string `json:"institution"`
		}{
			ID:     team.ID,
			Name:   team.Name,
			League: team.League,
		}

		data = append(data, t)
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func (a *API) createTeam(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("")
}

func (a *API) updateTeam(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("")
}

func (a *API) deleteTeam(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("")
}

func (a *API) csvTeam(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("")
}
