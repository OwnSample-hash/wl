package util

import (
	"database/sql"
	"fmt"
	"store/types/cfg"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func getDb(cfg *cfg.Database) (db *sql.DB, err error) {
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Database)
	db, err = sql.Open("mysql", connectionStr)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}
