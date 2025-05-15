package database

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"rcjv-app/backend/config"
)

func GetRedis(cfg *config.Config) *redis.Client {
	rdp := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Database.Redis.Host, cfg.Database.Redis.Port),
		Username: cfg.Database.Redis.User,
		Password: cfg.Database.Redis.Password,
		DB:       cfg.Database.Redis.DB,
	})

	pong := rdp.Ping(context.Background())
	log.Println(pong)

	return rdp
}
