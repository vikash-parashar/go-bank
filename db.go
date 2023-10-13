package main

import (
	"database/sql"
)

func (s *APIServer) ConnectDB() *sql.DB {
	//TODO:
	//FIXME:
	// s.DB = sql.OpenDB(postgres.OpenDB())
	return s.DB
}
