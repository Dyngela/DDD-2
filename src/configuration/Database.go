package configuration

import (
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Config struct {
	Conn *gorm.DB
}

var connection *gorm.DB

func ConnectToPostgres() error {
	var err error
	var db *gorm.DB
	var conn *sql.DB
	dsn := "host=localhost user=postgres password=postgrespw dbname=test port=49153 sslmode=disable TimeZone=Europe/Berlin"

	if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true}); err != nil {
		return err
	}

	if conn, err = db.DB(); err != nil {
		return err
	}
	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(100)
	conn.SetConnMaxLifetime(time.Hour)

	connection = db
	return nil
}

func GetConfig() *Config {
	return &Config{
		Conn: connection,
	}
}
