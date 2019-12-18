package main

import (
	// Database
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"

	// Config
	"github.com/spf13/viper"

	// Logging
	log "github.com/sirupsen/logrus"
)

func DbConnect() (*pgxpool.Pool, error) {
	if globalDb == nil {
		var err error
		var config *pgxpool.Config
		dbConnection := viper.GetString("DbConnection")
		config, err = pgxpool.ParseConfig(dbConnection)
		if err != nil {
			log.Fatal(err)
		}

		// Read current log level and use one less-fine level
		// below that
		config.ConnConfig.Logger = logrusadapter.NewLogger(log.New())
		levelString, _ := (log.GetLevel() - 1).MarshalText()
		pgxLevel, _ := pgx.LogLevelFromString(string(levelString))
		config.ConnConfig.LogLevel = pgxLevel

		// Connect!
		globalDb, err = pgxpool.ConnectConfig(context.Background(), config)
		if err != nil {
			log.Fatal(err)
		}
		dbName := config.ConnConfig.Config.Database
		dbUser := config.ConnConfig.Config.User
		dbHost := config.ConnConfig.Config.Host
		log.Infof("Connected as '%s' to '%s' @ '%s'", dbUser, dbName, dbHost)

		return globalDb, err
	}
	return globalDb, nil
}

func DBTileRequest(tr *TileRequest) ([]byte, error) {
	db, err := DbConnect()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	row := db.QueryRow(context.Background(), tr.Sql, tr.Args...)
	var mvtTile []byte
	err = row.Scan(&mvtTile)
	if err != nil {
		log.Warn(err)
		return nil, err
	} else {
		return mvtTile, nil
	}
}
