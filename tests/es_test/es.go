package main

import (
	"encoding/json"
	"fmt"
	"fyoukuApi/services/es"
	"strconv"
)

func addES() {
	for i := 0; i < 50; i++ {
		idStr := strconv.Itoa(i)
		body := map[string]interface{}{
			"id":    1,
			"title": "小黑子没有素质:" + "小黑子" + idStr + "号",
		}
		if ok := es.EsAdd("fyouku_demo", "user-"+idStr, body); !ok {
			fmt.Println("写入失败")
			return
		}
		fmt.Println("写入成功")
	}
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

func searchES() {
	sort := []map[string]string{map[string]string{"id": "desc"}}
	query := map[string]interface{}{
		"bool": map[string]interface{}{
			"must": map[string]interface{}{
				"term": map[string]interface{}{
					"id": 1,
				},
			},
		},
	}
	res := es.EsSearch("fyouku_demo", query, 0, 10, sort)
	fmt.Println("total:", res.Total)
	fmt.Println("hits:", res.Hits)

	var resList []resData
	for _, v := range res.Hits {
		var data resData
		err := json.Unmarshal([]byte(v.Source), &data)
		if err != nil {
			fmt.Println("err:", err.Error())
			return
		}
		resList = append(resList, data)
	}
	fmt.Println("data:", resList)
}

type resData struct {
	ID    int
	Title string
}

func main() {
	addES()
	updateES()
	searchES()
	deleteES()

}
