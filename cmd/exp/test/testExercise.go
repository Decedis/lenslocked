package main

import (
	"errors"
	"fmt"
)

func main() {
	err := B()
	// TODO: Determine if the `err` variable is an `ErrNotFound`
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Error is ErrNotFound")
	} else {
		fmt.Println("Error is not ErrNotFound")
	}
}

var ErrNotFound = errors.New("not found")

func A() error {
	return ErrNotFound
}

func B() error {
	err := A()
	if err != nil {
		return fmt.Errorf("b: %w", err)
	}
	return nil
}
