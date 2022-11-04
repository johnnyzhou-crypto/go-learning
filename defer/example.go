package main

import "fmt"

func main() {
	fmt.Printf("f1() num is: %d\n", f1(0))
	fmt.Printf("f2() num is: %d\n", f2(0))
}

// wrong example
func f1(num int) int {
	defer func() {
		num++
	}()
	return num
}

// wrong example
func f2(num int) int {
	i := 3
	defer func() {
		i = i + num
	}()
	return i
}
