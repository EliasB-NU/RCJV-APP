package database

import (
	glideValkey "github.com/valkey-io/valkey-glide/go/api"
	"log"
	"rcjv-app/backend/config"
)

func GetValkey(cfg *config.Config) glideValkey.GlideClientCommands {
	options := glideValkey.NewGlideClientConfiguration().
		WithAddress(&glideValkey.NodeAddress{
			Host: cfg.Database.Valkey.Host,
			Port: cfg.Database.Valkey.Port,
		})

	client, err := glideValkey.NewGlideClient(options)
	if err != nil {
		log.Fatalf("Error getting valkey: %v\n", err)
	}

	res, err := client.Ping()
	if err != nil {
		log.Fatalf("Error pinging valkey: %v\n", err)
	}
	log.Printf("Valkey Ping Response: %v\n", res)

	return client
}
