package web

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/websocket"
	"gorm.io/gorm"
	"log"
	"rcjv-app/backend/config"
	"rcjv-app/backend/util"
	"strings"
)

type API struct {
	CTX     context.Context
	PSQL    *gorm.DB
	RDB     *redis.Client
	CFG     *config.Config
	Clients map[*websocket.Conn]bool
}

func InitWeb(cfg *config.Config, psql *gorm.DB, rdb *redis.Client, ctx context.Context, mst *util.MST) {
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
				fiber.MethodDelete,
			}, ","),

			AllowCredentials: false,
		})
	)
	// Internal tools
	// RCJV App
	rcjvApp.Use(c)                                          // Cors Middleware
	rcjvApp.Use(healthcheck.New(healthcheck.ConfigDefault)) // Healthcheck Middleware
	rcjvApp.Get("/healthcheck", getHealthcheck)             // Healthcheck
	// RCJV game site
	rcjvGameSite.Use(c)
	rcjvGameSite.Use(healthcheck.New(healthcheck.ConfigDefault))

	// API
	apiV1 := fiber.New()
	rcjvApp.Mount("/api/v1", apiV1)
	a := API{
		CTX:     ctx,
		PSQL:    psql,
		RDB:     rdb,
		CFG:     cfg,
		Clients: make(map[*websocket.Conn]bool),
	}
	// Login
	apiV1.Post("/login", a.login)                      // <- Email&Password || -> returns new session token
	apiV1.Delete("/logout", a.logout)                  // <- Token, deletes session
	apiV1.Post("/checkLogin", a.checkIfUserIsLoggedIn) // -> Bool&Perms, checks if the session is valid and returns the users permissions
	// Admin API - Config
	apiV1.Get("/config", a.getConfig)            // [Auth] -> Returns config
	apiV1.Post("/config/update", a.updateConfig) // [Auth] <- Updates the config
	apiV1.Get("/enabled", a.getEnabled)          // -> Returns state of enabled config
	apiV1.Get("/name", a.getName)                // -> Returns the current name of the event
	apiV1.Get("/rescueURL", a.getRescueURL)      // -> Returns the Rescue Standings URL
	// Admin API - Users
	apiV1.Get("/users", a.getUsers)                 // [Auth] -> Returns all users
	apiV1.Post("/users/create", a.createUser)       // [Auth] <- Creates a new user
	apiV1.Post("/users/update/:id", a.updateUser)   // [Auth] <- Updates a user
	apiV1.Delete("/users/delete/:id", a.deleteUser) // [Auth] <- Deletes a user based on id
	// Leagues
	apiV1.Get("/leagues", a.getLeagues)             // -> Returns Leagues Body
	apiV1.Patch("/leagues/update", a.updateLeagues) // [Auth] <- Sends the struct with all leagues and updates accordingly
	// Teams
	apiV1.Get("/teams", a.getTeams)                 // -> Returns all teams with their name, league and institution
	apiV1.Post("/teams/create", a.createTeam)       // [Auth] <- Name, League and Institution, creates a new team
	apiV1.Post("/teams/update/:id", a.updateTeam)   // [Auth] <- Name, League and Institution, updates a team
	apiV1.Delete("/teams/delete/:id", a.deleteTeam) // [Auth] <- team id, deletes the team (can be restored)
	// Institutions
	apiV1.Get("/institutions", a.getInstitutions)                 // -> Returns all institutions with their name and amount of teams
	apiV1.Post("/institutions/create", a.createInstitution)       // [Auth] <- Name, creates an institution
	apiV1.Post("/institutions/update/:id", a.updateInstitution)   // [Auth] <- Name, updates an existing institution
	apiV1.Delete("/institutions/delete/:id", a.deleteInstitution) // [Auth] <- Deletes an existing institution
	// Fields
	apiV1.Get("/fields", a.getFields)                 // [Auth] -> Returns all fields
	apiV1.Post("/fields/create", a.createField)       // [Auth] <- Name&League, creates a new field
	apiV1.Post("/fields/update/:id", a.updateField)   // [Auth] <- Name&League&Id, updates an existing field
	apiV1.Delete("/fields/delete/:id", a.deleteField) // [Auth] <- Id, deletes an existing field
	// Matches
	apiV1.Get("/matches", a.getAllMatches)                                    // -> Returns all games
	apiV1.Get("/matches/league/:league", a.getMatchesLeague)                  // -> Returns all games by league
	apiV1.Get("/matches/team/:teamID", a.getMatchesTeam)                      // -> Returns all games by team
	apiV1.Get("/matches/institution/:institutionID", a.getMatchesInstitution) // -> Returns all games by institution
	apiV1.Get("/matches/field/:league/:field", a.getMatchesField)             // -> Returns all games by field (Due to sometimes similar naming conventions in different leagues, you also have to define the league
	apiV1.Post("/matches/upload/:league", a.uploadMatches)                    // [Auth] <- Uploads the ods with all the matches
	apiV1.Get("/matches/generate/:league", a.generateODS)                     // [Auth] -> Generates a new ods file with the teams of the league
	apiV1.Delete("/matches/delete/:id", a.deleteMatch)                        // [Auth] <- Delete a match based on its idea

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
	log.Println("Started RCJV V1")
	err = rcjvApp.Listen(addrRCJVApp)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
