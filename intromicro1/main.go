package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GoING FAST BOOIII!!!!  !! ")
	})

	http.HandleFunc("/money", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Go and give me your MOANAYAAYA !! ")
	})

	http.ListenAndServe(":8710", nil)
}
