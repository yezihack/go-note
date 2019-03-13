package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)



func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3308)/kindled?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	Find(db)
}

func Find(db *gorm.DB) {
	//bk := make([]Book, 0)

	bk := new(Book)
	spew.Dump(bk)
	db.First(&bk, "id=?", 7)
	fmt.Println(bk)
}
//新增
func Create(db *gorm.DB) {
	var bk Book
	bk.BookName = "相约星期二"
	bk.BookAuthor = "米奇·阿尔博姆"
	bk.BookProvince = "美国"
	db.Create(&bk)
	fmt.Println(bk.Id) //获取新增ID
}

type Book struct {
	gorm.Model
	Id int `gorm:"id"`
	BookName string `gorm:"book_name"`
	BookAuthor string `gorm:"book_author"`
	BookProvince string `gorm:"book_province"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

func (Book) TableName() string {
	return "book"
}