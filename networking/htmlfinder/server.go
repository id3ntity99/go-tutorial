package htmlfinder

import (
	"log"
	"net/http"
	"os"
)

type FetchResourceHandler struct {
	path       string
	infoLogger *log.Logger
}

func (handler *FetchResourceHandler) NewHandler(path string) *FetchResourceHandler {
	return &FetchResourceHandler{
		path: path,
	}
}

func (handler *FetchResourceHandler) HandleFetchResource(w http.ResponseWriter, r *http.Request) error {
	//check if the requested file exists
	stat, err := os.Stat(handler.path)

	if err != nil {
		return err
	} else {
		//Log stat
	}
	//if it exists, write it into response and send it to the client and reutnr nil
	//else, return error
	return nil
}
