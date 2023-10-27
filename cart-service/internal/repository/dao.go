package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type DAO interface {
	NewCartQuery() CartQuery
}

type dao struct{}

var rdb *redis.Client

func RDS() (*redis.Client, error) {
	// config, err := util.LoadConfig("../../")
	// // fmt.Println(config.RedisAddress)
	// if err != nil {
	// 	log.Fatalln("config error")
	// 	return nil, err
	// }

	rdb = redis.NewClient(&redis.Options{
		// Addr: "redis:6379", // Ganti dengan alamat dan port Redis Anda.
		Addr: "redis:6379",
		DB:   0, // Gunakan database Redis default (0).
	})

	// Use a background context
	ctx := context.Background()

	// Check the connection to the Redis server
	err := rdb.Ping(ctx).Err()
	if err != nil {
		// Log the error and return the error to the caller
		log.Printf("Error connecting to Redis: %v\n", err)
		return nil, fmt.Errorf("error connecting to Redis: %w", err)
	}

	// Log success and return the Redis client
	log.Println("Connected to Redis")

	return rdb, nil
}

func NewDAO() DAO {
	return &dao{}
}

func (d *dao) NewCartQuery() CartQuery {
	return &cartQuery{}
}
