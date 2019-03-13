package models

import (
	"database/sql"
	"fmt"
	"testing"
	_ "github.com/go-sql-driver/mysql"
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
var DB *sql.DB
var BookDb *BookModel

func init() {
	cfg := DBConfig{
		Host:"localhost",
		Port:3308,
		Name:"root",
		Pass:"123456",
		DBName:"kindled",
		Charset:"utf8mb4",
	}
	DB = InitDB(cfg)
	BookDb = NewBook(DB)
}

func TestBookModel_Find(t *testing.T) {
	_, err := BookDb.Find(nil)
	if err != nil {
		t.Error(err)
	}
}
func TestBookModel_Create(t *testing.T) {
	bk := Book{
		BookName:"书",
		BookAuthor:"未知",
		BookProvince:"未知",
	}
	lastId, err := BookDb.Create(&bk)
	if err != nil {
		t.Error(err)
	}
	if lastId == 0 {
		t.Error("error")
	}
}
func TestBookModel_Count(t *testing.T) {
	c, err := BookDb.Count()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(c)
}

func TestBookModel_First(t *testing.T) {
	row, err := BookDb.First(nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(row)
}

func TestBookModel_Last(t *testing.T) {
	row, err := BookDb.Last(nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(row)
}
func TestBookModel_Update(t *testing.T) {
	row, err := BookDb.Update(&Book{Id: 256265, BookName:"自卑与超越"})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(row)
}
func BenchmarkBookModel_Find(b *testing.B) {
	b.StopTimer()
	result, err := BookDb.Find(nil)
	if err != nil {
		b.Error(err)
	}
	b.StartTimer()
	for _, item := range result {
		_ = item
	}
}