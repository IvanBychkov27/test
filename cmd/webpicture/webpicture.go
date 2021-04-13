package main

import (
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	d := []byte(`<html><body>Hello <br> <img src="/data/img/imag.png"></body></html>`)
	w.Write(d)
}

func main() {
	http.HandleFunc("/", handle)

	staticHandler := http.StripPrefix("data", http.FileServer(http.Dir("./static")))
	http.Handle("/data/", staticHandler)
	log.Println("server is listening... 127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
