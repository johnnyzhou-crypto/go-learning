package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("types: ", reflect.TypeOf(arg))
	fmt.Println("values: ", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.3241
	reflectNum(num)
}
