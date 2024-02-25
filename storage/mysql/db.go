package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	User         string
	Password     string
	Address      string
	Port         int
	DataBaseName string
}

type MySQLDB struct {
	DB *sql.DB
}

func New(cfg Config) MySQLDB {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s(%s:%d)/%s", cfg.User, cfg.Password, cfg.Address, cfg.Port, cfg.DataBaseName))
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return MySQLDB{DB: db}

}
