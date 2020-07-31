package persistence

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

func NewDB() (*gorm.DB, func()) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	for i := 0; i < 10; i++ {
		db, err := gorm.Open("postgres", psqlInfo)
		if err != nil {
			log.Info("DB not yet ready, waiting...")
			time.Sleep(time.Second)
		} else {
			log.Info("DB Connection is ready")
			db.DB().SetMaxOpenConns(maxConnections)
			return db, func() {
				err = db.Close()
				if err != nil {
					panic("Cannot close DB: " + err.Error())
				}
			}
		}
	}
	panic("Cannot connect database: " + err.Error())
}
