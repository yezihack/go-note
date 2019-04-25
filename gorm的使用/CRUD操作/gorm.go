package main

import "github.com/jinzhu/gorm"


func (Book) TableName() string {
	return "book"
}

type BookModel struct {
	DB *gorm.DB
}
