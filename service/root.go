package service

import (
	"go-ecommerce/config"
	"go-ecommerce/repository"
	"go-ecommerce/service/mongo"
)

type Service struct {
	config *config.Config

	repository *repository.Repository

	MService *mongo.MService
}

func NewService(config *config.Config, repository *repository.Repository) (*Service, error) {
	r := &Service{
		config:     config,
		repository: repository,
		MService:   mongo.NewMService(repository),
	}
	return r, nil
}
