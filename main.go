package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.String("p", ":80", "set port number")

func handler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "(｀･ω･´)"); err != nil {
		log.Println("simple-shakiin-server:", err)
		return
	}
	log.Println("simple-shakiin-server: (｀･ω･´)")
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Println("simple-shakiin-server:", err)
	}
}
