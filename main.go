package main

import (
	"log"
	"net/http"
	"wxcloudrun-golang/service"
)

func main() {
	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/qigua", service.QiguaHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
