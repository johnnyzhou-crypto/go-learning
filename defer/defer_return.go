package main

import "fmt"

func main() {

	fmt.Println(deferReturn())
}

func deferReturn() (i int) {

	defer func() {
		i += 3
	}()

	i = 3

	return
}
