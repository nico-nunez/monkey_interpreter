package main

import (
	"fmt"
)

func main() {
	foo := "testing"
	bar := &foo
	foobar := *bar
	fmt.Println(foo)
	fmt.Println(bar)
	fmt.Println(foobar)
}
