package main

import (
	"encoding/json"
	"fmt"

)

type A struct {
	Num int
}
type B struct {
	Num int
}

type Monster struct{
	Name string `json:"name"` // `json:"name"` 就是 struct tag
	Age int `json:"age"`
	Skill string `json:"skill"`
	Hobby []string `json:"hobby"`
	MA   map[string]string `json:"ma"`
}
func main() {
	var a A
	var b B
	a = A(b) // ? 可以转换，但是有要求，就是结构体的的字段要完全一样(包括:名字、个数和类型！)
	fmt.Println(a, b)

	//1. 创建一个Monster变量
	var h []string
	m := map[string]string{"1":"1"}
	monster := Monster{"牛魔王", 500, "芭蕉扇~",append(h,"yyj"),m}

	//2. 将monster变量序列化为 json格式字串
	//   json.Marshal 函数中使用反射，这个讲解反射时，我会详细介绍
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误 ", err)
	}
	fmt.Println("jsonStr", string(jsonStr))
	res := []byte(string(jsonStr))
	fmt.Println(string(res))

}