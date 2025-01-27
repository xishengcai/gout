package gout

import (
	"bytes"
	"io"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xishengcai/gout/core"
)

func testTcpSocket(out *bytes.Buffer, quit chan bool, t *testing.T) (addr string) {
	addr = core.GetNoPortExists()

	addr = ":" + addr
	go func() {
		defer close(quit)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			t.Errorf("%v\n", err)
			return
		}
		defer listener.Close()

		conn, err := listener.Accept()
		if err != nil {
			t.Errorf("%v\n", err)
			return
		}

		defer conn.Close()
		if _, err = io.Copy(out, conn); err != nil {
			t.Errorf("%v\n", err)
			return
		}
	}()

	return addr

}

func Test_Use_Chunked(t *testing.T) {
	var out bytes.Buffer
	quit := make(chan bool)

	addr := testTcpSocket(&out, quit, t)
	time.Sleep(time.Second / 100) //等待服务起好

	// 这里超时返回错误, 原因tcp服务没有构造http返回报文
	assert.Error(t, POST(addr).SetTimeout(time.Second/100).Chunked().SetBody("11111111111").Do())
	<-quit
	//time.Sleep(time.Second)

	assert.NotEqual(t, bytes.Index(out.Bytes(), []byte("Transfer-Encoding: chunked")), -1)
}
