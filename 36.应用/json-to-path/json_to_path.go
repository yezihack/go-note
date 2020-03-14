package json_to_path

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
)

//定义接口
type JsonChanger interface {
	ToPath() map[string]interface{}          //转换成path格式的
	Get(key string) string                   //获取节点数据
	Put(key string, value interface{})       //设置节点数据, 如果存在则覆盖
	Set(key string, value interface{}) error //设置节点数据, 如果存在则设置失败
	SetSeparator(flag string)                //设置分隔符号
}

//定义结构体
type jsonChanger struct {
	data       map[string]interface{} //数据对象
	separator  string                 //分隔符号
	resultData map[string]interface{} //存储返回的数据
	mu         sync.RWMutex
}

//设置节点数据, 如果存在则设置失败
func (j *jsonChanger) Set(key string, value interface{}) error {
	j.mu.Lock()
	defer j.mu.Unlock()
	if !j.check(key) || !j.check(value) {
		return errors.New("key or value only allow string, int, float")
	}
	if _, ok := j.resultData[key]; ok {
		return fmt.Errorf("key: %s is exists", key)
	}
	j.resultData[key] = value
	return nil
}

//获取
func (j *jsonChanger) Get(key string) string {
	j.mu.RLock()
	defer j.mu.RUnlock()
	if val, ok := j.resultData[key]; ok {
		return fmt.Sprint(val)
	}
	return ""
}

//设置节点数据, 如果存在则覆盖
func (j *jsonChanger) Put(key string, value interface{}) error {
	j.mu.Lock()
	defer j.mu.Unlock()
	if !j.check(key) || !j.check(value) {
		return errors.New("key or value only allow string, int, float")
	}
	j.resultData[key] = value
	return nil
}

//初使化一下对象
func NewJsonChanger(data []byte) JsonChanger {
	d := make(map[string]interface{})
	err := json.Unmarshal(data, &d)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return &jsonChanger{
		data:       d,
		separator:  "/",
		resultData: make(map[string]interface{}),
	}
}

//设置分隔符号
func (j *jsonChanger) SetSeparator(flag string) {
	j.separator = flag
}

//转换成路径格式
func (j *jsonChanger) ToPath() map[string]interface{} {
	j.gen(j.data, "")
	return j.resultData
}

//递归实现将json全部分隔出来存入map里
func (j *jsonChanger) gen(kv map[string]interface{}, key string) {
	for k, value := range kv {
		newKey := key + j.separator + k
		if key == "" {
			newKey = k
		}
		if v, ok := value.(map[string]interface{}); ok {
			j.gen(v, newKey)
			continue
		}
		switch value.(type) {
		case map[string]interface{}:
			j.gen(value.(map[string]interface{}), newKey)
			continue
		case []interface{}:
			for ik, iv := range value.([]interface{}) {
				j.gen(iv.(map[string]interface{}), newKey+j.separator+strconv.Itoa(ik))
			}
			continue
		}
		j.resultData[newKey] = value
	}
}

//检查字符
func (j *jsonChanger) check(target interface{}) bool {
	switch target.(type) {
	case string, int, float64:
		return true
	}
	return false
}
