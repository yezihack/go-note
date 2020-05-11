package CURD

import (
	"database/sql"
)

type Modeler interface {
	Find(sql string, args ...interface{}) ([]map[string]interface{}, error)    //查询所有数据,即2维数据
	First(sql string, args ...interface{}) (map[string]interface{}, error)     //查询一行数据,即1维数据
	Pluck(sql string, name string, args ...interface{}) ([]interface{}, error) //查询一列数据,即1维数据 map
	Update(sql string, args ...interface{}) (int64, error)                     //更新
	Delete(sql string, args ...interface{}) (int64, error)                     //删除
	Insert(sql string, args ...interface{}) (int64, error)                     //增加
}

type ModelS struct {
	DB *sql.DB
}

func NewDB(db *sql.DB) Modeler {
	return &ModelS{
		DB: db,
	}
}

func (m *ModelS) Find(sql string, args ...interface{}) ([]map[string]interface{}, error) {
	stmt, err := m.DB.Prepare(sql)
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	//获取列名称
	var columns []string
	if columns, err = rows.Columns(); err != nil {
		return nil, err
	}
	//新建存储结果变量
	result := make([]map[string]interface{}, 0)
	count := len(columns)
	//存储数据变量
	values := make([]interface{}, count)
	//扫描变量
	scanValues := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			scanValues[i] = &values[i]
		}
		err = rows.Scan(scanValues...)
		if err != nil {
			continue
		}
		entity := make(map[string]interface{})
		for i, field := range columns {
			var v interface{}
			val := values[i]
			if b, ok := val.([]byte); ok {
				v = string(b)
			} else {
				v = val
			}
			entity[field] = v
		}
		result = append(result, entity)
	}
	return result, nil
}

//查询一行数据,即1维数据
func (m *ModelS) First(sql string, args ...interface{}) (map[string]interface{}, error) {
	stmt, err := m.DB.Prepare(sql)
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var columns []string
	columns, err = rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]interface{}, len(columns))
	scanValues := make([]interface{}, len(columns))
	result := make(map[string]interface{})
	for rows.Next() {
		for i := range columns {
			scanValues[i] = &values[i]
		}
		err = rows.Scan(scanValues...)
		if err != nil {
			continue
		}
		entity := make(map[string]interface{})
		for i, field := range columns {
			var v interface{}
			val := values[i]
			if b, ok := val.([]byte); ok {
				v = string(b)
			} else {
				v = val
			}
			entity[field] = v
		}
		result = entity
		break
	}
	return result, nil
}

//查询一列数据,即1维数据 map
func (m *ModelS) Pluck(sql string, name string, args ...interface{}) ([]interface{}, error) {
	stmt, err := m.DB.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var columns []string
	columns, err = rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]interface{}, len(columns))
	scanValues := make([]interface{}, len(columns))
	result := make([]interface{}, 0)
	for rows.Next() {
		for i := range columns {
			scanValues[i] = &values[i]
		}
		err = rows.Scan(scanValues...)
		if err != nil {
			continue
		}
		var one interface{}
		for i, field := range columns {
			var v interface{}
			val := values[i]
			if b, ok := val.([]byte); ok {
				v = string(b)
			} else {
				v = val
			}
			if field == name {
				one = v
				break
			}
		}
		result = append(result, one)
	}
	return result, nil
}

//更新
func (m *ModelS) Update(sql string, args ...interface{}) (int64, error) {
	stmt, err := m.DB.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

//删除
func (m *ModelS) Delete(sql string, args ...interface{}) (int64, error) {
	stmt, err := m.DB.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

//增加
func (m *ModelS) Insert(sql string, args ...interface{}) (int64, error) {
	stmt, err := m.DB.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}
