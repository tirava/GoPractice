package main

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
)

func main() {
	conf, err := yaml.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Error opening: ", err)
	}

	fmt.Println(conf.GetBool("enabled"))
	fmt.Println(conf.Get("path"))
}
