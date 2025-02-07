package main

import (
	"log"
	"net/http"
	"os"

	"github.com/id3ntity99/go/networking/htmlfinder"
)

// log.fatal kept calling exit(1), which was constantly shutting down my program
func main() {
	wlog := *log.New(os.Stdout, "[Warning]: ", log.LstdFlags)
	wd, err := os.Getwd()

	h := htmlfinder.NewHandler(wd)

	if err != nil {
		wlog.Println("Failed while getting currently working directory")
	}

	http.HandleFunc("/resources", htmlfinder.Make(h.HandleHTML))

	http.ListenAndServe(":8080", nil)
}
