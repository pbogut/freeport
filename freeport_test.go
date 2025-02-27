package main

import (
	"net"
	"strconv"
	"testing"
)

func TestGetFreePort(t *testing.T) {
	port, err := GetFreePort()
	if err != nil {
		t.Error(err)
	}
	if port == 0 {
		t.Error("port == 0")
	}

	// Try to listen on the port
	l, err := net.Listen("tcp", "localhost"+":"+strconv.Itoa(port))
	if err != nil {
		t.Error(err)
	}
	defer l.Close()
}

func TestGetFreePortEx(t *testing.T) {
	opts, err := MakeOptions(false)
	if err != nil {
		t.Error(err)
	}

	opts.Min = 10000
	port, err := GetFreePortEx(opts)
	if err != nil {
		t.Error(err)
	}
	if port == 0 {
		t.Error("port == 0")
	}

	if port < opts.Min {
		t.Error("expected port to be > 10000")
	}

	// Try to listen on the port
	l, err := net.Listen("tcp", "localhost"+":"+strconv.Itoa(port))
	if err != nil {
		t.Error(err)
	}
	defer l.Close()
}

func TestGetFreePorts(t *testing.T) {
	count := 3
	ports, err := GetFreePorts(count)
	if err != nil {
		t.Error(err)
	}
	if len(ports) == 0 {
		t.Error("len(ports) == 0")
	}
	for _, port := range ports {
		if port == 0 {
			t.Error("port == 0")
		}

		// Try to listen on the port
		l, err := net.Listen("tcp", "localhost"+":"+strconv.Itoa(port))
		if err != nil {
			t.Error(err)
		}
		defer l.Close()
	}
}
