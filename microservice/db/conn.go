package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

func Connect() (*sql.DB, error) {

	db, err := sql.Open(os.Getenv("DB_DRIVER"), connectionUrl())
	if err != nil {
		logrus.Errorf("Error while connecting to database: %+v", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		logrus.Errorf("Could not ping to database: %+v", err)
		return nil, err
	}

	max_conn := os.Getenv("DB_POOL_MAX_CONN")
	max_conn_db, err := strconv.Atoi(max_conn)
	if err != nil {
		logrus.Errorf("Error while setting max connection for db while parsing variable string to int: %+v", err)
		return nil, err
	}
	db.SetMaxOpenConns(max_conn_db)
	return db, nil
}

func connectionUrl() string {
	conn_url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_PASSWORD"))
	return conn_url
}
