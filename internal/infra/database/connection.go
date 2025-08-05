package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"prova-fattocs/internal/config"
)

func NewPostgresConnection(cfg *config.Config) *sql.DB {
	defaultDSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword)

	defaultDB, err := sql.Open("postgres", defaultDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer defaultDB.Close()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
