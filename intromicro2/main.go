package main

import (
	"github.com/jnprogrammer/go_microservices/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	heyhandler := handlers.NewHello(logger)
	byehandler := handlers.NewGoodBye(logger)

	servemux := http.NewServeMux()
	servemux.Handle("/", heyhandler)
	servemux.Handle("/bye", byehandler)

	//s := &http.Server
	http.ListenAndServe(":8710", servemux)
}
