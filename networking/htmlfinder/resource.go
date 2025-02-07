package htmlfinder

import (
	"log"
	"net/http"
	"os"
	"strings"
)

type FetchResourceHandler struct {
	wd         string
	infoLogger *log.Logger
}

// Constructor
func NewHandler(wd string) *FetchResourceHandler {
	return &FetchResourceHandler{
		infoLogger: log.New(os.Stdout, "[INFO] ", log.LstdFlags),
		wd:         wd,
	}
}

// This function returns URI for target file
func (h *FetchResourceHandler) toURL(r *http.Request) string {
	//FIXME Return value has duplicate of "/"
	fname := r.URL.Query().Get("name")
	sep := string(os.PathSeparator)
	uri := strings.Replace(r.URL.Path, "/", "", 1)
	uri = strings.ReplaceAll(uri, "/", sep)
	return h.wd + sep + uri + sep + fname
}

func (h *FetchResourceHandler) HandleHTML(w http.ResponseWriter, r *http.Request) error {
	// 1. Extract file name from query string
	uri := h.toURL(r)
	// 3. stat() the target file
	w.Write([]byte("<h3>" + "uri: " + uri + "<h3>"))
	// 4. if it exists, read the file, else return error
	// 5. Write the file into response body
	return nil
}
