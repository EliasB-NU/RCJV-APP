package web

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"rcjv-app/backend/database"
	"rcjv-app/backend/util"
)

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

func (a *API) activateLeague(c *fiber.Ctx) error {
	if util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("You are not authorized")
	}

	var leagues database.Leagues
	err := a.PSQL.Find(&leagues).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Error creating database, leagues table not found")
			return c.Status(fiber.StatusInternalServerError).JSON("Leagues not found, unexpected behaviour")
		}
		log.Printf("Error getting leagues: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error fetching leagues")
	}

	var league = c.Params("league")
	switch league {
	case "rescueLineEntry":
		leagues.RescueLineEntry = true
		break
	case "rescueLine":
		leagues.RescueLine = true
		break
	case "rescueMazeEntry":
		leagues.RescueMaze = true
		break
	case "rescueMaze":
		leagues.RescueMaze = true
		break
	case "soccerEntry":
		leagues.SoccerEntry = true
		break
	case "soccerLightWeightEntry":
		leagues.SoccerLightWeight = true
		break
	case "soccerLightWeight":
		leagues.SoccerLightWeight = true
		break
	case "soccerOpen":
		leagues.SoccerOpen = true
		break
	case "onStageEntry":
		leagues.OnStageEntry = true
		break
	case "onStage":
		leagues.OnStage = true
		break
	default:
		return c.Status(fiber.StatusBadRequest).JSON("League not valid")
	}
	err = a.PSQL.Save(&league).Error
	if err != nil {
		log.Printf("Error updating league table: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error updating league table")
	}

	return c.Status(fiber.StatusOK).JSON("Activated League successfully")
}

func (a *API) deactivateLeague(c *fiber.Ctx) error {
	if util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON("You are not authorized")
	}

	var leagues database.Leagues
	err := a.PSQL.Find(&leagues).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Error creating database, leagues table not found")
			return c.Status(fiber.StatusInternalServerError).JSON("Leagues not found, unexpected behaviour")
		}
		log.Printf("Error getting leagues: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error fetching leagues")
	}

	var league = c.Params("league")
	switch league {
	case "rescueLineEntry":
		leagues.RescueLineEntry = false
		break
	case "rescueLine":
		leagues.RescueLine = false
		break
	case "rescueMazeEntry":
		leagues.RescueMaze = false
		break
	case "rescueMaze":
		leagues.RescueMaze = false
		break
	case "soccerEntry":
		leagues.SoccerEntry = false
		break
	case "soccerLightWeightEntry":
		leagues.SoccerLightWeight = false
		break
	case "soccerLightWeight":
		leagues.SoccerLightWeight = false
		break
	case "soccerOpen":
		leagues.SoccerOpen = false
		break
	case "onStageEntry":
		leagues.OnStageEntry = false
		break
	case "onStage":
		leagues.OnStage = false
		break
	default:
		return c.Status(fiber.StatusBadRequest).JSON("League not valid")
	}
	err = a.PSQL.Save(&league).Error
	if err != nil {
		log.Printf("Error updating league table: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error updating league table")
	}

	return c.Status(fiber.StatusOK).JSON("Deactivated league successfully")
}
