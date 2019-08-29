package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1902")
	if err != nil {
		panic("failed to conect localhost:1902")
	}
	defer conn.Close()

	logger := log.New(conn, "testik ", log.LstdFlags|log.Lshortfile)

	logger.Println("Simple message.")
	logger.Panicln("Panicccc.")
}