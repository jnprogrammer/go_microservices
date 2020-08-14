package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

//ServeHTTP implements the go http.Handler interface
//https://golang.org/pkg/net/http/#Handler
func (handler *Hello) ServeHTTP(responsewriter http.ResponseWriter, r *http.Request) {
	handler.logger.Println("Handle Hello requests !!!  !! ")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handler.logger.Println("HOW??!", err)
		http.Error(responsewriter, "Unable to read request body", http.StatusBadRequest)
		return
	}
	// write the response
	fmt.Fprintf(responsewriter, "Hello %s", body)
}
