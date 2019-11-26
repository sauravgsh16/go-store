package postgres

import (
	"database/sql"
	"fmt"
	"log"

	// postgres driver
	_ "github.com/lib/pq"
)

const (
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "postgres"
	dbPwd  = "postgres"
	dbName = "account"
)

// GetDBConn returns a DB connection
func GetDBConn() *sql.DB {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPwd,
		dbName,
	)

	var err error

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	log.Printf("DB Configured successfully\n")
	return db
}
