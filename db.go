package main

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	// Database
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"

	// Config
	"github.com/spf13/viper"

	// Logging
	log "github.com/sirupsen/logrus"
)

func dbConnect() (*pgxpool.Pool, error) {
	if globalDb != nil {
		return globalDb, nil
	}

	var err error
	var config *pgxpool.Config
	dbConnection := viper.GetString("DbConnection")
	config, err = pgxpool.ParseConfig(dbConnection)
	if err != nil {
		log.Fatal(err)
	}

	// Read and parse connection lifetime
	dbPoolMaxLifeTime, errt := time.ParseDuration(viper.GetString("DbPoolMaxConnLifeTime"))
	if errt != nil {
		log.Fatal(errt)
	}
	config.MaxConnLifetime = dbPoolMaxLifeTime

	// Read and parse max connections
	dbPoolMaxConns := viper.GetInt32("DbPoolMaxConns")
	if dbPoolMaxConns > 0 {
		config.MaxConns = dbPoolMaxConns
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

func loadVersions() error {
	db, err := dbConnect()
	if err != nil {
		return err
	}
	row := db.QueryRow(context.Background(), "SELECT postgis_full_version()")
	var verStr string
	err = row.Scan(&verStr)
	if err != nil {
		return err
	}
	// Parse full version string
	//   POSTGIS="3.0.0 r17983" [EXTENSION] PGSQL="110" GEOS="3.8.0-CAPI-1.11.0 "
	//   PROJ="6.2.0" LIBXML="2.9.4" LIBJSON="0.13" LIBPROTOBUF="1.3.2" WAGYU="0.4.3 (Internal)"
	re := regexp.MustCompile(`([A-Z]+)="(.+?)"`)
	vers := make(map[string]string)
	for _, mtch := range re.FindAllStringSubmatch(verStr, -1) {
		vers[mtch[1]] = mtch[2]
	}

	pgisVer, ok := vers["POSTGIS"]
	if !ok {
		return errors.New("POSTGIS key missing from postgis_full_version")
	}
	// Convert Postgis version string into a lexically (and/or numerically) sortable form
	// "3.1.1 r17983" => "3001001"
	pgisMajMinPat := strings.Split(strings.Split(pgisVer, " ")[0], ".")
	pgisMaj, _ := strconv.Atoi(pgisMajMinPat[0])
	pgisMin, _ := strconv.Atoi(pgisMajMinPat[1])
	pgisPat, _ := strconv.Atoi(pgisMajMinPat[2])
	pgisNum := 1000000*pgisMaj + 1000*pgisMin + pgisPat
	vers["POSTGISFULL"] = strconv.Itoa(pgisNum)
	globalVersions = vers
	globalPostGISVersion = pgisNum

	return nil
}

func dBTileRequest(ctx context.Context, tr *TileRequest) ([]byte, error) {
	db, err := dbConnect()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	row := db.QueryRow(ctx, tr.SQL, tr.Args...)
	var mvtTile []byte
	err = row.Scan(&mvtTile)
	if err != nil {
		log.Warn(err)

		// check for errors retrieving the rendered tile from the database
		// Timeout errors can occur if the context deadline is reached
		// or if the context is canceled during/before a database query.
		if pgconn.Timeout(err) {
			return nil, tileAppError{
				SrcErr:  err,
				Message: fmt.Sprintf("Timeout: deadline exceeded on %s/%s", tr.LayerID, tr.Tile.String()),
			}
		}

		return nil, tileAppError{
			SrcErr:  err,
			Message: fmt.Sprintf("SQL error on %s/%s", tr.LayerID, tr.Tile.String()),
		}

	}
	return mvtTile, nil
}

func dBProjectionBoundsRequest(ctx context.Context, srid int) (float64, float64, float64, float64, error) {
	db, err := dbConnect()
	if err != nil {
		log.Error(err)
		return 0, 0, 0, 0, err
	}
	boundsSQL := fmt.Sprintf(
		"SELECT x_min, y_min, x_max, y_max FROM %s WHERE srid = $1",
		globalProjectionBoundsTableName,
	)
	row := db.QueryRow(ctx, boundsSQL, srid)

	var (
		xmin, ymin, xmax, ymax float64
	)
	err = row.Scan(&xmin, &ymin, &xmax, &ymax)
	if err != nil {
		log.Warn(err)
		return 0, 0, 0, 0, fmt.Errorf("failed to get bounds from database: %w", err)

	}
	return xmin, ymin, xmax, ymax, nil
}
