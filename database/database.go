package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
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
	url := fmt.Sprintf("file:///migrations")
	fmt.Println("url: ", url)
	m, err := migrate.NewWithDatabaseInstance(
		url,
		"gql1",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("#37", err)
		log.Fatal(err)
	}

}
