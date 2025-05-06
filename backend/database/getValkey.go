package database

import (
	"fmt"
	"github.com/valkey-io/valkey-go"
	"log"
	"rcjv-app/backend/config"
)

func GetValkey(cfg *config.Config) valkey.Client {
	client, err := valkey.NewClient(valkey.ClientOption{
		Username:    cfg.Database.Valkey.User,
		Password:    cfg.Database.Valkey.Password,
		InitAddress: []string{fmt.Sprintf("%s:%d", cfg.Database.Valkey.Host, cfg.Database.Valkey.Port)},
		SelectDB:    0,
	})
	if err != nil {
		log.Fatalf("Failed to create valkey client: %v\n", err)
	}

	return client
}
