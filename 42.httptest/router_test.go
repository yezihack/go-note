package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// TestOnGetStringRequest 测试以Get方式获取一段字符串的接口
func TestOnGetStringRequest(t *testing.T) {

	// 初始化请求地址
	uri := "/getString"

	// 发起Get请求
	body := Get(uri, router)
	fmt.Printf("response:%v\n", string(body))

	// 判断响应是否与预期一致
	if string(body) != "success" {
		t.Errorf("响应字符串不符，body:%v\n",string(body))
	}
}
// TestOnLoginRequestForForm 测试以表单形式传递参数的登录接口
func TestOnLoginRequestForForm(t *testing.T) {
	// 初始化请求地址和请求参数
	uri := "/loginForm"

	param := make(map[string]string)
	param["username"] = "valiben"
	param["password"] = "123"
	param["age"] = "1"

	// 发起post请求，以表单形式传递参数
	body := PostForm(uri, param, router)
	fmt.Printf("response:%v\n", string(body))

	// 解析响应，判断响应是否与预期一致
	response := &LoginResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Errorf("解析响应出错，err:%v\n",err)
	}
	if response.Data.Username != "valiben" {
		t.Errorf("响应数据不符，username:%v\n",response.Data.Username)
	}
}
// TestOnLoginRequestForJson 测试以Json形式传递参数的登录接口
func TestOnLoginRequestForJson(t *testing.T) {
	// 初始化请求地址和请求参数
	uri := "/loginJson"

	param := make(map[string]interface{})
	param["username"] = "valiben"
	param["password"] = "123"
	param["age"] = 1

	// 发起post请求，以Json形式传递参数
	body := PostJson(uri, param, router)
	fmt.Printf("response:%v\n", string(body))

	// 解析响应，判断响应是否与预期一致
	response := &LoginResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Errorf("解析响应出错，err:%v\n",err)
	}
	if response.Data.Username != "valiben" {
		t.Errorf("响应数据不符，username:%v\n",response.Data.Username)
	}
}