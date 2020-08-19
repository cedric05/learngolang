package main

import (
	"fmt"
	"net"
	"os"

	"./tcp"
	"./unix"
)

type IOConn interface {
	Read(ch chan byte)
	Write(ch chan byte)
	Connect(addr string)
}

const outbound = "localhost:2375"
const ip = "127.0.0.1"
const port = "2375"

func handle(con *net.UnixConn) {
	unixToTCP := make(chan []byte)
	tcpToUnix := make(chan []byte)
	unixCon := unix.New(con)

	remoteaddr := net.TCPAddr{
		[]byte{127, 0, 0, 1},
		2735,
		"",
	}
	// handle tcperror
	conn, tcperror := net.DialTCP("tcp", nil, &remoteaddr)
	if tcperror != nil {
		fmt.Println(" error happened %v", tcperror)
		return
	}
	tcpconn := tcp.New(conn)

	go unixCon.Read(tcpToUnix)
	go tcpconn.Write(tcpToUnix)

	go tcpconn.Read(unixToTCP)
	go unixCon.Write(unixToTCP)
	// defer close(tcpToUnix)
	// defer close(unixToTCP)

}

func main() {
	addr, err := net.ResolveUnixAddr("unix", "/tmp/foobar")
	if err != nil {
		fmt.Printf("Failed to resolve: %v\n", err)
		os.Exit(1)
	}
	list, err := net.ListenUnix("unix", addr)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		os.Exit(1)
	}
	for {
		conn, _ := list.AcceptUnix()
		go handle(conn)
	}
}
