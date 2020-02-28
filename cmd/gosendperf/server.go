package main

import (
	"fmt"
)

func server() {
	var tp TestParams

	fmt.Println("I'm the server")

	tp.addr = new(VsockAddr)

	err := tp.addr.Connect("1:1234")
	if err != nil {
		panic(err)
	}
}
