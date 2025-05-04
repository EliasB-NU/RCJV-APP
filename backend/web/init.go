package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	glideValkey "github.com/valkey-io/valkey-glide/go/api"
	"golang.org/x/net/websocket"
	"gorm.io/gorm"
	"log"
	"rcjv-app/backend/config"
	"rcjv-app/backend/util"
	"strings"
)

type API struct {
	PSQL    *gorm.DB
	Valkey  glideValkey.GlideClientCommands
	CFG     *config.Config
	Clients map[*websocket.Conn]bool

	LoadedEvents map[uint64]bool
}

func InitWeb(cfg *config.Config, psql *gorm.DB, valkey glideValkey.GlideClientCommands, mst *util.MST) {
	var (
		addrRCJVApp = "0.0.0.0:3006"

		err error

		// Fiber
		rcjvApp = fiber.New(fiber.Config{
			ServerHeader: "rcjv-app:fiber",
			AppName:      "rcjv-app",
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
	rcjvApp.Use(c)                                          // Cors Middleware
	rcjvApp.Use(healthcheck.New(healthcheck.ConfigDefault)) // Healthcheck Middleware
	rcjvApp.Use("/monitor", mon)                            // Monitor
	rcjvApp.Get("/healthcheck", getHealthcheck)             // Healthcheck

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

	mst.ElapsedTime()
	// Start server
	log.Println("Started Showmaster V3")
	err = rcjvApp.Listen(addrRCJVApp)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
