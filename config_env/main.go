package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println(os.Getenv("SSH_AGENT_PID"))
	http.HandleFunc("/", homePage)
	err := http.ListenAndServe(":"+os.Getenv("SSH_AGENT_PID"), nil)
	//err := http.ListenAndServe(":"+"5202", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprintf(res, "Klim homepage!")
}
