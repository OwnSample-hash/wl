package util

import (
	"database/sql"
	"fmt"
	"store/types/cfg"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDb(cfg cfg.Database) (err error) {
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&tls=preferred", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Database)
	Db, err = sql.Open("mysql", connectionStr)
	if err != nil {
		return err
	}
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
	return nil
}
