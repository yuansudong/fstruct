package main

import (
	"log"

	"github.com/yuansudong/fstruct"
)

const (
	_Tag = "example"
)

// Student 学生
type Student struct {
	// Name 名称
	Name string `example:"name"`
	// Age 年龄
	Age int `example:"age"`
	// Address 地址
	Address string `example:"address"`
	// School 学校
	School string `example:"school"`
	// Score 分值
	Score string `example:"score"`
	// Ignore 测试忽略字段
	Ignore string `example:"-"`
}

func main() {
	log.Println("hello,world")
	stu := new(Student)
	if err := fstruct.FillFromMap(
		stu,
		map[string]interface{}{
			"school":  "嘻哈学校",
			"age":     28,
			"address": "地爆天星",
			"name":    "漩涡鸣人",
			"score":   79.8,
		},
		_Tag,
	); err != nil {
		log.Println(err.Error())
		return
	}
	log.Printf("%+v\n", stu)
}
