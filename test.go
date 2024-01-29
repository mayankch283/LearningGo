package main

import "fmt"

func cray() {
	fmt.Println(foo())
}

func foo() string {
	return "bar"
}
