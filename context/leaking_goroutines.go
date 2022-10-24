package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

// Once the caller is done with the generator (when it breaks the loop), the goroutine will run forever executing the infinite loop. Our code will leak a goroutine.
func generate() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
		}
	}()
	return ch
}

/*
*
We can avoid the problem by signaling the internal goroutine with a stop channel but there is a better solution: cancellable contexts.
The generator can select on a context’s Done channel and once the context is done, the internal goroutine can be cancelled.
*/
func avoid_leaking_generate(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				//immediately return to avoid leaking
				return
			case ch <- n:
				n++
				//default:
				//	fmt.Println("default...")
			}
		}
	}()
	return ch
}

func g1() {
	//The generator above starts a goroutine with an infinite loop, but the caller consumes the values until n is equal to 5
	for i := range generate() {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
}

func g2() {
	ctx, cancel := context.WithCancel(context.Background())
	// make sure all paths cancel the context to avoid context leak
	defer cancel()
	for i := range avoid_leaking_generate(ctx) {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
}

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	select {}

	g2()

}
