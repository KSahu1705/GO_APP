package category

import (
	"GO_APP/internal/model/config"
	"GO_APP/internal/repository"
)

//Implementation is a struct
type Implementation struct {
	Config                        *config.Config
	UserRepository            	  repository.User
}

//New returned Implementation struct
func New(cfg *config.Config, payload Implementation) *Implementation {
	return &Implementation{
		Config:                        cfg,
		UserRepository:            	   payload.UserRepository,
	}
}