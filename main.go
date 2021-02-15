package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Cardano time")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oopps? ?", http.StatusBadRequest)
			return
		}

		//log.Printf("Data %s", d)
		fmt.Fprintf(rw, "You sent the data: %s \n", d)
	})

	http.HandleFunc("/ADA", func(http.ResponseWriter, *http.Request) {
		log.Println("₳")
	})

	http.ListenAndServe(":9090", nil)
}
