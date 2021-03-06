package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	nl, err := net.Listen("tcp", "192.168.1.123:8880") //(network,addr[ip:port](socket)})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //stop
	}
	conn, err := nl.Accept()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //stop
	}
	remoteAddress := conn.RemoteAddr().String()
	fmt.Println(remoteAddress)
	conn.Write([]byte("Welcome to coodingbootcamp"))
	conn.Close()
	nl.Close()

}
