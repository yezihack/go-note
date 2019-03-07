package mysql

import (
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/ThreeKing2018/k3log"
)

//打印错误日志
func (t *DataComment) errWrite(err error, sql string, addParams ...interface{}) {
	type ErrorType struct {
		DataComment
		AddParams []interface{} `json:"add_params"`
	}
	var infoList []interface{}
	for k, _ := range addParams {
		infoList = append(infoList, addParams[k])
		errType := ErrorType{
			DataComment: *t,
			AddParams:   infoList,
		}
		paramsJson, _ := json.Marshal(errType)
		k3log.Error("msg", "dbError", "tableName", DATA_COMMENT, "error", err.Error(), "sql", sql, "params", string(paramsJson))
	}
}

//获取表的所有字段
func (t *DataComment) selectColumn() string {
	return " `id`,`data` "
}

//查询数据,基础函数
func (t *DataComment) _selectBody(db *sql.DB, sqlText string, params []interface{}) (_bodyArr []*DataComment, err error) {
	_bodyArr = make([]*DataComment, 0)
	rows, err := db.Query(sqlText, params...)
	if nil != err {
		t.errWrite(err, sqlText, params)
		return
	}
	defer rows.Close()
	for rows.Next() {
		_one := nullDataComment{}
		err = rows.Scan(
			&_one.Id,   //
			&_one.Data, //
		)
		if nil != err {
			t.errWrite(err, sqlText, params)
			continue
		}
		_bodyArr = append(_bodyArr, &DataComment{
			Id:   int(_one.Id.Int64), //
			Data: _one.Data.String,   //
		})
	}
	return
}

//更新数据基础函数
func (t *DataComment) _updateBody(db *sql.DB, sqlText string, params []interface{}) (b bool, err error) {
	stmt, err := db.Prepare(sqlText)
	if nil != err {
		t.errWrite(err, sqlText, params)
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		params...,
	)
	_count, err := res.RowsAffected()
	if nil != err {
		t.errWrite(err, sqlText, params)
		return
	}
	return _count > 0, nil
}

//插入数据
func (t *DataComment) Insert(db *sql.DB) (b bool, err error) {
	const sqlText = "INSERT INTO " + DATA_COMMENT + " (data) VALUES (?)"
	stmt, err := db.Prepare(sqlText)
	if nil != err {
		t.errWrite(err, sqlText)
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		t.Data, //
	)
	if nil != err {
		t.errWrite(err, sqlText)
		return
	}
	n, err := res.RowsAffected()
	if nil != err {
		t.errWrite(err, sqlText)
		return
	}
	return n > 0, nil
}

//查询单列数据
func (t *DataComment) First(db *sql.DB, id int64) (result *DataComment, err error) {
	sqlText := "SELECT" + t.selectColumn() + "FROM " + DATA_COMMENT + " WHERE id = ? limit 1"
	list, err := t._selectBody(db, sqlText, []interface{}{id})
	if err != nil {
		return
	}
	if len(list) > 0 {
		result = list[0]
	}
	return
}

//查询多条数据,传入id,id2,id3
func (t *DataComment) Find(db *sql.DB, ids ...int) (resultList []*DataComment, err error) {
	resultList = make([]*DataComment, 0)
	sqlText := "SELECT" + t.selectColumn() + "FROM " + DATA_COMMENT + " WHERE id in (" + strings.TrimRight(strings.Repeat("?,", len(ids)), ",") + ")"
	var args []interface{}
	for _, id := range ids {
		args = append(args, id)
	}
	resultList, err = t._selectBody(db, sqlText, args)
	if err != nil {
		return
	}
	return
}

//查询多条数据,传入id,id2,id3
func (t *DataComment) Select(db *sql.DB) (resultList []*DataComment, err error) {
	resultList = make([]*DataComment, 0)
	sqlText := "SELECT" + t.selectColumn() + "FROM " + DATA_COMMENT
	var args []interface{}
	resultList, err = t._selectBody(db, sqlText, args)
	if err != nil {
		return
	}
	return
}

//主键更新数据
func (t *DataComment) Update(db *sql.DB, id int) (err error) {
	sqlText := "UPDATE " + DATA_COMMENT + " SET data=? WHERE id = ?"
	params := make([]interface{}, 0)
	params = append(params, t.Data)

	//主键id
	params = append(params, id)
	_, err = t._updateBody(db, sqlText, params)
	if err != nil {
		return
	}
	return
}
