package ch3

import (
	"io"
	"net"
	"testing"
)

func TestDial(t *testing.T) {
	// 랜덤 포트에 리스너 생성
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		defer func() {
			t.Logf("closing listener")
			done <- struct{}{}
		}()

		for {
			t.Logf("accepting connection")
			conn, err := listener.Accept()
			if err != nil {
				t.Log("err:", err)
				return
			}
  
			go func(c net.Conn) {
				defer func() {
					c.Close()
					t.Logf("closing connection")
					done <- struct{}{}
				}()

				buf := make([]byte, 1024)
				for {
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						t.Logf("===EOF===")
						return
					}
					t.Logf("received: %q", buf[:n])
				}
			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	conn.Close()
	<-done

	listener.Close()
	<-done
}
