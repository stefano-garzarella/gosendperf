package main

import (
	"golang.org/x/sys/unix"
)

func (tp *TestParams) TPTestServer() SendResult {
	var res SendResult

	buf := make([]byte, tp.buf_size)

	n, err := unix.Write(tp.addr.GetFD(), buf)
	if err != nil {
		panic(err)
	}

	println("n: ", n)
	return res
}

func (tp *TestParams) TPTestClient() SendResult {
	var res SendResult

	buf := make([]byte, tp.buf_size)

	n, err := unix.Read(tp.addr.GetFD(), buf)
	if err != nil {
		panic(err)
	}

	println("n: ", n)
	return res
}
