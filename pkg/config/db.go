package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

var DB *db

type db struct {
	Driver    string `env:"DB_DRIVER"`
	Database  string `env:"DB_DATABASE"`
	User      string `env:"DB_USER"`
	Password  string `env:"DB_PASSWORD"`
	Protocol  string `env:"DB_PROTOCOL"`
	Addr      string `env:"DB_ADDR"`
	Charset   string
	ParseTime bool
	DSN       string
}

func init() {
	db := db{
		Charset:   "utf8mb4",
		ParseTime: true,
	}

	if err := env.Parse(&db); err != nil {
		panic(err)
	}

	db.DSN = fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s&parseTime=%t",
		db.User, db.Password, db.Protocol, db.Addr, db.Database, db.Charset, db.ParseTime)

	DB = &db
}
