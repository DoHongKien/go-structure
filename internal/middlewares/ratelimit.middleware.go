package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/DoHongKien/go-structure/global"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/redis"
)

type RateLimiter struct {
	globalRateLimiter     *limiter.Limiter
	publicAPIRateLimiter  *limiter.Limiter
	privateAPIRateLimiter *limiter.Limiter
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		globalRateLimiter:     rateLimiter("10000-S"),
		publicAPIRateLimiter:  rateLimiter("10-S"),
		privateAPIRateLimiter: rateLimiter("50-S"),
	}
}

// Set litmit request for all api
func rateLimiter(interval string) *limiter.Limiter {
	store, err := redis.NewStoreWithOptions(global.Rdb, limiter.StoreOptions{
		Prefix:          "rate_limiter",
		MaxRetry:        3,
		CleanUpInterval: time.Hour,
	})

	if err != nil {
		return nil
	}

	rate, err := limiter.NewRateFromFormatted(interval)

	if err != nil {
		return nil
	}

	return limiter.New(store, rate)
}

func (rl *RateLimiter) GlobalRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "global"

		log.Println("global limiter ->>>")

		limiterContext, err := rl.globalRateLimiter.Get(c, key)

		if err != nil {
			fmt.Println("Failed to check limiter global: ", err)
			c.Next()
			return
		}

		if limiterContext.Reached {
			log.Printf("Rate limit breached global %s", key)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit branched global, try later"})
			return
		}
		c.Next()
	}
}

func (rl *RateLimiter) PublicAPIRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		rateLimitPath := rl.filterLimitUrlPath(urlPath)

		if rateLimitPath != nil {
			clientIP := c.ClientIP()
			log.Printf("Path: %s, Client IP: %s", urlPath, clientIP)

			key := fmt.Sprintf("%s-%s", clientIP, urlPath)
			limitContext, err := rateLimitPath.Get(c, key)

			if err != nil {
				fmt.Println("Failed to check limiter", err)
				c.Next()
				return
			}

			if limitContext.Reached {
				fmt.Println("Rate limit reached", key)
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached, try later"})
				return
			}
		}
		c.Next()
	}
}

func (rl *RateLimiter) UserAndPrivateRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		rateLimiterPath := rl.filterLimitUrlPath(urlPath)

		if rateLimiterPath != nil {
			userId := 1

			key := fmt.Sprintf("%d:%s", userId, urlPath)
			limitContext, err := rateLimiterPath.Get(c, key)

			if err != nil {
				fmt.Println("Failed to check limiter", err)
				c.Next()
				return
			}

			if limitContext.Reached {
				log.Println("Rate limit reached", key)
				c.AbortWithStatusJSON(
					http.StatusTooManyRequests, 
					gin.H{"error": "Too many request. Rate limit breached, try later"})
				return
			}
			c.Next()
		}
	}
}

func (rl *RateLimiter) filterLimitUrlPath(urlPath string) *limiter.Limiter {
	if strings.HasPrefix(urlPath, "/api/v2/customer") {
		return rl.globalRateLimiter
	} else if strings.HasPrefix(urlPath, "/api/v2/order") {
		return rl.publicAPIRateLimiter
	} else {
		return rl.privateAPIRateLimiter
	}
}
