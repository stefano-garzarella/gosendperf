package main

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

type VsockAddr struct {
	fd int
	svm unix.SockaddrVM
}

func (va *VsockAddr) parse(addr string) (error) {
	fmt.Println(addr)

	parsed := strings.Split(addr, ":")
	if (len(parsed) != 2) {
		return fmt.Errorf("Expected cid:port")
	}

	tmp, err := strconv.ParseUint(parsed[0], 10, 32)
	if err != nil {
		return err
	}
	va.svm.CID = uint32(tmp)

	tmp, err = strconv.ParseUint(parsed[1], 10, 32)
	if err != nil {
		return err
	}
	va.svm.Port = uint32(tmp)

	return nil
}

func (va *VsockAddr) Connect(addr string) (error) {
	var err error

	err = va.parse(addr)
	if err != nil {
		return err
	}

	va.fd, err = unix.Socket(unix.AF_VSOCK, unix.SOCK_STREAM, 0)
	if err != nil {
		return err
	}

	return nil
}

func (va *VsockAddr) Listen(addr string) (error) {
	return nil
}

func (va *VsockAddr) GetFD() (int) {
	return va.fd
}

func (va *VsockAddr) Accept() (SendAddr, error) {
	return nil, nil
}
func (va *VsockAddr) Close() (error) {
	return nil
}