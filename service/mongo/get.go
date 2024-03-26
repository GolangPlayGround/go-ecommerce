package mongo

import (
	"go-ecommerce/types"
	"log"
)

func (m *MService) GetUserBucket(user string) (*types.User, error) {

	if r, err := m.repository.Mongo.GetUserBucket(user); err != nil {
		log.Println(err)
		return nil, err
	} else {
		return r, nil
	}
}
