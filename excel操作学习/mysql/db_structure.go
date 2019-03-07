//数据库表内结构体信息
package mysql

import "database/sql" //
type Tmp struct {
	Id  int `json:"id"`  //
	T1  int `json:"t1"`  //
	T2  int `json:"t2"`  //
	T3  int `json:"t3"`  //
	T4  int `json:"t4"`  //
	T10 int `json:"t10"` //
	Tt  int `json:"tt"`  //

}
type nullTmp struct {
	Id  sql.NullInt64 //
	T1  sql.NullInt64 //
	T2  sql.NullInt64 //
	T3  sql.NullInt64 //
	T4  sql.NullInt64 //
	T10 sql.NullInt64 //
	Tt  sql.NullInt64 //

} //活动统计表
type ActivitiyStat struct {
	Id         int    `json:"id"`          //
	StatType   int    `json:"stat_type"`   //统计类型, 1 打款
	StatDay    string `json:"stat_day"`    //统计日期
	StatData   string `json:"stat_data"`   //统计数据
	StatRemark string `json:"stat_remark"` //备注
	Identify   int    `json:"identify"`    //乐观锁
	CreatedAt  string `json:"created_at"`  //创建时间
	UpdatedAt  string `json:"updated_at"`  //更新时间

}
type nullActivitiyStat struct {
	Id         sql.NullInt64  //
	StatType   sql.NullInt64  //统计类型, 1 打款
	StatDay    sql.NullString //统计日期
	StatData   sql.NullString //统计数据
	StatRemark sql.NullString //备注
	Identify   sql.NullInt64  //乐观锁
	CreatedAt  sql.NullString //创建时间
	UpdatedAt  sql.NullString //更新时间

} //书表
type Book struct {
	Id           int    `json:"id"`            //
	BookName     string `json:"book_name"`     //名称
	BookAuthor   string `json:"book_author"`   //作者
	BookProvince string `json:"book_province"` //省
	CreatedAt    string `json:"created_at"`    //创建时间
	UpdatedAt    string `json:"updated_at"`    //更新时间

}
type nullBook struct {
	Id           sql.NullInt64  //
	BookName     sql.NullString //名称
	BookAuthor   sql.NullString //作者
	BookProvince sql.NullString //省
	CreatedAt    sql.NullString //创建时间
	UpdatedAt    sql.NullString //更新时间

} //
type DataComment struct {
	Id   int    `json:"id"`   //
	Data string `json:"data"` //

}
type nullDataComment struct {
	Id   sql.NullInt64  //
	Data sql.NullString //

}
