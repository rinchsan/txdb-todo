package repository

import (
	"context"

	"github.com/rinchsan/txdb-todo/pkg/entity"
)

//go:generate mockgen -source=./user.go -destination=./mock/user.go -package=mock -mock_names=User=User
type User interface {
	Add(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id uint64) (*entity.User, error)
	GetAll(ctx context.Context) (entity.UserSlice, error)
	Update(ctx context.Context, user *entity.User) error
}
