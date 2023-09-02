package es6

import (
	"encoding/json"
)

type JSObject struct {
	data map[string]interface{}
}

func NewJSObject() *JSObject {
	return &JSObject{
		data: make(map[string]interface{}),
	}
}

func (o *JSObject) Set(key string, value interface{}) {
	o.data[key] = value
}

func (o *JSObject) Get(key string) interface{} {
	return o.data[key]
}

func (o *JSObject) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(o.data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// func main() {
// 	obj := NewJSObject()
// 	obj.Set("name", "John")
// 	obj.Set("age", 30)
// 	obj.Set("isStudent", false)

// 	jsonStr, err := obj.ToJSON()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println(jsonStr)
// }
