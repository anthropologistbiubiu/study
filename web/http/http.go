package main

import (
	"log"
	"net/http"
)

func main() {

	server := http.Server{
		Addr: ":9090",
		//Handler: http.HandleFunc("http", httpRequest),
	}
	http.HandleFunc("/fetch1", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	log.Print(r.URL)
	w.Write([]byte("{\"name\":\"sunweiming\"}"))
}
