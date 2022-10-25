package main

import (
	"errors"
	"fmt"
)

func main() {
	errA := errors.New("this is error A; ")
	errB := fmt.Errorf("this is error B; %w", errA)
	errC := fmt.Errorf("this is error C; %w", errB)
	fmt.Println(errC)
	errB = errors.Unwrap(errC)
	errA = errors.Unwrap(errB)
	fmt.Println(errA)
}
