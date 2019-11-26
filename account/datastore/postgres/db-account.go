package postgres

import (
	"database/sql"
	"fmt"
	"log"

	// postgres driver
	_ "github.com/lib/pq"
)

// DB Connection
var DB *sql.DB

const (
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "postgres"
	dbPwd  = "postgres"
	dbName = "account"
)

func init() {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPwd,
		dbName,
	)

	var err error

	DB, err := sql.Open("postgres", connInfo)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	log.Printf("DB Configured successfully\n")
}
