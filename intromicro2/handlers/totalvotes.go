package handlers

import (
	"log"
	"net/http"
)

type getVotes struct {
	l *log.Logger
}

func newgetVotes(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("total number of votes"))
}
