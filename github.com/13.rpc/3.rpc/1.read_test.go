package __rpc

import (
	"net"
	"sync"
	"testing"
)

func TestSession_ReadWriter(t *testing.T) {
	// 定义地址
	addr := "127.0.0.1:8000"
	my_data := "hello"

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}

		conn, _ := lis.Accept()
		s := Session{conn: conn}
		err = s.Write([]byte, )
	}()
}
