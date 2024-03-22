package router

import (
	"context"
	"errors"
	"fmt"
	"go-ecommerce/service/mongo"
	"go-ecommerce/types"

	"github.com/gin-gonic/gin"
)

type MongoRouter struct {
	router   *Router
	mService *mongo.MService
}

func NewMongoRouter(router *Router, mService *mongo.MService) {
	m := &MongoRouter{
		router: router,
	}

	baseUri := "/mongo"

	m.router.GET(baseUri+"/health", m.health)

	m.router.GET(baseUri+"/user-bucket", m.userBucket) //장바구니 대한 정보 가져오기

	m.router.GET(baseUri+"/content", nil)             // 상품 정보를 가져오기
	m.router.GET(baseUri+"/user-bucket-history", nil) // 유저의 구매 이력 정보
}

func (m *MongoRouter) userBucket(c *gin.Context) {
	var req types.BucketRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		// TODO 에러 처리
	} else if r, err := m.mService.GetUserBucket(req.User); err != nil {
		//TODO
	} else {
		//TODO
		fmt.Println(r)
	}
}

func (m *MongoRouter) health(c *gin.Context) {

	if errors.Is(c.Request.Context().Err(), context.DeadlineExceeded) {
		fmt.Println("에러발생")
	} else {
		fmt.Println("incoming")
	}
}
