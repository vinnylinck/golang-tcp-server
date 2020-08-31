package main

import (
	"log"
	"net"
)

// main function - entry point
func main() {
	// retrieving env vars
	port := getPort()

	// creating / starting tcp listener
	tcpli, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer tcpli.Close()

	// reading buffer from connections
	for {
		conn, err := tcpli.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handle(conn)
	}
}
