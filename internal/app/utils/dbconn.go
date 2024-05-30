package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/configuration"

	_ "github.com/lib/pq"
)

var (
	sqlOpen = sql.Open
)

func InitDBConn(cfg configuration.DatabaseConfig) error {
	createDatabaseIfNotExist(cfg)

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

	err = initTable(db)
	if err != nil {
		log.Printf("[initDBConn] initTable() got error: %+v\n", err)
		return err
	}

	log.Println("successfully connect to database - initDBConn")
	return nil
}

func createDatabaseIfNotExist(cfg configuration.DatabaseConfig) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password)

	tempDB, err := sql.Open(cfg.Driver, psqlInfo)
	if err != nil {
		log.Fatalf("[createDatabaseIfNotExist] sql.Open() got error: %v", err)
	}
	defer tempDB.Close()

	var exists bool
	err = tempDB.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = 'simple_dating_app')").Scan(&exists)
	if err != nil {
		log.Fatalf("[createDatabaseIfNotExist] tempDB.QueryRow() got error: %+v\n", err)
	}

	if !exists {
		_, err = tempDB.Exec("CREATE DATABASE simple_dating_app")
		if err != nil {
			log.Fatalf("[createDatabaseIfNotExist] tempDB.Exec() got error: %+v\n", err)
		}
	}
}

func initTable(db *sql.DB) error {
	script := "files/query.sql"

	queries, err := os.ReadFile(script)
	if err != nil {
		log.Printf("[initTable] os.ReadFile() got error: %+v\n", err)
		return err
	}

	sql := string(queries)
	_, err = db.Exec(sql)
	if err != nil {
		log.Printf("[initTable] db.Exec() got error: %+v\n", err)
		return err
	}

	return nil
}
