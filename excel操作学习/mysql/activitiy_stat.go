package mysql

import (
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/ThreeKing2018/k3log"
)

//打印错误日志
func (t *ActivitiyStat) errWrite(err error, sql string, addParams ...interface{}) {
	type ErrorType struct {
		ActivitiyStat
		AddParams []interface{} `json:"add_params"`
	}
	var infoList []interface{}
	for k, _ := range addParams {
		infoList = append(infoList, addParams[k])
		errType := ErrorType{
			ActivitiyStat: *t,
			AddParams:     infoList,
		}
		paramsJson, _ := json.Marshal(errType)
		k3log.Error("msg", "dbError", "tableName", ACTIVITIY_STAT, "error", err.Error(), "sql", sql, "params", string(paramsJson))
	}
}

//获取表的所有字段
func (t *ActivitiyStat) selectColumn() string {
	return " `id`,`stat_type`,`stat_day`,`stat_data`,`stat_remark`,`identify`,`created_at`,`updated_at` "
}

//查询数据,基础函数
func (t *ActivitiyStat) _selectBody(db *sql.DB, sqlText string, params []interface{}) (_bodyArr []*ActivitiyStat, err error) {
	_bodyArr = make([]*ActivitiyStat, 0)
	rows, err := db.Query(sqlText, params...)
	if nil != err {
		t.errWrite(err, sqlText, params)
		return
	}
	defer rows.Close()
	for rows.Next() {
		_one := nullActivitiyStat{}
		err = rows.Scan(
			&_one.Id,         //
			&_one.StatType,   //统计类型, 1 打款
			&_one.StatDay,    //统计日期
			&_one.StatData,   //统计数据
			&_one.StatRemark, //备注
			&_one.Identify,   //乐观锁
			&_one.CreatedAt,  //创建时间
			&_one.UpdatedAt,  //更新时间
		)
		if nil != err {
			t.errWrite(err, sqlText, params)
			continue
		}
		_bodyArr = append(_bodyArr, &ActivitiyStat{
			Id:         int(_one.Id.Int64),       //
			StatType:   int(_one.StatType.Int64), //统计类型, 1 打款
			StatDay:    _one.StatDay.String,      //统计日期
			StatData:   _one.StatData.String,     //统计数据
			StatRemark: _one.StatRemark.String,   //备注
			Identify:   int(_one.Identify.Int64), //乐观锁
			CreatedAt:  _one.CreatedAt.String,    //创建时间
			UpdatedAt:  _one.UpdatedAt.String,    //更新时间
		})
	}
	return
}

//更新数据基础函数
func (t *ActivitiyStat) _updateBody(db *sql.DB, sqlText string, params []interface{}) (b bool, err error) {
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
func (t *ActivitiyStat) Insert(db *sql.DB) (b bool, err error) {
	const sqlText = "INSERT INTO " + ACTIVITIY_STAT + " (stat_type,stat_day,stat_data,stat_remark,identify) VALUES (?,?,?,?,?)"
	stmt, err := db.Prepare(sqlText)
	if nil != err {
		t.errWrite(err, sqlText)
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		t.StatType,   //统计类型, 1 打款
		t.StatDay,    //统计日期
		t.StatData,   //统计数据
		t.StatRemark, //备注
		t.Identify,   //乐观锁
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
func (t *ActivitiyStat) First(db *sql.DB, id int64) (result *ActivitiyStat, err error) {
	sqlText := "SELECT" + t.selectColumn() + "FROM " + ACTIVITIY_STAT + " WHERE id = ? limit 1"
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
func (t *ActivitiyStat) Find(db *sql.DB, ids ...int) (resultList []*ActivitiyStat, err error) {
	resultList = make([]*ActivitiyStat, 0)
	sqlText := "SELECT" + t.selectColumn() + "FROM " + ACTIVITIY_STAT + " WHERE id in (" + strings.TrimRight(strings.Repeat("?,", len(ids)), ",") + ")"
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
func (t *ActivitiyStat) Update(db *sql.DB, id int) (err error) {
	sqlText := "UPDATE " + ACTIVITIY_STAT + " SET stat_type=?,stat_day=?,stat_data=?,stat_remark=?,identify=identify+1 WHERE id = ?"
	params := make([]interface{}, 0)
	params = append(params, t.StatType)
	params = append(params, t.StatDay)
	params = append(params, t.StatData)
	params = append(params, t.StatRemark)

	//主键id
	params = append(params, id)
	_, err = t._updateBody(db, sqlText, params)
	if err != nil {
		return
	}
	return
}
