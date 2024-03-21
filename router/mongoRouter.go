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

	m.router.GET(baseUri+"/bucket", nil) //장바구니 대한 정보 가져오기

	m.router.GET(baseUri+"/content", nil)             // 상품 정보를 가져오기
	m.router.GET(baseUri+"/user-bucket-history", nil) // 유저의 구매 이력 정보
}

func (m *MongoRouter) health(c *gin.Context) {

	if errors.Is(c.Request.Context().Err(), context.DeadlineExceeded) {
		fmt.Println("에러발생")
	} else {
		fmt.Println("incoming")
	}
}
