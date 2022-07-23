package main

import (
	"fmt"
	"log"
	"net/http"
)

func serverApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "你好")
	})
	if err := http.ListenAndServe("0.0.0.0:10000", mux); err != nil {
		log.Fatal(err)
	}
}
func serverDeBug() {

	if err := http.ListenAndServe("127.0.0.1:10010", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}

func main() {
	go serverApp()
	go serverDeBug()
	select {}
}
