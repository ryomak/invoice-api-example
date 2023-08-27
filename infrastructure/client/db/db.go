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

func New() (*Conn, error) {
	envCfg := env.GetCfg()
	c := mysql.Config{
		DBName:    envCfg.MySQLDatabase,
		User:      envCfg.MySQLUser,
		Passwd:    envCfg.MySQLPassword,
		Addr:      fmt.Sprintf("%s:%s", envCfg.MySQLHost, envCfg.MySQLPort),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       time.UTC,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}
	return &Conn{db}, nil
}
