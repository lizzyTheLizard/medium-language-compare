package persistence

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	host           = "postgres"
	port           = 5432
	user           = "postgres"
	password       = "postgres"
	dbname         = "postgres"
	maxConnections = 100
)

var db *sql.DB

func Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic("Cannot connect to database: %w", err)
	}
	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err != nil {
			log.Info("DB not yet ready, waiting...")
			time.Sleep(time.Second)
		} else {
			log.Info("DB Connection is ready")
			db.SetMaxOpenConns(maxConnections)
			return
		}
	}
	log.Panic("Cannot ping database: %w", err)

}

func Close() {
	err := db.Close()
	if err != nil {
		log.Panic("Cannot close database connection: %w", err)
	}
}
