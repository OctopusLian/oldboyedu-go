package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name" zhoulin:"嘿嘿嘿"`
	Score int    `json:"score" zhoulin:"哈哈哈"`
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)

	fmt.Println(t.Name(), t.Kind()) // student struct

	// 通过for循环遍历结构体的所有字段信息
	fmt.Println(t.NumField()) // 2
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("zhoulin"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}
