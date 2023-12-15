package main

import (
	"fmt"
	"fyoukuApi/services/es"
)

func addES() {
	body := map[string]interface{}{
		"id":    1,
		"title": "我是练习时长两年半的蔡徐坤",
	}
	if ok := es.EsAdd("fyouku_demo", "user-1", body); !ok {
		fmt.Println("写入失败")
		return
	}
	fmt.Println("写入成功")
}

func updateES() {
	body := map[string]interface{}{
		"id":    1,
		"title": "小黑子没有素质",
	}
	ok := es.EsEdit("fyouku_demo", "user-1", body)
	if !ok {
		fmt.Println("更新失败")
		return
	}
	fmt.Println("更新成功")
}

func deleteES() {

	ok := es.EsDelete("fyouku_demo", "user-1")
	if !ok {
		fmt.Println("删除失败")
		return
	}
	fmt.Println("删除成功")
}

func main() {
	addES()
	updateES()
	deleteES()

}
