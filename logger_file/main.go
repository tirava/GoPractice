package main

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.OpenFile("./log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	defer logfile.Close()

	logger := log.New(logfile, "testik ", log.LstdFlags|log.Lshortfile)

	logger.Println("Simple message.")
	logger.Fatalln("Fatal error.")
	logger.Println("End message.")
}