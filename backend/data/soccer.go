package data

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
)

type Soccer struct {
	CTX context.Context

	PSQL *gorm.DB
	RDB  *redis.Client
}

type SoccerLeagues struct {
	Abbreviation string   `json:"abbrev"`
	Name         string   `json:"name"`
	LeagueStages []string `json:"leagues"`
}

func (s *Soccer) FetchSoccerLeagues() {
	url, err := s.RDB.Get(s.CTX, "rcjv:soccerURL").Result()
	if err != nil {
		log.Printf("Error fetching url from RDB: %v\n", err)
	}
	abbrev, err := s.RDB.Get(s.CTX, "rcjv:soccerAbbreviation").Result()
	if err != nil {
		log.Printf("Error fetching abbreviation from RDB: %v\n", err)
	}
	rURL := fmt.Sprintf("https://%s/rest/v1/%s", url, abbrev)

	log.Println(rURL)
}
