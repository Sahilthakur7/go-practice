package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	val := scanner.Scan()
	fmt.Println(val)
	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
	}

	defer conn.Close()
}