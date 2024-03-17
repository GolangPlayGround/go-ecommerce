package router

import (
	"context"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type MongoRouter struct {
	router *Router
}

func NewMongoRouter(router *Router) {
	m := &MongoRouter{
		router: router,
	}

	baseUri := "/mongo"

	m.router.GET(baseUri+"/health", m.health)
}

func (m *MongoRouter) health(c *gin.Context) {

	if errors.Is(c.Request.Context().Err(), context.DeadlineExceeded) {
		fmt.Println("에러발생")
	} else {
		fmt.Println("incoming")
	}
}
