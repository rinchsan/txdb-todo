package dao_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rinchsan/txdb-todo/pkg/dao"
	"github.com/rinchsan/txdb-todo/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestUserImpl_Add(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		user  *entity.User
		noErr bool
	}{
		"new user": {
			user:  &entity.User{Username: "rinchsan"},
			noErr: true,
		},
		"duplicate username": {
			user:  &entity.User{Username: "John"},
			noErr: true,
		},
		"empty username": {
			user:  &entity.User{Username: ""},
			noErr: true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, err := sql.Open("txdb", uuid.New().String())
			assert.NoError(t, err)
			defer db.Close()
			impl := dao.NewUser(db)

			err = impl.Add(context.Background(), c.user)
			if c.noErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestUserImpl_GetByID(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		id    uint64
		user  *entity.User
		noErr bool
	}{
		"user 1 exists": {
			id: 1,
			user: &entity.User{
				ID:        1,
				Username:  "John",
				CreatedAt: time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
				UpdatedAt: time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
			},
			noErr: true,
		},
		"user 2 exists": {
			id: 2,
			user: &entity.User{
				ID:        2,
				Username:  "Charles",
				CreatedAt: time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
				UpdatedAt: time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
			},
			noErr: true,
		},
		"user 100 does not exist": {
			id:    100,
			user:  nil,
			noErr: false,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, err := sql.Open("txdb", uuid.New().String())
			assert.NoError(t, err)
			defer db.Close()
			impl := dao.NewUser(db)

			user, err := impl.GetByID(context.Background(), c.id)
			assert.Equal(t, c.user, user)
			if c.noErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestUserImpl_GetAll(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		users entity.UserSlice
		noErr bool
	}{
		"all users": {
			users: entity.UserSlice{
				&entity.User{
					ID:        1,
					Username:  "John",
					CreatedAt: time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
					UpdatedAt: time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
				},
				&entity.User{
					ID:        2,
					Username:  "Charles",
					CreatedAt: time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
					UpdatedAt: time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
				},
				&entity.User{
					ID:        3,
					Username:  "Herbert",
					CreatedAt: time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
					UpdatedAt: time.Date(2020, time.December, 29, 23, 59, 59, 0, time.UTC),
				},
			},
			noErr: true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, err := sql.Open("txdb", uuid.New().String())
			assert.NoError(t, err)
			defer db.Close()
			impl := dao.NewUser(db)

			users, err := impl.GetAll(context.Background())
			assert.Equal(t, len(c.users), len(users))
			assert.Equal(t, c.users, users)
			if c.noErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestUserImpl_Update(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		user  *entity.User
		noErr bool
	}{
		"user 1": {
			user: &entity.User{
				ID:       1,
				Username: "new name",
			},
			noErr: true,
		},
		"user 2": {
			user: &entity.User{
				ID:       2,
				Username: "new name",
			},
			noErr: true,
		},
		"empty username": {
			user: &entity.User{
				ID:       3,
				Username: "",
			},
			noErr: true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, err := sql.Open("txdb", uuid.New().String())
			assert.NoError(t, err)
			defer db.Close()
			impl := dao.NewUser(db)

			err = impl.Update(context.Background(), c.user)
			if c.noErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
