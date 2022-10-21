package main

import (
	"errors"
	"fmt"
)

func main() {

	b := new(int)
	fmt.Printf("%T\n", b)

	var p *[]int = new([]int)

	fmt.Printf("p: %v\n", p)

	v := make([]int, 10)

	fmt.Printf("v: %v\n", v)

	//c := make([]byte, 10)
	//
	//var by bytes.Buffer

}

/**
带error返回
*/
func test(s string) (string, error) {
	if s == "" {
		err := errors.New("null string")
		return "", err
	} else {
		return s, nil
	}
}
