package main

import (
	"fmt"

	"github.com/real-qgl/go-helper"
)

func main() {
	inp1 := []string{"a", "b", "c", "d", "e"}
	val1 := "a"
	fmt.Printf("helper.InArray(%s, slice1): %t\n", val1, helper.InArray(val1, inp1))
	val2 := "f"
	fmt.Printf("helper.InArray(%s, slice1): %t\n", val2, helper.InArray(val2, inp1))
}
