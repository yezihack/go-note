package _0_基础

import (
	"code.qschou.com/golang/errors"
	"fmt"
	"reflect"
	"testing"
)

//反射
type Book struct {
	Id int `field:"id" type:"int(11)"`
	Name string `field:"name" type:"varchar(50)"`
}

func GetTag(t reflect.Type, name string) (field reflect.StructTag, err error) {
	f, ok := t.FieldByName(name)
	if !ok {
		err = errors.New(name + " is not exists")
		return
	}
	field = f.Tag
	return
}

func TestBook(t *testing.T) {
	var b Book
	tf := reflect.TypeOf(b)
	f, ok := tf.FieldByName("Id")
	if !ok {
		t.Error("id is not exists")
	}
	fmt.Println(f.Tag)
	fmt.Println(f.Tag.Get("field"))
	fmt.Println(f.Tag.Get("type"))
}
//获取标签里的值
func TestBook2(t *testing.T) {
	var book Book
	tt := reflect.TypeOf(book)
	tags, err := GetTag(tt, "Name")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tags.Get("field"))
	fmt.Println(tags.Get("type"))
}
//valueOf Elem获取值
func TestBook3(t *testing.T) {
	book := &Book{1, "正面管教"}
	v := reflect.ValueOf(book).Elem()
	fmt.Println(v.FieldByName("Id").Int())
	fmt.Println(v.FieldByName("Name").String())
}