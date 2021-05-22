package main

//
import (
	"flag"
	"log"
	"webclient/webclient"
)

func main() {
	var (
		url = flag.String("url", "http://localhost:8080", "input markdown file")
	)
	flag.Parse()

	log.SetFlags(log.Lshortfile)
	log.Printf("START")

	err := webclient.Get(*url, "")
	//err := webclient.Get("http://www.yahoo.co.jp", "")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("END")
}
