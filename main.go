package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/jnprogrammer/go_microservices/product-api/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(logger)

	servemux := http.NewServeMux()
	servemux.Handle("/", ph)

	s := &http.Server{
		Addr:         ":8710",
		Handler:      servemux,
		IdleTimeout:  120 * time.Second, //max time for connections using TCP keep-alive
		ReadTimeout:  1 * time.Second,   //max time to read request from the client
		WriteTimeout: 1 * time.Second,   //max time to write response to client
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("I'm shutting down dudes!", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}
