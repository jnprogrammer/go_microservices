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
	//creates the handlers
	producthandler := handlers.NewProducts(logger)

	servemux := mux.NewRouter()

	getRouter := servemux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", producthandler.GetProducts)

	putRouter := servemux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", producthandler.UpdateProducts) //adds identifier by pulling it out of the string using regex

	postRouter := servemux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", producthandler.AddProduct)

	server := &http.Server{
		Addr:         ":8710",
		Handler:      servemux,
		IdleTimeout:  1240 * time.Second, //max time for connections using TCP keep-alive
		ReadTimeout:  3 * time.Second,    //max time to read request from the client
		WriteTimeout: 3 * time.Second,    //max time to write response to client
	}

	go func() {
		err := server.ListenAndServe()
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

	server.Shutdown(tc)
}
