// +build !linux

package main

import (
	"fmt"
)

type VsockAddr struct {
}

func (va *VsockAddr) parse(addr string) error {
	return fmt.Errorf("AF_VSOCK not supported!")
}

func (va *VsockAddr) Connect(addr string) error {
	return fmt.Errorf("AF_VSOCK not supported!")
}

func (va *VsockAddr) Listen(addr string) error {
	return fmt.Errorf("AF_VSOCK not supported!")
}

func (va *VsockAddr) GetFD() int {
	return -1
}

func (va *VsockAddr) Accept() (SendAddr, error) {
	return nil, fmt.Errorf("AF_VSOCK not supported!")
}
func (va *VsockAddr) Close() error {
	return fmt.Errorf("AF_VSOCK not supported!")
}
