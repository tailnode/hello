package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	output := "hostname: " + getHostName() + " version: " + os.Getenv("HELLO_VERSION") + "\n"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("receive request")
		io.WriteString(w, output)
	})

	fmt.Println("listening :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown_host"
	}
	return hostname
}
