package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8090
	}

	srv := new(server)

	// srv.auth.username = os.Getenv("AUTH_USERNAME")
	// srv.auth.password = os.Getenv("AUTH_PASSWORD")

	// if srv.auth.username == "" {
	// 	log.Fatal("basic auth username must be provided")
	// }

	// if srv.auth.password == "" {
	// 	log.Fatal("basic auth password must be provided")
	// }

	if os.Getenv("HOSTNAME") != "" {
		srv.info.host = os.Getenv("HOSTNAME")
	} else if os.Getenv("HOST") != "" {
		srv.info.host = os.Getenv("HOST")
	} else {
		srv.info.host = "Undetermined"
	}

	srv.info.version = os.Getenv("VERSION")
	if srv.info.version == "" {
		srv.info.version = "v1"
	}

	srv.info.startTime = time.Now()

	// http.HandleFunc("/", srv.basicAuth(srv.handler))
	http.HandleFunc("/", srv.handler)
	http.HandleFunc("/healthz", srv.healthHandler)

	fmt.Printf("Starting server on port %v...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)

	// TODO: How to load certificates in the cluster?
	// http.ListenAndServeTLS(
	// 	fmt.Sprintf(":%v", port),
	// 	"./localhost.pem",
	// 	"./localhost-key.pem",
	// 	nil)
}
