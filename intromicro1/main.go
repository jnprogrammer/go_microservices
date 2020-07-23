package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("GoING FAST BOOIII!!!!  !! ")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "HOW??!", http.StatusBadRequest)
			//rw.WriteHeader(http.StatusBadRequest)  http.Error does the same as these two lines
			//rw.WriteHeader([]byte("HAHT"))
			return
		}

		log.Printf("Data %s\n", d) //takes data to the server but not the user

		fmt.Fprintf(rw, "Hey %s", d) //this allows me
	})

	http.HandleFunc("/money", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Go and give me your MOANAYAAYA !! ")
	})

	http.ListenAndServe(":8710", nil)
}
