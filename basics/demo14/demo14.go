package main

import (
	"fmt"
)

// map集合
func main() {
	// 创建map
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["American"]
	// fmt.Println(capital)
	// fmt.Println(ok)
	if ok {
		fmt.Println("American的首都是", capital)
	} else {
		fmt.Println("nil")
	}

	// delete删除map元素
	countryCapitalMap2 := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("原始地图")

	for country := range countryCapitalMap2 {
		fmt.Println(country, "首都是", countryCapitalMap2[country])
	}

	delete(countryCapitalMap2, "France")

	fmt.Println("delete France")

	fmt.Println("删除元素后地图")

	for country := range countryCapitalMap2 {
		fmt.Println(country, "首都是", countryCapitalMap2[country])
	}
}
