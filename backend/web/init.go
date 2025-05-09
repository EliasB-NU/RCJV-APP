package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/valkey-io/valkey-go"
	"golang.org/x/net/websocket"
	"gorm.io/gorm"
	"log"
	"rcjv-app/backend/config"
	"rcjv-app/backend/util"
	"strings"
)

type API struct {
	PSQL    *gorm.DB
	Valkey  valkey.Client
	CFG     *config.Config
	Clients map[*websocket.Conn]bool
}

func InitWeb(cfg *config.Config, psql *gorm.DB, valkey valkey.Client, mst *util.MST) {
	var (
		addrRCJVApp      = "0.0.0.0:3006"
		addrRCJVGameSite = "0.0.0.0:3007"

		err error

		// Fiber
		rcjvApp = fiber.New(fiber.Config{
			ServerHeader: "rcjv-app:fiber",
			AppName:      "rcjv-app",
		})

		rcjvGameSite = fiber.New(fiber.Config{
			ServerHeader: "rcjv-game:fiber",
			AppName:      "rcjv-game",
		})

		// Cors
		c = cors.New(cors.Config{
			AllowOrigins: strings.Join([]string{
				"*",
			}, ","),

			AllowHeaders: strings.Join([]string{
				"Origin",
				"Content-Type",
				"Accept",
			}, ","),

			AllowMethods: strings.Join([]string{
				fiber.MethodGet,
				fiber.MethodPost,
				fiber.MethodPatch,
				fiber.MethodDelete,
			}, ","),

			AllowCredentials: false,
		})

		// Monitor
		mon = monitor.New(monitor.Config{
			Title: "Showmaster Monitor",
		})
	)
	// Internal tools
	// RCJV App
	rcjvApp.Use(c)                                          // Cors Middleware
	rcjvApp.Use(healthcheck.New(healthcheck.ConfigDefault)) // Healthcheck Middleware
	rcjvApp.Use("/monitor", mon)                            // Monitor
	rcjvApp.Get("/healthcheck", getHealthcheck)             // Healthcheck
	// RCJV game site
	rcjvGameSite.Use(c)
	rcjvGameSite.Use(healthcheck.New(healthcheck.ConfigDefault))

	// API
	api := fiber.New()
	rcjvApp.Mount("/api", api)
	a := API{
		PSQL:    psql,
		Valkey:  valkey,
		CFG:     cfg,
		Clients: make(map[*websocket.Conn]bool),
	}
	// Login
	api.Post("/login", a.login)                      // <- Email&Password || -> returns new session token
	api.Delete("/logout", a.logout)                  // <- Token, deletes session
	api.Post("/checkLogin", a.checkIfUserIsLoggedIn) // -> Bool&Perms, checks if the session is valid and returns the users permissions
	// Admin API - Config
	api.Get("/config", a.getConfig)            // [Auth] -> Returns config
	api.Post("/config/update", a.updateConfig) // [Auth] <- Updates the config
	api.Get("/enabled", a.getEnabled)          // -> Returns state of enabled config
	api.Get("/name", a.getName)                // -> Returns the current name of the event
	// Admin API - Users
	api.Get("/users", a.getUsers)                 // [Auth] -> Returns all users
	api.Post("/users/create", a.createUser)       // [Auth] <- Creates a new user
	api.Post("/users/update/:id", a.updateUser)   // [Auth] <- Updates a user
	api.Delete("/users/delete/:id", a.deleteUser) // [Auth] <- Deletes a user based on id
	// Leagues
	api.Get("/leagues", a.getLeagues)             // -> Returns Leagues Body
	api.Patch("/leagues/update", a.updateLeagues) // [Auth] <- Sends the struct with all leagues and updates accordingly
	// Teams
	api.Get("/teams", a.getTeams)                 // -> Returns all teams with their name, league and institution
	api.Post("/teams/create", a.createTeam)       // [Auth] <- Name, League and Institution, creates a new team
	api.Post("/teams/update/:id", a.updateTeam)   // [Auth] <- Name, League and Institution, updates a team
	api.Delete("/teams/delete/:id", a.deleteTeam) // [Auth] <- team id, deletes the team (can be restored)
	// Institutions
	api.Get("/institutions", a.getInstitutions) // -> Returns all institutions with their name and amount of teams
	// WebSites
	rcjvApp.Static("/", "adminsite/dist/")
	rcjvGameSite.Static("/", "webview/dist/")

	// Start WebView Server
	go func() {
		err = rcjvGameSite.Listen(addrRCJVGameSite)
		if err != nil {
			log.Fatal("Error starting server: ", err)
		}
	}()

	mst.ElapsedTime()
	// Start server
	log.Println("Started Showmaster V3")
	err = rcjvApp.Listen(addrRCJVApp)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
