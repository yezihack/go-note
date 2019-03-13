package CURD

import (
	"fmt"
	"testing"
)

var config = DBConfig{
	Host:    "localhost",
	Port:    3308,
	Name:    "root",
	Pass:    "123456",
	DBName:  "kindled",
	Charset: "utf8mb4",
}
var (
	MasterDB Modeler
)

func init() {
	db := InitDB(config)
	model := NewDB(db)
	MasterDB = model
}

/*
CREATE TABLE `book` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `book_name` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '名称',
  `book_author` varchar(125) COLLATE utf8mb4_bin NOT NULL COMMENT '作者',
  `book_province` varchar(25) COLLATE utf8mb4_bin NOT NULL COMMENT '省',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间 ',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间 ',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='书表';
*/

func TestModelS_Find(t *testing.T) {
	ls, err := MasterDB.Find("select * from book")
	if err != nil {
		t.Error(err)
	}
	for f, item := range ls {
		fmt.Println(f, item)
	}
}
func TestModelS_First(t *testing.T) {
	ls, err := MasterDB.First("select * from book where id = ?", 1)
	if err != nil {
		t.Error(err)
	}
	for f, item := range ls {
		fmt.Println(f, item)
	}
}
func TestModelS_Pluck(t *testing.T) {
	ls, err := MasterDB.Pluck("select * from book where id = ?", "book_name", 3)
	if err != nil {
		t.Error(err)
	}
	for f, item := range ls {
		fmt.Println(f, item)
	}
}
func TestModelS_Insert(t *testing.T) {
	ls, err := MasterDB.Insert("insert into book set book_name=?, book_author=?, book_province=?", "论语", "孔子", "山东")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ls)
}
func TestModelS_Update(t *testing.T) {
	ls, err := MasterDB.Update("update book set book_name=? where id=?", "国学-论语", 3)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ls)
}
func TestModelS_Delete(t *testing.T) {
	ls, err := MasterDB.Delete("delete from book where id = ?", 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ls)
}
