package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("here")
		fmt.Fprint(w, "hello")
		//w.WriteHeader(419)
		//w.Write([]byte("SSL certificate has been checked"))

	})
	log.Printf("About to listen on 9000. Go to https://notionassistant.publicvm.com/")
	err := http.ListenAndServe(":9000", nil)
	log.Fatal(err)
}
