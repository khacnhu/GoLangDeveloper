package middlewares

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitRedis() *redis.Client {
	// Create a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis address (default localhost:6379)
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	// Test the connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	return rdb
}

func CacheMiddleware(rdb *redis.Client, cacheDuration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		cacheKey := c.Request.URL.Path

		// Try to get the cached response
		cachedData, err := rdb.Get(ctx, cacheKey).Result()
		if err == redis.Nil {
			// No cache, proceed to the next handler
			c.Next()

			// Cache the response for future requests
			responseData, exists := c.Get("response_data")
			if exists {
				rdb.Set(ctx, cacheKey, responseData, cacheDuration)
			}
		} else if err != nil {
			log.Printf("Error retrieving cache: %v", err)
			c.Next()
		} else {
			// Cache exists, return the cached data
			c.String(200, cachedData)
			c.Abort()
		}
	}
}
