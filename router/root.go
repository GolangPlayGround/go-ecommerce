package router

import (
	"context"
	"errors"
	"go-ecommerce/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Router struct {
	config *config.Config

	engine *gin.Engine
}

func NewRouter(config *config.Config) (*Router, error) {
	r := &Router{
		config: config,
		engine: gin.New(),
	}

	r.engine.Use(requestTimeOutMiddleWare(5 * time.Second))

	NewMongoRouter(r)

	return r, r.engine.Run(config.ServerInfo.Port)
}

func requestTimeOutMiddleWare(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		done := make(chan struct{})

		go func() {
			defer close(done)
			c.Next()
		}()

		select {
		case <-done:
			return
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{"error": "request time out"})
				return
			}
		}
	}
}
