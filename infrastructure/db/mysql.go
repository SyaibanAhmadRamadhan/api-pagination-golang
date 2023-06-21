package db

import (
	"database/sql"
	"fmt"

	"github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewMysqlConnection() *sql.DB {
	dbHost := config.Get().DB.Mysql.Host
	dbPort := config.Get().DB.Mysql.Port
	dbUser := config.Get().DB.Mysql.User
	dbPass := config.Get().DB.Mysql.Pass
	dbName := config.Get().DB.Mysql.Name

	// fDB := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	fDB := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", fDB)
	if err != nil {
		log.Err(err).Dict("errors", zerolog.Dict().Str("file", "mysql.go").Str("line", "24")).Msg("error load to connect db")
	}

	// db.SetMaxIdleConns(10)
	// db.SetMaxOpenConns(100)
	// db.SetConnMaxIdleTime(5 * time.Minute)
	// db.SetConnMaxLifetime(60 * time.Minute)
	log.Info().Msg("mysql started")
	return db
}
