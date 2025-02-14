package htmlfinder

import (
	"fmt"
	"net/http"
	"os"
)

//A helper class

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func Make(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: Handle all errors here.
		err := f(w, r)

		if err != nil {
			switch e := err.(type) {
			case (*os.PathError):
				msg := fmt.Sprintf("<h1 style=\"color: red\">404 - Not Found</h1><p>stack trace </br>%s</p>", e.Error())
				w.WriteHeader(http.StatusNotFound)
				w.Header().Set("Content-type", "text/html;charset=utf-8")
				w.Write([]byte(msg))
			}
		}
	}
}
