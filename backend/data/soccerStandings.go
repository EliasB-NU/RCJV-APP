package data

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
	"time"
)

type SoccerStandings struct {
}

func (s *Soccer) FetchSoccerStandings() {
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for {
			<-ticker.C
			// Make initial request
			var rUrl = strings.TrimPrefix(fmt.Sprintf("%s%s", s.RDB.Get(s.CTX, "rcj:soccerRURL").String(), "/standings?format=json"), "get rcj:soccerRURL: ")
			log.Println(rUrl)
			agent := fiber.Get(rUrl).InsecureSkipVerify()
			statusCode, body, errs := agent.Bytes()
			if len(errs) > 0 {
				log.Printf("Error fetching soccer leagues with code %d: %v\n", statusCode, errs)
			}

			var data fiber.Map
			err := json.Unmarshal(body, &data)
			if err != nil {
				log.Printf("Error parsing soccer standings: %v\n", err)
			}
		}
	}()
}
