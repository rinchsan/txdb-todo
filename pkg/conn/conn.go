package conn

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rinchsan/txdb-todo/pkg/config"
)

var (
	DB *sql.DB
)

func SetupDB() func() {
	var err error
	DB, err = sql.Open(config.DB.Driver, config.DB.DSN)
	if err != nil {
		panic(err)
	}

	maxOpenConns := 30
	DB.SetMaxOpenConns(maxOpenConns)
	DB.SetMaxIdleConns(maxOpenConns)
	DB.SetConnMaxLifetime(time.Duration(maxOpenConns) * time.Second)

	return func() {
		DB.Close()
	}
}
