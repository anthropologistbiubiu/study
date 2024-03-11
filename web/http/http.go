package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	server := http.Server{
		Addr: ":9090",
		//Handler: http.HandleFunc("http", httpRequest),
	}
	http.HandleFunc("/submit", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err := r.ParseForm()
	if err != nil {
		log.Println("err", err.Error())
	}
	w.Write([]byte("{\"name\":\"sunweiming\"}"))
	username := r.FormValue("username")
	passwd := r.FormValue("password")
	fmt.Println(username, passwd)
}
