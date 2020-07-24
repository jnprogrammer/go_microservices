package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("So much LOGGING !!!  !! ")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "HOW??!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hey %s", d) //this allows me to send data back to user
}
