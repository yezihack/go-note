package json_to_path

import (
	"fmt"
	"reflect"
	"testing"
)

var Buf2 = []byte(`{
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

func TestJsonChanger_ToPath(t *testing.T) {
	jc := NewJsonChanger(Buf2)
	list := jc.ToPath()
	for k, v := range list {
		fmt.Println(k, v)
	}
}
func TestRe(t *testing.T) {
	s := []interface{}{
		1, "aa", []int{11}, map[string]string{"a": "bb"}, 0.6, 11111111111111,
	}
	for _, val := range s {
		t := reflect.TypeOf(val).String()
		fmt.Println(t)
	}
}
