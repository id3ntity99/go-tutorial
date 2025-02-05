package basics

import (
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	path := r.URL.Path
	message := "<h3>Hello, world! You are at " + path
	w.Write([]byte(message))
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html;charest=utf-8")
	path := r.URL.Path
	message := "<h3>Welcome, world! You are at " + path
	w.Write([]byte(message))
}
