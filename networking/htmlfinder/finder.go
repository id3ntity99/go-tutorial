package htmlfinder

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// TODO: Refactor this error handling into elegance way
func ResourceHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the requested file exists
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("[Fatal] Failed getting current working directory\n%s", err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	sep := string(os.PathSeparator)
	paths := strings.Split(r.URL.Path, "/")
	target := wd + sep + "resources" + sep + paths[len(paths)-1]
	fmt.Printf("[INFO] Checking %s...\n", target)
	_, err = os.Stat(target)

	if err != nil { // If the file doesn't exist
		log.Fatalf("[Fatal] Failed getting current working directory\n%s", err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	data, err := os.ReadFile(target)

	if err != nil {
		log.Fatalf("[Fatal] Failed getting current working directory\n%s", err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	//Write bytes read from the target file into response body
	w.Write(data)
}
