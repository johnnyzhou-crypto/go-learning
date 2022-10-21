package main

import (
	"container/list"
	"fmt"
)

func main() {
	//双向链表
	l := list.New()
	l.PushBack(41)
	l.PushBack(32)
	l.PushBack(88)
	//在元素v1前插入3
	v2 := l.PushFront(18)
	l.InsertBefore(3, v2)
	l.PushFront(21)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
}
