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

func Connect() (*sql.DB, func()) {
	db := open()
	for i := 0; i < 10; i++ {
		if isConnectionUp(db) {
			return db, func() { close(db) }
		}
		time.Sleep(time.Second)
	}
	panic("Cannot connect database")
}

func open() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic("Cannot connect to database: %w", err)
	}
	db.SetMaxOpenConns(maxConnections)
	return db
}

func close(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic("Cannot close DB: " + err.Error())
	}
}

func isConnectionUp(db *sql.DB) bool {
	err := db.Ping()
	if err != nil {
		log.Debug("DB not yet ready: " + err.Error())
		log.Info("Waiting for DB to come up")
		return false
	}
	log.Info("DB is ready")
	return true
}
