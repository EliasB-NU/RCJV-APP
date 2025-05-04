package main

import (
	"log"
	"rcjv-app/backend/config"
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

	// Get config
	var cfg = config.GetConfig()

	// Get databases
	var (
		psql   = database.GetPSQL(cfg)
		valkey = database.GetValkey(cfg)
	)
	defer valkey.Close()
	// Init PSQL
	err := database.InitPSQLDatabase(psql)
	if err != nil {
		log.Fatalf("Error initializing PSQL database: %v\n", err)
	}

	// Routines
	util.DeleteOldSessions(psql)
	util.DeleteSoftDeletedUserKeys(psql)

	// Init Web
	web.InitWeb(cfg, psql, valkey, &mst)
}
