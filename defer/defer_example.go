package main

import "fmt"

func main() {
	deferFunc()
}

func deferFunc() {
	//注意，defer的执行顺序遵循后进先出的原则，后面的defer语句将会被先执行，因此执行的顺序应该是从下至上
	//先进先出
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")

}
