package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"os"
	"rcjv-app/backend/config"
	"rcjv-app/backend/data"
	"rcjv-app/backend/database"
	"rcjv-app/backend/util"
	"rcjv-app/backend/web"
)

func main() {
	// Create a new stopwatch to measure startup time
	var mst util.MST
	mst.StartTimer()

	// RCJV APP V1
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting RCJV Backend ...")

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
	go func(psql *gorm.DB, rdb *redis.Client) {
		// New context
		var ctx = context.Background()
		// Get current config
		var dbConfig database.Config
		err := psql.First(&dbConfig).Error
		if err != nil {
			log.Fatalf("Error getting config: %v\n", err)
		}

		// Store config in database
		rdb.Set(ctx, "rcjv:appEnabled", dbConfig.AppEnabled, 0)
		rdb.Set(ctx, "rcjv:appName", dbConfig.EventName, 0)
		rdb.Set(ctx, "rcjv:soccerURL", dbConfig.SoccerURL, 0)
		rdb.Set(ctx, "rcjv:soccerAbbreviation", dbConfig.SoccerAbbreviation, 0)
	}(psql, rdb)

	// Routines
	util.DeleteOldSessions(psql)
	util.DeleteSoftDeletedUserKeys(psql)

	// Soccer
	var soccer = data.Soccer{
		CTX:  context.Background(),
		PSQL: psql,
		RDB:  rdb,
	}
	soccer.FetchSoccerLeagues()

	// Init Web
	web.InitWeb(cfg, psql, rdb, &mst)
}
