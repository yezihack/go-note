package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	json_str := []byte(`{"age":1}`)
	var value map[string]interface{}
	json.Unmarshal(json_str, &value)
	age := value["age"]
	fmt.Println(reflect.TypeOf(age)) //float64 所有的整形都是float64
}
