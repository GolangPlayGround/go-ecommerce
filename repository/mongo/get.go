package mongo

import (
	"context"
	"go-ecommerce/types"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *Mongo) GetUserBucket(user string) (*types.User, error) {

	// find, findOne, aggregate

	filter := bson.M{"user": user}

	// bson.M 조회하지만 순서보장은 안됨, bson.D (권장) , bson.A : 배열 값

	var u types.User

	if err := m.user.FindOne(context.Background(), filter).Decode(&user); err != nil {
		return nil, err
	} else {
		return &u, nil
	}

}
