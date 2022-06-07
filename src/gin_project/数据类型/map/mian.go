package main

import (
	"fmt"
	"sort"
)

func createMap() {
	aMap := make(map[string]int, 8)
	aMap["小明"] = 3
	aMap["小李"] = 4
	fmt.Println(aMap)
	fmt.Println(aMap["小明"])
	fmt.Printf("type: %T, value: %[1]v", aMap)
}

func createMap2() {
	aMap := map[string]int{
		"张三": 1,
		"李四": 2,
	}
	fmt.Println(aMap)
	fmt.Println(aMap["张三"])
	fmt.Printf("type: %T, value: %[1]v", aMap)
}

func createMultiMap() {
	aMap := make([]map[string]string, 2)
	for k, v := range aMap {
		fmt.Printf("index:%d value:%v\n", k, v)
	}
	fmt.Println("----赋值后----")
	aMap[0] = make(map[string]string, 3)
	aMap[0]["name"] = "张三"
	aMap[1] = make(map[string]string, 3)
	aMap[1]["age"] = "李四"
	for k, v := range aMap {
		fmt.Printf("index:%d value:%v\n", k, v)
	}
	fmt.Printf("type: %T, len: %v,value: %[1]v", aMap, len(aMap))

}

func isMapKey() {
	aMap := map[string]int{
		"张三": 1,
		"李四": 2,
	}
	name, isKey := aMap["张三"]
	if isKey {
		fmt.Println(name)
	} else {
		fmt.Println("key 不存在")
	}
}

func deleteMap() {
	aMap := make(map[string]int, 2)
	aMap["小明"] = 3
	aMap["小李"] = 4
	delete(aMap, "小明")
	for k, v := range aMap {
		fmt.Printf("key:%v value:%v\n", k, v)
	}
}

func forMap() {
	aMap := make(map[string]int, 2)
	aMap["小明"] = 3
	aMap["小李"] = 4
	for k, v := range aMap {
		fmt.Printf("key:%v value:%v\n", k, v)
	}
}

func forMapKey() {
	aMap := make(map[string]int, 2)
	aMap["小明"] = 3
	aMap["小李"] = 4
	for k := range aMap {
		fmt.Printf("key:%v \n", k)
	}
}

func forSortMapKey() {
	aMap := make(map[string]int, 20)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		aMap[key] = i
	}
	fmt.Println(aMap)

	var keys = make([]string, 0, 50)
	for key := range aMap {
		keys = append(keys, key)
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key, ":", aMap[key])
	}
}

func main() {
	forSortMapKey()
}
