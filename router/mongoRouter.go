package router

import (
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
	fmt.Println("incoming")
}