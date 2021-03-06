package main

import (
	"fmt"
	"io"
	"net"
	"tcp-proxy-go/util"
)

func main() {
	fmt.Println("proxy listening for connection@", util.ProxyAddress)
	// listen on all interfaces
	ln, err := net.Listen("tcp", util.ProxyAddress)
	util.HandleErr(err)
	// accept connection on port
	proxyConn, err := ln.Accept()
	util.HandleErr(err)
	defer proxyConn.Close()
	// connect to proxy server
	serverConn, err := net.Dial("tcp", util.ServerAddress)
	util.HandleErr(err)
	defer proxyConn.Close()
	sendFile(proxyConn, serverConn)
}

func sendFile(proxyConn, serverConn net.Conn) {
	_, err := io.Copy(serverConn, proxyConn)
	util.HandleErr(err)
}
