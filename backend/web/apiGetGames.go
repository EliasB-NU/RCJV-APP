package web

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"rcjv-app/backend/database"
	"strconv"
	"time"
)

// This the struct send on all functions, better define it ones than 5 times
type payload struct {
	ID        uint64    `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`

	League    string        `json:"league"`
	Name      string        `json:"name"`
	StartTime time.Time     `json:"startTime"`
	Duration  time.Duration `json:"duration"`
	Field     string        `json:"field"`

	InstitutionID   uint64 `json:"institutionID"`
	InstitutionName string `json:"institutionName"`

	TeamID   uint64 `json:"teamID"`
	TeamName string `json:"teamName"`
}

func (a *API) getAllMatches(c *fiber.Ctx) error {
	var (
		matches []database.Match

		err error
	)

	// Payload of the message
	var data []payload

	// Get all matches and preload the team and institution
	err = a.PSQL.Preload("Team").Preload("Institution").Find(&matches).Error
	if err != nil {
		log.Printf("Error fetching games by league: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error fetching games")
	}

	//  Appends all the matches to the data
	for _, v := range matches {
		d := payload{
			ID:        v.ID,
			UpdatedAt: v.UpdatedAt,

			League:          v.League,
			Name:            v.Name,
			StartTime:       v.StartTime,
			Duration:        v.Duration,
			Field:           v.Field,
			InstitutionID:   v.InstitutionID,
			InstitutionName: v.Institution.Name,
			TeamID:          v.TeamID,
			TeamName:        v.Team.Name,
		}

		data = append(data, d)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"lastRequested": time.Now(),
		"data":          data,
	})
}

func (a *API) getMatchesLeague(c *fiber.Ctx) error {
	league := c.Params("league")
	var (
		matches []database.Match

		err error
	)

	var data []payload

	// Load all matches and preload the team and institution
	err = a.PSQL.Preload("Team").Preload("Institution").Where("league = ?", league).Find(&matches).Error
	if err != nil {
		// If the league parameter is invalid, there should be a gorm.ErrRecordNotFound Error, so I check for that separately
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON("Invalid or outdated league")
		}
		log.Printf("Error fetching games by league: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error fetching games")
	}

	// Appends all the matches to the data
	for _, v := range matches {
		d := payload{
			ID:        v.ID,
			UpdatedAt: v.UpdatedAt,

			League:    v.League,
			Name:      v.Name,
			StartTime: v.StartTime,
			Duration:  v.Duration,
			Field:     v.Field,

			InstitutionID:   v.InstitutionID,
			InstitutionName: v.Institution.Name,

			TeamID:   v.TeamID,
			TeamName: v.Team.Name,
		}

		data = append(data, d)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"lastRequested": time.Now(),
		"data":          data,
	})
}

func (a *API) getMatchesTeam(c *fiber.Ctx) error {
	var (
		teamID uint64
		err    error
	)

	// Some parsing and error checking, if the data is valid
	teamID, err = strconv.ParseUint(c.Params("teamID"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid teamID")
	}
	if teamID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid teamID")
	}

	// Load the team and preload the matches and Institution
	var team = database.Team{}
	err = a.PSQL.Preload("Matches").Preload("Institution").First(&team, teamID).Error
	if err != nil {
		log.Printf("Error fetching team: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error fetching team")
	}

	var data []payload

	// Appends all the matches to the data
	for _, m := range team.Matches {
		d := payload{
			ID:        m.ID,
			UpdatedAt: m.UpdatedAt,

			League:    m.League,
			Name:      m.Name,
			StartTime: m.StartTime,
			Duration:  m.Duration,
			Field:     m.Field,

			InstitutionID:   team.InstitutionID,
			InstitutionName: team.Institution.Name,

			TeamID:   team.ID,
			TeamName: team.Name,
		}

		data = append(data, d)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"lastRequested": time.Now(),
		"data":          data,
	})
}

func (a *API) getMatchesInstitution(c *fiber.Ctx) error {
	var (
		institutionID uint64

		err error
	)
	// Parsing and validating the institutionID
	institutionID, err = strconv.ParseUint(c.Params("institutionID"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid institutionID")
	}
	if institutionID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid institutionID")
	}

	// Get all the matches and preload the team and institution
	var matches []database.Match
	err = a.PSQL.Preload("Teams").Preload("Institution").Where("institution_id = ?", institutionID).Find(&matches).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON("Invalid institutionID")
		}
		log.Printf("Error fetching games by institution: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error fetching games")
	}

	var data []payload

	for _, m := range matches {
		d := payload{
			ID:        m.ID,
			UpdatedAt: m.UpdatedAt,

			League:    m.League,
			Name:      m.Name,
			StartTime: m.StartTime,
			Duration:  m.Duration,
			Field:     m.Field,

			InstitutionID:   institutionID,
			InstitutionName: m.Institution.Name,

			TeamID:   m.TeamID,
			TeamName: m.Team.Name,
		}

		data = append(data, d)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"lastRequested": time.Now(),
		"data":          data,
	})
}

func (a *API) getMatchesField(c *fiber.Ctx) error {
	var (
		league string
		field  string

		err error
	)
	// Parsing and validating the league and field
	league = c.Params("league")
	field = c.Params("field")
	if league == "" || field == "" {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid league or field")
	}

	// Get all the matches and preload the team and institution
	var matches []database.Match
	err = a.PSQL.Preload("Teams").Preload("Institution").Where("league = ? AND field = ?", league, field).Find(&matches).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON("Invalid institutionID")
		}
		log.Printf("Error fetching games by institution: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error fetching games")
	}

	var data []payload

	for _, m := range matches {
		d := payload{
			ID:        m.ID,
			UpdatedAt: m.UpdatedAt,

			League:    m.League,
			Name:      m.Name,
			StartTime: m.StartTime,
			Duration:  m.Duration,
			Field:     m.Field,

			InstitutionID:   m.InstitutionID,
			InstitutionName: m.Institution.Name,

			TeamID:   m.TeamID,
			TeamName: m.Team.Name,
		}

		data = append(data, d)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"lastRequested": time.Now(),
		"data":          data,
	})
}
