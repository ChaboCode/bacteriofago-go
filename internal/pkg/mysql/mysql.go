package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/golang-migrate/migrate/v4"
)

var Db *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:dbpass@tcp(localhost)/bacteriofago")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
}

func CloseDB() error {
	return Db.Close()
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}

	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/mysql",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
