package configuration

import (
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func ConnectToPostgres() (*gorm.DB, error) {
	var err error
	var db *gorm.DB
	var conn *sql.DB
	dsn := "host=localhost user=postgres password=postgrespw dbname=test port=49153 sslmode=disable TimeZone=Europe/Berlin"

	if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true}); err != nil {
		return nil, err
	}

	if conn, err = db.DB(); err != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(100)
	conn.SetConnMaxLifetime(time.Hour)

	return db, nil
}
