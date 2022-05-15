package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

var host string
var startTime time.Time
var version string
var requests = 0

func handler(w http.ResponseWriter, req *http.Request) {

	now := time.Now()
	requests++
	fmt.Printf("Running On: %v | Requests: %v | Uptime: %vs | Log Time: %v\n", host, requests, now.Sub(startTime), now.Format(time.UnixDate))

	fmt.Fprintf(w, "Hello Kubernetes bootcamp go! | Running on: %v | v=%v\n", host, version)

}

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8090
	}

	if os.Getenv("HOSTNAME") != "" {
		host = os.Getenv("HOSTNAME")
	} else if os.Getenv("HOST") != "" {
		host = os.Getenv("HOST")
	} else {
		host = "Undetermined"
	}

	version = os.Getenv("VERSION")
	if version == "" {
		version = "v1"
	}

	startTime = time.Now()

	http.HandleFunc("/", handler)

	fmt.Printf("Starting server on port %v...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
