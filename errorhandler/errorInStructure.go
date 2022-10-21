package errorhandler

import (
	"errors"
	"fmt"
	"log"
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
}
