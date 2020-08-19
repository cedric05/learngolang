package unix

import (
	"fmt"
	"net"
)

type Conn struct {
	conn *net.UnixConn
}

func (conn Conn) Read(ch chan []byte) {
	buf := make([]byte, 204800)
	for {
		u, _, err := conn.conn.ReadFromUnix(buf)
		if err == nil {
			fmt.Println("unix read", u)
			ch <- buf[:u]
		} else {
			fmt.Println("unix ran into error ", err)
			return
		}
	}
}

func (conn Conn) Write(ch chan []byte) {
	for a := range ch {
		fmt.Println("unix write ", len(a))
		_, error := conn.conn.Write(a)
		if error != nil {
			fmt.Println("unix write ran into error")
		}
	}
	// TODO close channel
}

func New(conn *net.UnixConn) Conn {
	return Conn{conn}
}
