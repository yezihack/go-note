package json_to_path

import (
	"fmt"
)

func main() {
	//buf := []byte(`{"a": {"b": 11}}`)
	buf := []byte(`{
  "code": 200,
  "message": "success",
  "data": {
    "page": 1,
    "pageCount": 10,
    "totalCount": 93,
    "list":
      {
        "id": 68188,
        "name": "火由者南会别",
        "guard_uuid": "Dd6aFE6c-ffC1-Ed64-24e7-D9bd6c1eAE1e",
        "app_name": "向话治价",
        "created_at": "2017-12-19",
        "file_format_name": "文本",
        "mode": "用户",
        "status": "已冻结"
      },
	"line":[
      {
        "id": 68188,
        "name": "火由者南会别",
        "guard_uuid": "Dd6aFE6c-ffC1-Ed64-24e7-D9bd6c1eAE1e",
        "app_name": "向话治价",
        "created_at": "2017-12-19",
        "file_format_name": "文本",
        "mode": "用户",
        "status": "已冻结"
      },
		{
        "id": 68188,
        "name": "火由者南会别",
        "guard_uuid": "Dd6aFE6c-ffC1-Ed64-24e7-D9bd6c1eAE1e",
        "app_name": "向话治价",
        "created_at": "2017-12-19",
        "file_format_name": "文本",
        "mode": "用户",
        "status": "已冻结"
      }
	]
  }
}`)
	jp := NewJsonChanger(buf)
	data := jp.ToPath()
	fmt.Println(data)
	return

	//kv := NewKv()
	//kv.SetDelim("/")
	//if err := json.Unmarshal(buf, kv.GetData()); err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(kv.String("a/b")) //输出11的结果.
	//
	//mKv := make(map[string]interface{})
	//err := json.Unmarshal(buf, &mKv)
	//if err != nil {
	//	panic(err)
	//}
	//
	//gen(mKv, "")
	//fmt.Println(data)
	//for k, v := range data {
	//	fmt.Printf("key:%s, value:%v\n", k, v)
	//}
}
