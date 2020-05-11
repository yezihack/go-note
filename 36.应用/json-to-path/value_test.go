package json_to_path

import (
	"encoding/json"
	"fmt"
	"testing"
)

var Buf = []byte(`{
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

func TestKv_String(t *testing.T) {
	kv := NewKv()
	if err := json.Unmarshal(Buf, kv.GetData()); err != nil {
		t.Error(err)
	}
	fmt.Println(kv.String("data.list.id"))
}

func TestKv_Strings(t *testing.T) {
	kv := NewKv()
	if err := json.Unmarshal(Buf, kv.GetData()); err != nil {
		t.Error(err)
	}
	fmt.Println(kv.Strings("data.list"))
}
