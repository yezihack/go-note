package mysql

import (
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/ThreeKing2018/k3log"
)

//打印错误日志
func (t *Tmp) errWrite(err error, sql string, addParams ...interface{}) {
	type ErrorType struct {
		Tmp
		AddParams []interface{} `json:"add_params"`
	}
	var infoList []interface{}
	for k, _ := range addParams {
		infoList = append(infoList, addParams[k])
		errType := ErrorType{
			Tmp:       *t,
			AddParams: infoList,
		}
		paramsJson, _ := json.Marshal(errType)
		k3log.Error("msg", "dbError", "tableName", TMP, "error", err.Error(), "sql", sql, "params", string(paramsJson))
	}
}

//获取表的所有字段
func (t *Tmp) selectColumn() string {
	return " `id`,`t1`,`t2`,`t3`,`t4`,`t10`,`tt` "
}

//查询数据,基础函数
func (t *Tmp) _selectBody(db *sql.DB, sqlText string, params []interface{}) (_bodyArr []*Tmp, err error) {
	_bodyArr = make([]*Tmp, 0)
	rows, err := db.Query(sqlText, params...)
	if nil != err {
		t.errWrite(err, sqlText, params)
		return
	}
	defer rows.Close()
	for rows.Next() {
		_one := nullTmp{}
		err = rows.Scan(
			&_one.Id,  //
			&_one.T1,  //
			&_one.T2,  //
			&_one.T3,  //
			&_one.T4,  //
			&_one.T10, //
			&_one.Tt,  //
		)
		if nil != err {
			t.errWrite(err, sqlText, params)
			continue
		}
		_bodyArr = append(_bodyArr, &Tmp{
			Id:  int(_one.Id.Int64),  //
			T1:  int(_one.T1.Int64),  //
			T2:  int(_one.T2.Int64),  //
			T3:  int(_one.T3.Int64),  //
			T4:  int(_one.T4.Int64),  //
			T10: int(_one.T10.Int64), //
			Tt:  int(_one.Tt.Int64),  //
		})
	}
	return
}

//更新数据基础函数
func (t *Tmp) _updateBody(db *sql.DB, sqlText string, params []interface{}) (b bool, err error) {
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
func (t *Tmp) Insert(db *sql.DB) (b bool, err error) {
	const sqlText = "INSERT INTO " + TMP + " (t1,t2,t3,t4,t10,tt) VALUES (?,?,?,?,?,?)"
	stmt, err := db.Prepare(sqlText)
	if nil != err {
		t.errWrite(err, sqlText)
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		t.T1,  //
		t.T2,  //
		t.T3,  //
		t.T4,  //
		t.T10, //
		t.Tt,  //
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
func (t *Tmp) First(db *sql.DB, id int64) (result *Tmp, err error) {
	sqlText := "SELECT" + t.selectColumn() + "FROM " + TMP + " WHERE id = ? limit 1"
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
func (t *Tmp) Find(db *sql.DB, ids ...int) (resultList []*Tmp, err error) {
	resultList = make([]*Tmp, 0)
	sqlText := "SELECT" + t.selectColumn() + "FROM " + TMP + " WHERE id in (" + strings.TrimRight(strings.Repeat("?,", len(ids)), ",") + ")"
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

//主键更新数据
func (t *Tmp) Update(db *sql.DB, id int) (err error) {
	sqlText := "UPDATE " + TMP + " SET t1=?,t2=?,t3=?,t4=?,t10=?,tt=? WHERE id = ?"
	params := make([]interface{}, 0)
	params = append(params, t.T1)
	params = append(params, t.T2)
	params = append(params, t.T3)
	params = append(params, t.T4)
	params = append(params, t.T10)
	params = append(params, t.Tt)

	//主键id
	params = append(params, id)
	_, err = t._updateBody(db, sqlText, params)
	if err != nil {
		return
	}
	return
}
