package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var Db *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "abhinav:Ab9450455997@tcp(192.168.1.2:3306)/gql1?multiStatements=true")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	url := fmt.Sprintf("file://migrations")
	m, err := migrate.NewWithDatabaseInstance(url, "gql1", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("#37", err)
		log.Fatal(err)
	}

}
