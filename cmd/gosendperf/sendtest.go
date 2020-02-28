package main

type TestParams struct {
	addr     SendAddr
	buf_size uint64
	len      uint64
}

type SendResult struct {
	time  uint64
	bytes uint64
}

type SendTest interface {
	server() SendResult
	client() SendResult
}
