package dao_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/rinchsan/txdb-todo/pkg/config"
)

func TestMain(m *testing.M) {
	prepare()

	txdb.Register("txdb", "mysql", config.DB.DSN)

	code := m.Run()
	os.Exit(code)
}

func prepare() {
	db, err := sql.Open(config.DB.Driver, config.DB.DSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory("/go/src/github.com/rinchsan/txdb-todo/testdata/fixtures"),
	)
	if err != nil {
		panic(err)
	}

	if err := fixtures.Load(); err != nil {
		panic(err)
	}
}
