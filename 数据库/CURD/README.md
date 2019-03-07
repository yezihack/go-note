# golang 数据库操作
> 实现golang的数据库增删改查操作,无需改表结果,快捷操作数据库

## CURD

### 初使数据库连接
```
var config = DBConfig{
	Host:    "localhost",
	Port:    3308,
	Name:    "root",
	Pass:    "123456",
	DBName:  "kindled",
	Charset: "utf8mb4",
}
var (
	MasterDB Modeler //定义主DB
)

func init() {
	db := InitDB(config)
	model := NewDB(db)
	MasterDB = model
}
```

### 查询所有数据

```
ls, err := MasterDB.Find("select * from book")
if err != nil {
    t.Error(err)
}
for f, item := range ls {
    fmt.Println(f, item)
}
```

### 查询单行数据

```
ls, err := MasterDB.First("select * from book where id = ?", 1)
	if err != nil {
		t.Error(err)
	}
	for f, item := range ls {
		fmt.Println(f, item)
	}
```

### 查询单列数据

```
ls, err := MasterDB.Pluck("select * from book where id = ?", "book_name", 3)
if err != nil {
    t.Error(err)
}
for f, item := range ls {
    fmt.Println(f, item)
}
```

### 增加数据
```
ls, err := MasterDB.Insert("insert into book set book_name=?, book_author=?, book_province=?", "论语", "孔子", "山东")
if err != nil {
    t.Error(err)
}
fmt.Println(ls)
```


### 修改数据
```
ls, err := MasterDB.Update("update book set book_name=? where id=?", "国学-论语", 3)
if err != nil {
    t.Error(err)
}
fmt.Println(ls)
```


### 删除数据
```
ls, err := MasterDB.Delete("delete from book where id = ?", 1)
if err != nil {
    t.Error(err)
}
fmt.Println(ls)
```

