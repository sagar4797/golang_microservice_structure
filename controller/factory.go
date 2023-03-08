package controller

import "database/sql"

type ControllerFactory struct {
	db *sql.DB
}

func NewControllerFactory(db *sql.DB) *ControllerFactory {
	return &ControllerFactory{db: db}
}
