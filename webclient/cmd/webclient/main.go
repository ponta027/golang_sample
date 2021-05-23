package main

//
import (
	"flag"
	"log"
	"webclient/webclient"
)

func main() {
	var (
		method = flag.String("m", "GET", "Request Method")
		url    = flag.String("url", "http://localhost:8080", "input markdown file")
		query  = flag.String("q", "", "query")
	)
	flag.Parse()

	log.SetFlags(log.Lshortfile)
	log.Printf("START")

	var err error = nil
	if *method == "GET" {
		err = webclient.Get(*url, *query)
	} else {
		err = webclient.Post(*url, *query)
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("END")
}
