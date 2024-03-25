package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/trile99/go_api_server/internal/app/databases"
)

func CacheMiddleware(redisClient *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()

		cachedData, err := redisClient.Get(ctx, "cached_data").Result()
		if err == redis.Nil {
			data := databases.DB.Db
			if err != nil {
				log.Error().Err(err).Msg("Failed to fetch data from database")
				return c.Next()
			}

			err = redisClient.Set(ctx, "cached_data", data, 24*time.Hour).Err()
			if err != nil {
				log.Error().Err(err).Msg("Failed to cache data")
			}
		} else if err != nil {
			log.Error().Err(err).Msg("Error checking cache")
		} else {
			fmt.Println("Data from cache:", cachedData)
			c.Locals("cached_data", cachedData)
		}

		return c.Next()
	}
}
