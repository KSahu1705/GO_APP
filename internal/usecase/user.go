package usecase

import (
	"context"

	"GO_APP/internal/model/entity"
)

//UserCase is an interface
type UserCase interface {
	PutUserData(ctx context.Context) (*entity.User, error)
}
