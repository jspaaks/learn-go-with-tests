package main

import (
	"fmt"
	"github.com/jspaaks/learn-go-with-tests/helloworld/pkg/helloworld"
)

func main() {
	var s string = helloworld.SayHello("John", "")
	fmt.Println(s)
}
