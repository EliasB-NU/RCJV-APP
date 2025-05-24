package data

import (
  "encoding/json"
  "fmt"
  "github.com/gofiber/fiber/v2"
  "log"
  "strings"
  "time"
)

type SoccerMatches struct {
  Matches []struct {
    Number int `json:"number"`

    League      string `json:"league"`
    LeagueStage string `json:"league_stage"`
    GroupName   string `json:"group_name"`
    Start       string `json:"start"`
    Duration    string `json:"duration"`
    Field       string `json:"pitch"`

    Goals1  int `json:"goals1"`
    Points1 int `json:"points1"`
    Team1   struct {
      ID          int    `json:"id"`
      Name        string `json:"name"`
      Institution string `json:"affiliation"`
      StartNumber int    `json:"startnumber"`
      ExternalKey string `json:"external_key"`
    } `json:"team1"`

    Goals2  int `json:"goals2"`
    Points2 int `json:"points2"`
    Team2   struct {
      ID          int    `json:"id"`
      Name        string `json:"name"`
      Institution string `json:"affiliation"`
      StartNumber int    `json:"startnumber"`
      ExternalKey string `json:"external_key"`
    } `json:"team2"`
    Referees []struct {
      FirstName string `json:"first_name"`
      LastName  string `json:"last_name"`
    } `json:"referees"`
  } `json:"matches"`

  TournamentAbbrev     string `json:"tournament_abbrev"`
  TournamentName       string `json:"tournament_name"`
  TournamentHeaderType string `json:"tournament_header_type"`

  LastPublished string    `json:"last_published"`
  LastUpdated   time.Time `json:"last_updated"`
}

func (s *Soccer) FetchSoccerMatches() {
  ticker := time.NewTicker(1 * time.Minute)
  go func() {
    for {
      <-ticker.C

      // Do initial request
      var rUrl = strings.TrimPrefix(fmt.Sprintf("%s%s", s.RDB.Get(s.CTX, "rcj:soccerRURL").String(), "/matches?format=json"), "get rcj:soccerRURL: ")
      agent := fiber.Get(rUrl).InsecureSkipVerify()
      statusCode, body, errs := agent.Bytes()
      if len(errs) > 0 || statusCode != 200 {
        log.Printf("Error fetching soccer matches with code %d: %v\n", statusCode, errs)
      }

      // Parse body
      var data SoccerMatches
      err := json.Unmarshal(body, &data)
      if err != nil {
        log.Printf("Error unmarshalling soccer matches: %v\n", err)
      }

      data.LastUpdated = time.Now()

      // Store in redis
      err = s.RDB.JSONSet(s.CTX, "rcj:soccerMatches", "$", data).Err()
      if err != nil {
        log.Printf("Error updating soccer matches: %v\n", err)
      }
    }
  }()
}
