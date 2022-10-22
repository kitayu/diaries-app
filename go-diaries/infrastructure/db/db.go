package db

import (
	"database/sql"
	"fmt"

	"github.com/kitayu/go-diaries/config"
)

func NewDB() (*sql.DB, error) {
	c := config.Config
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		c.DBUsername, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
