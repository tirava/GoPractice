package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type configuration struct {
	Enabled bool
	Path    string
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Error opening: ", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatal("Error decoding: ", err)
	}

	fmt.Println(conf.Path, conf.Enabled)
}
