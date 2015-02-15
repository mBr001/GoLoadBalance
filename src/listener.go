package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
)

func startServer(listenPort int, backends *Backends) {
	port := strconv.Itoa(listenPort)
	fmt.Println("Starting server on port ", port)
	addr, _ := net.ResolveTCPAddr("tcp", ":"+port)

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println("Error occured accepting a connection", err.Error())
		}

		go handleConnection(con, backends.NextAddress())
	}

}

func handleConnection(cli_conn net.Conn, srv_addr string) {
	srv_conn, err := net.Dial("tcp", srv_addr)
	if err != nil {
		fmt.Println("Could not connect to server, connection dropping")
		return
	}

	go io.Copy(cli_conn, srv_conn)
	io.Copy(srv_conn, cli_conn)
}
