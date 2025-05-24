package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

type Soccer struct {
	CTX context.Context

	PSQL *gorm.DB
	RDB  *redis.Client
}

type SoccerLeagues struct {
	Leagues []struct {
		Abbreviation string `json:"abbrev"`
		Name         string `json:"name"`

		SeedingLastPublished   string `json:"seeding_last_published"`
		SeedingPublishingState string `json:"seeding_publishing_state"`

		LeagueStages []struct {
			Name string `json:"name"`

			StandingsLastPublished   string `json:"standings_last_published"`
			StandingsPublishingState string `json:"standings_publishing_state"`
		} `json:"league_stages"`
	}
	TournamentAbbrev     string `json:"tournament_abbrev"`
	TournamentName       string `json:"tournament_name"`
	TournamentHeaderType string `json:"tournament_header_type"`

	LastUpdated time.Time `json:"last_updated"`
}

// FetchSoccerLeagues fetches all the soccer leagues and stores them into the redis database.
// The values are used to cache the matches and standings of the published leagues
func (s *Soccer) FetchSoccerLeagues() {
	// Runs every minute
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for {
			<-ticker.C

			// Do the initial request
			var rUrl = strings.TrimPrefix(fmt.Sprintf("%s%s", s.RDB.Get(s.CTX, "rcj:soccerRURL").String(), "/leagues?format=json"), "get rcj:soccerRURL: ")
			agent := fiber.Get(rUrl).InsecureSkipVerify()
			statusCode, body, errs := agent.Bytes()
			if errs != nil || statusCode != 200 {
				log.Printf("Error fetching soccer leagues with code %d: %v\n", statusCode, errs)
				return
			}

			// Parse into data
			var data SoccerLeagues
			err := json.Unmarshal(body, &data)
			if err != nil {
				log.Printf("Error unmarshalling leagues: %v\n", err)
				return
			}
			data.LastUpdated = time.Now()

			// Store in redis
			err = s.RDB.JSONSet(s.CTX, "rcj:soccerLeagues", "$", data).Err()
			if err != nil {
				log.Printf("Error updating leagues in redis: %v\n", err)
				return
			}
		}
	}()
}
