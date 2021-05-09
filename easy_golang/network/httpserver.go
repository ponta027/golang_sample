package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>httpserverへようこそ</h1>")
	fmt.Fprintf(w, "<p>Host:%q</p>", r.Host)
	fmt.Fprintf(w, "<p>Remote:%q</p>", r.RemoteAddr)
	fmt.Fprintf(w, "</body></html>")

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
}

func main() {

	fmt.Println("start")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("end")

}
