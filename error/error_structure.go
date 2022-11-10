package main

import (
	"errors"
	"fmt"
	"log"
	"my-test-service/error/example"
	_ "my-test-service/error/example"
)

type badValueError struct {
	key   string
	value interface{}
}

func (bv *badValueError) Error() string {
	return fmt.Sprintf("bad value %v for key %s", bv.value, bv.key)
}

func testErr() error {
	v := 3
	k := "test"

	return &badValueError{key: k, value: v}
}

func main() {
	err := testErr()

	var e *badValueError

	if errors.As(err, &e) {
		log.Printf("%v", err)
	} else {
		log.Println("wrong")
	}

	//errorf
	simpleError := errors.New("a simple error")
	simpleError2 := fmt.Errorf("an error from a %s string", simpleError)
	fmt.Println(simpleError2)

	err2 := errors.New("an_error")
	wrappedError := example.Wrapf(err2, "error %s", "[customised Error]")
	fmt.Println(wrappedError)

}
