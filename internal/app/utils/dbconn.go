package utils

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/configuration"

	_ "github.com/lib/pq"
)

var (
	sqlOpen = sql.Open
)

func InitDBConn(cfg configuration.DatabaseConfig) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sqlOpen(cfg.Driver, psqlInfo)
	if err != nil {
		log.Printf("[initDBConn] sql.Open() got error: %+v\n", err)
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Printf("[initDBConn] db.Ping() got error: %+v\n", err)
		return err
	}

	log.Println("successfully connect to database - initDBConn")
	return nil
}
