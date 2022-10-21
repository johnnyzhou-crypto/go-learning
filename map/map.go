package main

import (
	"fmt"
	"reflect"
)

func main() {

	var m map[string]int
	m = map[string]int{"golang": 1, "java": 2}
	fmt.Println(m)

	m2 := map[string]map[string]int{}
	innerm := map[string]int{}
	innerm["im"] = 123
	m2["m"] = innerm
	fmt.Println(m2)

	var m3 map[string]int
	m3 = map[string]int{"mon": 1, "tue": 2}
	fmt.Println(m3)
	if _, ok := m3["tue"]; !ok {
		fmt.Println("The key does not exist")
	}
	delete(m3, "tue")
	fmt.Println(m3, len(m3))

	//delete this m3
	m3 = nil

	m5 := map[string]int{"a": 1, "b": 2, "c": 3}
	m6 := map[string]int{"a": 1, "c": 3, "b": 2}
	//方法一
	//其中方法一用到了反射，效率相对比较低，相差大约10倍
	fmt.Println(reflect.DeepEqual(m5, m6))
	fmt.Println(compareMap(m5, m6))

}

func compareMap(m1, m2 map[string]int) bool {
	if len(m1) == len(m2) {
		for k1, v1 := range m1 {
			if v2, ok := m2[k1]; ok {
				if v1 != v2 {
					return false
				}
			} else {
				return false
			}
		}
		return true
	}
	return false
}
