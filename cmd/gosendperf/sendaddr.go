package main

type SendAddr interface {
	Connect(address string) error
	Listen(address string) error
	Accept() (SendAddr, error)

	GetFD() int
	Close() error
}
