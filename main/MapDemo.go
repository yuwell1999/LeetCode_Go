package main

import (
	"fmt"
)

func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 100
	scoreMap["李四"] = 200
	scoreMap["李四"] = 300
	scoreMap["王五"] = 400
	// 也支持声明的时候添加元素
	//scoreMap :=map[string]int{
	//	"张三":100,
	//	"李四":300,
	//}
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["李四"])
	fmt.Println(len(scoreMap))

	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(ok)
	} else {
		_ = v
		fmt.Println("查无此人")
	}

	// map的遍历

	// for-range
	fmt.Println("删除键值对前")
	for key, value := range scoreMap {
		fmt.Println(key, value)
	}
	//// 只想遍历key
	//for key := range scoreMap{
	//	fmt.Println(key)
	//}

	// 删除键值对
	delete(scoreMap, "张三")
	fmt.Println("删除键值对后")
	for key := range scoreMap {
		fmt.Println(key, scoreMap[key])
	}

	mapInit := map[string]string{
		"小李": "湖南",
		"小红": "广东",
	}
	_ = mapInit
	mapTemp := make(map[string]string, 10)
	mapTemp["小明"] = "北京"
	mapTemp["小王"] = "河北"

	v1, ok := mapTemp["小明"]
	fmt.Println(v1, ok)

}
