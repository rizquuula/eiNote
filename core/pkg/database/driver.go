package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Driver struct {
	User         string
	Password     string
	DatabaseName string
	Host         string
	Port         string
}

func (d Driver) NewPostgresConn() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", d.User, d.Password, d.DatabaseName, d.Host, d.Port)
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Check if the connection is alive
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
