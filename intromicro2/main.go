package main

import (
	"context"
	"github.com/jnprogrammer/go_microservices/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	heyhandler := handlers.NewHello(logger)
	byehandler := handlers.NewGoodBye(logger)

	servemux := http.NewServeMux()
	servemux.Handle("/", heyhandler)
	servemux.Handle("/bye", byehandler)

	s := &http.Server{
		Addr:         ":8710",
		Handler:      servemux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
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
