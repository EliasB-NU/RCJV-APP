package web

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"rcjv-app/backend/database"
	"time"
)

func (a *API) getInstitutions(c *fiber.Ctx) error {
	var (
		data []struct {
			ID          uint64 `json:"id"`
			Name        string `json:"name"`
			NumberTeams int    `json:"numberTeams"`
		}
	)

	var institutions []database.Institution
	err := a.PSQL.Preload("Teams").Find(&institutions).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error getting institutions: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting institutions")
	}

	for _, institution := range institutions {
		var d = struct {
			ID          uint64 `json:"id"`
			Name        string `json:"name"`
			NumberTeams int    `json:"numberTeams"`
		}{
			ID:          institution.ID,
			Name:        institution.Name,
			NumberTeams: len(institution.Teams),
		}

		data = append(data, d)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"lastUpdated": time.Now(),
		"data":        data,
	})

}
