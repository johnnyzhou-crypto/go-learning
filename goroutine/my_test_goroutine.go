package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 200)
	}
}

func responseItsSize(url string) {
	fmt.Println("follow the step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("follow the step2: ", url)
	defer response.Body.Close()

	fmt.Println("follow the step3: ", url)
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("final step4: ", len(body))

}

func main() {
	//start a go routine
	go showMsg("from java to go")
	//start another go routine
	go showMsg("from go to gogogo")
	//start a third go routine
	//which will not ensure all go routine can finish on time, once main thread is dead, all go routines dead
	//but if we set enough time sleep, all go routine can be finished on time
	time.Sleep(time.Millisecond * 2000)
	fmt.Println("main end")

	go responseItsSize("https://duoke360.com")
	go responseItsSize("https://bing.com")
	go responseItsSize("https://jd.com")
	time.Sleep(time.Second * 20)
}
