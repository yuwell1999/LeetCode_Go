package main

import (
	"fmt"
	"sort"
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

	//v2, ok := mapTemp["小王"]
	//if ok{
	//	fmt.Println(v2)
	//}

	// 相当于
	if v2, ok := mapTemp["小王"]; ok {
		fmt.Println(v2)
	}

	fmt.Println(len(mapInit))
	fmt.Println(len(mapTemp) + len(mapInit) + len(scoreMap))

	printByOrder()
}

func printByOrder() {
	map1 := make(map[int]string, 5)

	map1[1] = "www.topgoer.com"
	map1[2] = "rpc.topgoer.com"
	map1[5] = "ceshi"
	map1[3] = "xiaohong"
	map1[4] = "xiaohuang"

	sli := []int{}
	for k := range map1 {
		sli = append(sli, k)
	}
	sort.Ints(sli)
	for i := 0; i < len(map1); i++ {
		fmt.Print(map1[sli[i]] + " ")
	}
}
