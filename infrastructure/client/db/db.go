package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/ryomak/invoice-api-example/infrastructure/env"
	"time"
)

type Conn struct {
	*sql.DB
}

func newByConfig(c mysql.Config) (*Conn, error) {
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(8)
	db.SetMaxIdleConns(8)
	db.SetConnMaxIdleTime(3 * time.Second)
	return &Conn{db}, nil
}

func New() (*Conn, error) {
	envCfg := env.GetCfg()
	c := mysql.Config{
		DBName:               envCfg.MySQLDatabase,
		User:                 envCfg.MySQLUser,
		Passwd:               envCfg.MySQLPassword,
		Addr:                 fmt.Sprintf("%s:%s", envCfg.MySQLHost, envCfg.MySQLPort),
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  time.UTC,
		AllowNativePasswords: true,
	}
	return newByConfig(c)
}

func NewWithOverride(port string) (*Conn, error) {
	envCfg := env.GetCfg()
	c := mysql.Config{
		DBName:               envCfg.MySQLDatabase,
		User:                 envCfg.MySQLUser,
		Passwd:               envCfg.MySQLPassword,
		Addr:                 fmt.Sprintf("%s:%s", envCfg.MySQLHost, port),
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  time.UTC,
		AllowNativePasswords: true,
	}
	return newByConfig(c)
}
