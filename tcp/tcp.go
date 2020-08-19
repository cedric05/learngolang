package tcp

import (
	"fmt"
	"net"
)

type Conn struct {
	conn *net.TCPConn
}

func (conn Conn) Read(ch chan []byte) {
	buf := make([]byte, 204800)
	for {
		u, err := conn.conn.Read(buf)
		if err == nil {
			ch <- buf[:u]
			fmt.Println("tcp read", u)
		} else {
			fmt.Println("tcp into error ", err)
			return
		}
	}
}

func (conn Conn) Write(ch chan []byte) {
	for a := range ch {
		fmt.Println("tcp write", len(a))
		_, error := conn.conn.Write(a)
		if error != nil {
			fmt.Println("tcp write ran into error")
		}
	}
	// TODO close channel
}

func New(conn *net.TCPConn) Conn {
	return Conn{conn}
}
