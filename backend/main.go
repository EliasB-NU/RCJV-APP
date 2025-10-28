package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"rcjv-app/backend/config"
	"rcjv-app/backend/data"
	"rcjv-app/backend/database"
	"rcjv-app/backend/util"
	"rcjv-app/backend/web"

	"github.com/redis/go-redis/v9"
)

func main() {
	// Create a new stopwatch to measure startup time
	var mst util.MST
	mst.StartTimer()

	// RCJV APP V1
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting RCJV Backend ...")

	// Create systemwide context
	var ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	// Check if dev or prod, create tmp dir if dev
	if os.Args[1] == "dev" {
		err := os.MkdirAll("./tmp", os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating ./tmp directory: %v\n", err)
		}
	}

	// Get config
	var cfg = config.GetConfig()

	// Get databases
	var (
		psql = database.GetPSQL(cfg)
		rdb  = database.GetRedis(cfg)
	)
	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {
			log.Fatalf("Error closing redis connection: %v\n", err)
		}
	}(rdb)
	// Init PSQL
	err := database.InitPSQLDatabase(psql)
	if err != nil {
		log.Fatalf("Error initializing PSQL database: %v\n", err)
	}

	// Cache config
	// Get current config
	var dbConfig database.Config
	err = psql.First(&dbConfig).Error
	if err != nil {
		log.Fatalf("Error getting config: %v\n", err)
	}

	// Store config in database
	rdb.Set(ctx, "rcj:appEnabled", dbConfig.AppEnabled, 0)
	rdb.Set(ctx, "rcj:appName", dbConfig.EventName, 0)
	rdb.Set(ctx, "rcj:soccerURL", dbConfig.SoccerURL, 0)
	rdb.Set(ctx, "rcj:soccerAbbreviation", dbConfig.SoccerAbbreviation, 0)
	rdb.Set(ctx, "rcj:rescueURL", dbConfig.RescueURL, 0)

	// Routines
	util.DeleteOldSessions(psql)
	util.DeleteSoftDeletedUserKeys(psql)

	// Soccer
	// Create url and initiate interface
	url, err := rdb.Get(ctx, "rcj:soccerURL").Result()
	if err != nil {
		log.Printf("Error fetching url from RDB: %v\n", err)
	}
	abbrev, err := rdb.Get(ctx, "rcj:soccerAbbreviation").Result()
	if err != nil {
		log.Printf("Error fetching abbreviation from RDB: %v\n", err)
	}
	// https:// is already in the url you should enter in the frontend
	rdb.Set(ctx, "rcj:soccerRURL", fmt.Sprintf("%s/rest/v1/%s", url, abbrev), 0)
	var soccer = data.Soccer{
		CTX:  ctx,
		PSQL: psql,
		RDB:  rdb,
	}
	// Run fetch functions
	soccer.FetchSoccerLeagues()
	soccer.FetchSoccerMatches()
	// soccer.FetchSoccerStandings()

	// Init Web
	web.InitWeb(cfg, psql, rdb, ctx, &mst)
}
