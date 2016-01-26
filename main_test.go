package hello_test

import (
	"fmt"

	"github.com/mattes/migrate/pipe"
	"github.com/raizyr/wercker_go_test"
)

func ExampleHelloWorld() {
	fmt.Println(hello.Hello())
	// Output: hello world
}

func ExamplePipe() {
	p := pipe.New()
	go func() { p <- "ping" }()
	msg := <-p
	fmt.Println(msg)
	// Output: ping
}
