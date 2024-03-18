package mongo

import (
	"context"
	"go-ecommerce/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	config *config.Config

	client *mongo.Client
	db     *mongo.Database
}

func NewMongo(config *config.Config) (*Mongo, error) {
	m := &Mongo{
		config: config,
	}

	ctx := context.Background()
	var err error

	if m.client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Mongo.URI)); err != nil {
		panic(err)
	} else if err = m.client.Ping(ctx, nil); err != nil {
		panic(err)
	} else {
		m.db = m.client.Database(config.Mongo.Db)
		// TODO 컬렉션 연결
	}

	return m, nil
}
