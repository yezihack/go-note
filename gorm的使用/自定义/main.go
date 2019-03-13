package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yezihack/studyGo/gorm的使用/自定义/models"
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

func main() {
	cfg := DBConfig{
		Host:"localhost",
		Port:3308,
		Name:"root",
		Pass:"123456",
		DBName:"kindled",
		Charset:"utf8mb4",
	}
	db := InitDB(cfg)
	model := models.NewBook(db)
	res, err := model.Find(nil)
	if err !=nil {
		panic(err)
	}

}