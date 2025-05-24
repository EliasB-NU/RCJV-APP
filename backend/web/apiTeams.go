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

// Returns all teams in the below defined schema.
// It also adds the name of the institution
func (a *API) getTeams(c *fiber.Ctx) error {
	var (
		data []struct {
			ID            uint64 `json:"id"`
			Name          string `json:"name"`
			League        string `json:"league"`
			InstitutionID uint64 `json:"institutionID"`
			Institution   string `json:"institution"`
		}

		teams []database.Team

		err error
	)

	err = a.PSQL.Preload("Institution").Find(&teams).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error getting teams: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting teams")
	}

	for _, v := range teams {
		var t = struct {
			ID            uint64 `json:"id"`
			Name          string `json:"name"`
			League        string `json:"league"`
			InstitutionID uint64 `json:"institutionID"`
			Institution   string `json:"institution"`
		}{
			ID:            v.ID,
			Name:          v.Name,
			League:        v.League,
			InstitutionID: v.Institution.ID,
			Institution:   v.Institution.Name,
		}

		data = append(data, t)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"lastRequested": time.Now(),
		"data":          data,
	})
}

// Create a new teamBodyData
func (a *API) createTeam(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
	}

	var (
		data = struct {
			Name          string `json:"name"`
			League        string `json:"league"`
			InstitutionID uint64 `json:"institutionID"`
		}{}

		inst database.Institution
		team database.Team

		err error
	)

	// Parse body
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid Body")
	}

	// Find the institution
	err = a.PSQL.First(&inst, data.InstitutionID).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error getting institution: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting institution")
	}

	// Create the team
	team.Name = data.Name
	team.League = data.League
	team.InstitutionID = data.InstitutionID
	team.Institution = inst

	err = a.PSQL.Create(&team).Error
	if err != nil {
		log.Printf("Error creating team: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error creating team")
	}

	return c.Status(fiber.StatusOK).JSON("Team created")
}

func (a *API) updateTeam(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
	}

	// Get the team to update
	var team database.Team
	err := a.PSQL.Preload("Institution").First(&team, c.Params("id")).Error
	if err != nil {
		log.Printf("Error getting teamBodyData: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error getting teamBodyData")
	}

	var (
		data = struct {
			Name          string `json:"name"`
			League        string `json:"league"`
			InstitutionID uint64 `json:"institutionID"`
		}{}

		inst database.Institution
	)

	// Parse the body
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid Body")
	}

	// Only update the institution if it has changed
	if team.InstitutionID != data.InstitutionID {
		err = a.PSQL.First(&inst, data.InstitutionID).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Error getting institution: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).JSON("Error getting institution")
		}
		team.Institution = inst
		team.InstitutionID = data.InstitutionID
	}

	team.Name = data.Name
	team.League = data.League

	err = a.PSQL.Save(&team).Error
	if err != nil {
		log.Printf("Error updating team: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error updating team")
	}

	return c.Status(fiber.StatusOK).JSON("")
}

func (a *API) deleteTeam(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
	}

	err := a.PSQL.Delete(&database.Team{}, c.Params("id")).Error
	if err != nil {
		log.Printf("Error deleting teamBodyData: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error deleting teamBodyData")
	}

	return c.Status(fiber.StatusOK).JSON("")
}

func (a *API) csvTeam(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
	}

	return c.Status(fiber.StatusOK).JSON("")
}
