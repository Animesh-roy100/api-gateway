package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
)

var cb *gobreaker.CircuitBreaker

func init() {
	var settings gobreaker.Settings
	settings.Name = "API-Gateway"
	settings.MaxRequests = 3
	settings.Interval = 10 * time.Second
	settings.Timeout = 60 * time.Second
	settings.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	settings.OnStateChange = func(name string, from gobreaker.State, to gobreaker.State) {
		if to == gobreaker.StateOpen {
			log.Println("State Open!")
		}
		if from == gobreaker.StateOpen && to == gobreaker.StateHalfOpen {
			log.Println("Going from Open to Half Open")
		}
		if from == gobreaker.StateHalfOpen && to == gobreaker.StateClosed {
			log.Println("Going from Half Open to Closed!")
		}
	}

	cb = gobreaker.NewCircuitBreaker(settings)
}

func CircuitBreaker() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := cb.Execute(func() (interface{}, error) {
			c.Next()
			return nil, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(503, gin.H{"error": "service unavailable"})
			return
		}
	}
}
