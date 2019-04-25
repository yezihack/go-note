package mysql_orm

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
	"github.com/yezihack/now"
	"fmt"
	"sync"
	"runtime"
)

type DBPool struct {
	OrmDB *gorm.DB
	DB    *sql.DB
}

var Pools DBPool

func init() {
	var err error
	Pools.OrmDB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3308)/sys?charset=utf8")
	if err != nil {
		panic(err)
	}
	Pools.DB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3308)/sys?charset=utf8")
	if err != nil {
		panic(err)
	}
}

type Category struct {
	Id           int    `gorm:"primary_key"`
	CategoryName string `gorm:"category_name"`
	CreatedAt    string `gorm:"created_at"`
	UpdatedAt    string `gorm:"updated_at"`
}
type BlackList struct {
	Id            int    `gorm:"primary_key"`
	UserId        int    `grom:"user_id"`
	CategoryId    int    `gorm:"category_id"`
	FromProjectId int    `gorm:"from_project_id"`
	DataValue     string `gorm:"data_value"`
	IsUse         int    `gorm:"is_use"`
	CreatedAt     string `gorm:"created_at"`
	UpdatedAt     string `gorm:"updated_at"`
}

//原生查询数据库
func GetListByOri() []Category {
	sqlText := `select id, category_name, created_at, updated_at from category`
	rows, err := Pools.DB.Query(sqlText)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	result := make([]Category, 0)
	for rows.Next() {
		item := Category{}
		rows.Scan(
			&item.Id,
			&item.CategoryName,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		result = append(result, item)
	}
	return result
}

// orm查询数据库
func GetListByOrm() []Category {
	var cate []Category
	Pools.OrmDB.Table("category").Find(&cate)
	return cate
}

//原生查询数据库
func GetBlackListByOri() []BlackList {
	sqlText := `select id, user_id, from_project_id, data_value, is_use, created_at, updated_at from blacklist`
	rows, err := Pools.DB.Query(sqlText)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	result := make([]BlackList, 0)
	for rows.Next() {
		item := BlackList{}
		rows.Scan(
			&item.Id,
			&item.UserId,
			&item.FromProjectId,
			&item.DataValue,
			&item.IsUse,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		result = append(result, item)
	}
	return result
}

// orm查询数据库
func GetBlackListByOrm() []BlackList {
	var cate []BlackList
	Pools.OrmDB.Table("blacklist").Find(&cate)
	return cate
}

func InsertBlackList() {
	ShowNumberGoruntine()
	wg := sync.WaitGroup{}
	wg.Add(10)
	for j := 0; j < 10; j ++ {
		go func(j int, wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 0; i < 100; i ++ {
				fmt.Printf("j:%d, i:%d\n", j, i)
				var list BlackList
				list.UserId = rand.Intn(100)
				list.IsUse = 1
				list.DataValue = "{}"
				list.FromProjectId = rand.Intn(20000)
				list.CreatedAt = new(now.Now).Format("2006-01-02 15:04:05")
				list.UpdatedAt = new(now.Now).Format("2006-01-02 15:04:05")
				Pools.OrmDB.Table("blacklist").Create(&list)
			}

		}(j,&wg)
	}
	wg.Wait()
}
func ShowNumberGoruntine() {
	go func() {
		fmt.Printf("目前有%d个goruntine", runtime.NumGoroutine())
	}()
}