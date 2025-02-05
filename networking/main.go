package main

import (
	"log"
	"net/http"

	"github.com/id3ntity99/go/networking/basics"
	"github.com/id3ntity99/go/networking/htmlfinder"
)

func main() {
	http.HandleFunc("/hello", basics.HelloHandler)
	http.HandleFunc("/hello/world", basics.HelloHandler)
	http.HandleFunc("/welcome", basics.WelcomeHandler)

	http.HandleFunc("/files/", htmlfinder.ResourceHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
