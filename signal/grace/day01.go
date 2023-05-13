package main

import (
	"log"
	"net/http"
)

// startHttpServer start http server
func startHttpServer() {
	mux := http.NewServeMux()
	if err := http.ListenAndServe(":1608", mux); err != nil {
		log.Fatal("startHttpServer ListenAndServe error: " + err.Error())
	}
}

// startHttpServer start http server
