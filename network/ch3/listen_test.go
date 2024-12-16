package ch3

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	defer func() { listener.Close() }()

	t.Logf("bound to %q", listener.Addr())
  }
