package htmlfinder

import (
	"log"
	"net/http"
	"os"
	"strings"
)

type FetchResourceHandler struct {
	path       string
	infoLogger *log.Logger
}

// Constructor
func NewHandler(path string) *FetchResourceHandler {
	return &FetchResourceHandler{
		infoLogger: log.New(os.Stdout, "[INFO] ", log.LstdFlags),
		path:       path,
	}
}

func (handler *FetchResourceHandler) HandleFetchResource(w http.ResponseWriter, r *http.Request) error {
	sepstr := string(os.PathSeparator)
	target := handler.path + strings.ReplaceAll(r.URL.Path, "/", sepstr)
	//check if the requested file exists
	stat, err := os.Stat(target)

	if err != nil { // If error occurs
		return err
	} else {
		//Log stat.Name()
		handler.infoLogger.Printf("Requested resource %s has found\n", stat.Name())
	}

	//if the file exists, write it into response body and send it to the client and return nil
	file, err := os.ReadFile(target)

	if err != nil {
		return err
	}

	w.Header().Set("Content-type", "text/html;charset=utf-8")
	w.Write(file)

	return nil
}
