package main

import (
	"database/sql"
	"fmt"
)

type DBConfig struct {
	Host    string
	Port    int
	Name    string
	Pass    string
	DBName  string
	Charset string
}

//连接数据库
func InitDB(cfg DBConfig) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", cfg.Name, cfg.Pass, cfg.Host, cfg.Port, cfg.DBName, cfg.Charset)
	connection, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return connection
}
