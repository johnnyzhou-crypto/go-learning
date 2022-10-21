package main

import (
	"fmt"
	"sort"
)

type testSlides []map[string]float64

func (l testSlides) Len() int {
	return len(l)
}

func (l testSlides) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l testSlides) Less(i, j int) bool {
	return l[i]["a"] < l[j]["a"]
}

//func Marchel(l interface{}) ([]byte, error) {
//
//	return []byte, nil
//}

func main() {
	ls := testSlides{
		{"a": 8, "b": 5},
		{"a": 4, "b": 1},
		{"a": 7, "b": 9},
	}

	fmt.Printf("%v\n", ls)
	sort.Sort(ls)
	fmt.Printf("%v\n", ls)
}
