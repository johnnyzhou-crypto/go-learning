package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	m := map[string]Person{}
	p := Person{Name: "Johnny", Age: 18}
	m["m1"] = p
	fmt.Println(m)
}