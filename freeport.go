package freeport

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
)

type Options struct {
	Address string
	Min     int
	Max     int
}

// MakeOptions makes options to get non-privileged port [1024:49151]. It will use
// env var PORT, if set, as min value.
func MakeOptions(isAllInterfaces bool) (*Options, error) {
	min := 0

	s := os.Getenv("PORT")
	if s != "" {
		min, _ = strconv.Atoi(s)
	}

	if min == 0 {
		min = 1024
	}

	address := "127.0.0.1"
	if isAllInterfaces {
		address = ""
	}

	// limit to non-privileged ports, sys apps should not be using FreePort
	return &Options{Address: address, Min: min, Max: 49151}, nil
}

var ErrPortNotFound = errors.New("port not found")

// GetFreePortEx asks the kernel for a free open port that is ready to use. If options is nil, then
// default options are used.
func GetFreePortEx(options *Options) (int, error) {
	var err error

	if options == nil {
		options, err = MakeOptions(false)
		if err != nil {
			return 0, err
		}
	}

	for port := options.Min; port <= options.Max; port++ {
		pingAddr := fmt.Sprintf("%s:%d", options.Address, port)
		addr, err := net.ResolveTCPAddr("tcp", pingAddr)
		if err != nil {
			continue
		}
		l, err := net.ListenTCP("tcp", addr)
		if err != nil {
			continue
		}

		defer l.Close()
		return l.Addr().(*net.TCPAddr).Port, nil
	}

	return 0, ErrPortNotFound
}

// GetFreePort gets non-privileged open port that is ready to use.
func GetFreePort() (int, error) {
	return GetFreePortEx(nil)
}

// MustGetFreePort calls GetFreePort and panics on error
func MustGetFreePort() int {
	port, err := GetFreePortEx(nil)
	if err != nil {
		panic(err)
	}
	return port
}

// GetFreePorts gets an array of non-privileged open ports that are ready to use.
func GetFreePorts(count int) ([]int, error) {
	ports := make([]int, count)
	options, err := MakeOptions(false)
	if err != nil {
		return nil, err
	}

	for i := 0; i < count; i++ {
		port, err := GetFreePortEx(options)
		if err != nil && err != ErrPortNotFound {
			return nil, err
		}
		ports[i] = port
		options.Min = port + 1
	}
	return ports, nil
}
