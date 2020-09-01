package main

import (
	"bufio"
	"fmt"
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

// connection handler
func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

// request handler
func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}
