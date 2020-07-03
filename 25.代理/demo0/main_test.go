package main

import (

	"fmt"
	"testing"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Server string `json:"server"`
}

func TestPing(t *testing.T) {
	uri := "/ping"
	body := Get(uri, router)
	fmt.Print(string(body))
	// resp := Response{}
	// if err := json.Unmarshal(body, &resp); err != nil {
	// 	t.Fatal(err)
	// }
	// if resp.Code != 200 {
	// 	t.Errorf("got:%d, want:%d\n", resp.Code, 200)
	// }
	// println(resp.Code, resp.Msg, resp.Server)
}