package method

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	m := make(map[int]int)
	m[1] = 50
	m[2] = 50
	m[3] = 50
	m[4] = 50
	m[5] = 50

	data, err := json.Marshal(m)
	if err != nil {
		t.Error("err:", err)
	}
	fmt.Println(string(data))
}