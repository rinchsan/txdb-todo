package dao

import (
	"context"
	"database/sql"

	"github.com/rinchsan/txdb-todo/pkg/entity"
	"github.com/rinchsan/txdb-todo/pkg/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func NewUser(db *sql.DB) repository.User {
	return userImpl{
		db: db,
	}
}

type userImpl struct {
	db *sql.DB
}

func (impl userImpl) Add(ctx context.Context, user *entity.User) error {
	if err := user.Insert(ctx, impl.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (impl userImpl) GetByID(ctx context.Context, id uint64) (*entity.User, error) {
	user, err := entity.FindUser(ctx, impl.db, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (impl userImpl) GetAll(ctx context.Context) (entity.UserSlice, error) {
	users, err := entity.Users().All(ctx, impl.db)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (impl userImpl) Update(ctx context.Context, user *entity.User) error {
	if _, err := user.Update(ctx, impl.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}
