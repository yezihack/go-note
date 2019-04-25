package models

import (
	"database/sql"
	"time"
	"github.com/go-sql-driver/mysql"
)

type Modeler interface {
	Create() //新增
	Update() //保存
	Find()   //查询多行
	Delete() //删除
	First()  //查询一行
	Last()   //查询最后一行
	Pluck()  //查询单列
	Count()  //统计行数
}

const (
	BookTable = "book"
)

type Book struct {
	Id           int64
	BookName     string
	BookAuthor   string
	BookProvince string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type BookNull struct {
	Id           sql.NullInt64
	BookName     sql.NullString
	BookAuthor   sql.NullString
	BookProvince sql.NullString
	CreatedAt    mysql.NullTime
	UpdatedAt    mysql.NullTime
}
type BookModel struct {
	DB *sql.DB
}

func NewBook(db *sql.DB) *BookModel {
	return &BookModel{
		DB:db,
	}
}


//获取所有的表字段
func (m *BookModel) getColumns() string {
	return " `id`,`book_name`,`book_author`,`book_province`,`updated_at`,`created_at` "
}

//获取多行数据.
func (m *BookModel) getRows(sqlTxt string, params ...interface{}) (rowsResult []*Book, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	defer query.Close()
	if err != nil {
		return
	}
	for query.Next() {
		row := BookNull{}
		err = query.Scan(
			&row.Id,           //主键
			&row.BookName,     //业务id
			&row.BookAuthor,   //业务名
			&row.BookProvince, //使用用户
			&row.UpdatedAt,    //更新时间
			&row.CreatedAt,    //添加时间
		)
		if nil != err {
			continue
		}
		rowsResult = append(rowsResult, &Book{
			Id:           row.Id.Int64, //主键
			BookName:     row.BookName.String,
			BookAuthor:   row.BookAuthor.String,
			BookProvince: row.BookProvince.String,
			UpdatedAt:    row.UpdatedAt.Time, //更新时间
			CreatedAt:    row.CreatedAt.Time, //添加时间
		})
	}
	return
}

//获取单行数据
func (m *BookModel) getRow(sql string, params ...interface{}) (rowResult *Book, err error) {
	query := m.DB.QueryRow(sql, params...)
	if err != nil {
		return
	}
	row := BookNull{}
	err = query.Scan(
		&row.Id,           //主键
		&row.BookName,     //业务id
		&row.BookAuthor,   //业务名
		&row.BookProvince, //使用用户
		&row.UpdatedAt,    //更新时间
		&row.CreatedAt,    //添加时间
	)
	if nil != err {
		return
	}
	rowResult = &Book{
		Id:           row.Id.Int64, //主键
		BookName:     row.BookName.String,
		BookAuthor:   row.BookAuthor.String,
		BookProvince: row.BookProvince.String,
		UpdatedAt:    row.UpdatedAt.Time, //更新时间
		CreatedAt:    row.CreatedAt.Time, //添加时间
	}

	return
}

//新增信息
func (m *BookModel) Create(value *Book) (lastId int64, err error) {
	sqlText := "INSERT INTO " + BookTable + "(id, book_name, book_author, book_province) VALUES (?,?,?,?)"
	stmt, err := m.DB.Prepare(sqlText)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		&value.Id,
		&value.BookName,
		&value.BookProvince,
		&value.BookProvince,
	)
	if err != nil {
		return
	}
	lastId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}
//更新数据
func (m *BookModel) saveBody(sqlTxt string, params []interface{}) (b bool, err error) {
	stmt, err := m.DB.Prepare(sqlTxt)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(params...)
	if err != nil {
		return
	}
	var affectCount int64
	affectCount, err = result.RowsAffected()
	if err != nil {
		return
	}
	b = affectCount > 0
	return
}
//更新数据
func (m *BookModel) Update(value *Book) (b bool, err error) {
	sqlText := "UPDATE " + BookTable + " SET book_name=?, book_author=?, book_province=? WHERE id=?"
	var params []interface{}
	params = append(params, value.BookName)
	params = append(params, value.BookAuthor)
	params = append(params, value.BookProvince)
	params = append(params, value.Id)
	return m.saveBody(sqlText, params)
}

//查询多行数据
func (m *BookModel) Find(value *Book) (resList []*Book, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + BookTable
	resList, err = m.getRows(sqlText)
	return
}

//获取单行数据
func (m *BookModel) First(value *Book) (result *Book, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + BookTable + " LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取单行数据
func (m *BookModel) Last(value *Book) (result *Book, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + BookTable + " ORDER BY ID DESC LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取行数
func (m *BookModel) Count() (count int64, err error) {
	sqlText := "SELECT COUNT(*) FROM " + BookTable
	query := m.DB.QueryRow(sqlText)
	if err != nil {
		return
	}
	err = query.Scan(&count)
	if err != nil {
		return
	}
	return
}
