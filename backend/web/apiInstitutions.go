package web

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"rcjv-app/backend/database"
	"rcjv-app/backend/util"
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

func (a *API) createInstitution(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("unauthorized")
	}

	err := a.PSQL.Create(&database.Institution{Name: c.Params("name")}).Error
	if err != nil {
		log.Printf("Error creating institution: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error creating institution")
	}

	return c.Status(fiber.StatusOK).JSON("")
}

func (a *API) updateInstitution(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("unauthorized")
	}

	var institution database.Institution
	err := a.PSQL.First(&institution, c.Params("id")).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error getting institution: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting institution")
	}

	if err := c.BodyParser(&institution); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid Body")
	}
	err = a.PSQL.Save(&institution).Error
	if err != nil {
		log.Printf("Error updating institution: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error updating institution")
	}

	return c.Status(fiber.StatusOK).JSON("")
}

func (a *API) deleteInstitution(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("unauthorized")
	}

	err := a.PSQL.Delete(&database.Institution{}, c.Params("id")).Error
	if err != nil {
		log.Printf("Error deleting institution: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error deleting institution")
	}

	return c.Status(fiber.StatusOK).JSON("")
}
