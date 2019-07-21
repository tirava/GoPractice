package main

import (
	"fmt"
	"gopkg.in/gcfg.v1"
	"log"
)

func main() {
	conf := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}
	err := gcfg.ReadFileInto(&conf, "config.ini")
	if err != nil {
		log.Fatal("Error opening: ", err)
	}

	fmt.Println(conf.Section.Enabled)
	fmt.Println(conf.Section.Path)
}
