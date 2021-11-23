package database

import (
	"context"
	"fmt"

	"github.com/SherinV/search-api/pkg/config"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
	klog "k8s.io/klog/v2"
)

// var Pool *pgxpool.Pool
// var err error

// const (
// 	IDLE_TIMEOUT     = 60 // ReadinessProbe runs every 30 seconds, this keeps the connection alive between probe intervals.
// 	maxConnections   = 8
// 	SINGLE_TABLE     = true
// 	CLUSTER_SHARDING = false
// 	TOTAL_CLUSTERS   = 2
// )

// func init() {
// 	glog.Info("In init dbconnector")
// 	// Pool, err = setUpDBConnection()

// 	if err != nil {
// 		glog.Error("Error connecting to db", err)
// 	}
// }

// func GetDBConnection() *pgxpool.Pool {
// 	Pool, err = setUpDBConnection()

// 	if Pool != nil {
// 		err := validateDBConnection(Pool, time.Now())

// 		if err != nil {
// 			glog.Fatal("Connection to db unsuccessful*** END")
// 			panic(err)
// 		} else {
// 			glog.Info("Connection to db successful")
// 		}
// 		return Pool
// 	} else {
// 		glog.Fatal("No Connection to db available. Pool is NIL")
// 	}
// 	return nil
// }

// func setUpDBConnection() (*pgxpool.Pool, error) {
// 	DB_HOST := config.Cfg.DB_HOST
// 	DB_USER := config.Cfg.DB_USER
// 	DB_NAME := config.Cfg.DB_NAME
// 	DB_PASSWORD := url.QueryEscape(config.Cfg.DB_PASSWORD)
// 	DB_PORT := config.Cfg.DB_PORT

// 	database_url := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
// 	glog.Info("Connecting to PostgreSQL at: ", strings.ReplaceAll(database_url, DB_PASSWORD, "*****"))
// 	config, connerr := pgxpool.ParseConfig(database_url)
// 	if connerr != nil {
// 		glog.Info("Error connecting to DB:", connerr)
// 	}
// 	config.MaxConns = maxConnections
// 	conn, err := pgxpool.ConnectConfig(context.Background(), config)

// 	if err != nil {
// 		glog.Error("Error connecting to database. Original error: ", err)
// 		return nil, err
// 	}
// 	return conn, nil
// }

// // Used by the pool to test if redis connections are still okay. If they have been idle for less than a minute,
// // just assumes they are okay. If not, calls PING.
// func validateDBConnection(c *pgxpool.Pool, t time.Time) error {
// 	// if time.Since(t) < IDLE_TIMEOUT*time.Second {
// 	// 	return nil
// 	// }
// 	err := c.Ping(context.Background()) //c.Do("PING")
// 	return err
// }

var pool *pgxpool.Pool

func init() {
	klog.Info("Initializing database connection.")
	// initializePool()
}

func initializePool() {
	cfg := config.New()

	database_url := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	klog.Info("Connecting to PostgreSQL at: ", fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", cfg.DB_USER, "*********", cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME))

	config, configErr := pgxpool.ParseConfig(database_url)
	if configErr != nil {
		klog.Error("Error parsing database connection configuration.", configErr)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		klog.Error("Unable to connect to database: %+v\n", err)
	}

	pool = conn
}

func GetConnection() *pgxpool.Pool {
	if pool == nil {
		initializePool()
	}

	if pool != nil {
		err := pool.Ping(context.Background())
		if err != nil {
			klog.Error("Unable to get a database connection. ", err)
			// Here we may need to add retry.
			return nil
		}
		klog.Info("Successfully connected to database!")
		return pool
	}
	return nil
}
